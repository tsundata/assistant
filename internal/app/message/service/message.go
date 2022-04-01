package service

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/message/repository"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/push"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc/md"
	"github.com/tsundata/assistant/internal/pkg/util"
	"github.com/tsundata/assistant/internal/pkg/vendors"
	"gorm.io/gorm"
	"io"
	"strings"
)

type Message struct {
	bus     event.Bus
	config  *config.AppConfig
	logger  log.Logger
	redis   *redis.Client
	repo    repository.MessageRepository
	chatbot pb.ChatbotSvcClient
	storage pb.StorageSvcClient
}

func NewMessage(
	bus event.Bus,
	logger log.Logger,
	redis *redis.Client,
	config *config.AppConfig,
	repo repository.MessageRepository,
	chatbot pb.ChatbotSvcClient,
	storage pb.StorageSvcClient) *Message {
	return &Message{
		bus:     bus,
		logger:  logger,
		redis:   redis,
		config:  config,
		repo:    repo,
		chatbot: chatbot,
		storage: storage,
	}
}

func (m *Message) Create(ctx context.Context, payload *pb.MessageRequest) (*pb.MessageReply, error) {
	// check uuid
	var message pb.Message
	message.UserId = payload.Message.GetUserId()
	message.GroupId = payload.Message.GetGroupId()
	message.Uuid = payload.Message.GetUuid()
	message.Sender = payload.Message.GetUserId()
	message.SenderType = enum.MessageUserType
	message.Receiver = payload.Message.GetGroupId()
	message.ReceiverType = enum.MessageGroupType
	if payload.Message.GetPayload() != "" {
		message.Payload = payload.Message.GetPayload()
	} else {
		message.Payload = "{}"
	}

	// before
	switch enum.MessageType(payload.Message.Type) {
	case enum.MessageTypeImage:
		message.Type = string(enum.MessageTypeImage)
	default:
		message.Type = string(enum.MessageTypeText)
		message.Text = strings.TrimSpace(payload.Message.GetText())
	}

	// check
	find, err := m.repo.GetByUUID(ctx, message.Uuid)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) && find.Id > 0 {
		return &pb.MessageReply{
			Message: &pb.Message{
				Uuid:    message.Uuid,
				Type:    message.Type,
				Text:    message.Text,
				Payload: message.Payload,
			},
		}, nil
	}

	// parse type
	message.Text = strings.TrimSpace(message.Text)
	if util.IsUrl(message.Text) {
		message.Type = string(enum.MessageTypeLink)
	}
	if message.IsMessageOfActionScript() {
		p, err := json.Marshal(pb.ScriptMsg{
			Kind: enum.ActionScript,
			Code: message.Text,
		})
		if err != nil {
			return nil, err
		}
		message.Type = string(enum.MessageTypeScript)
		message.Payload = util.ByteToString(p)
	}

	// store
	_, err = m.repo.Create(ctx, &message)
	if err != nil {
		return nil, err
	}

	// after
	switch enum.MessageType(message.Type) {
	case enum.MessageTypeScript:
		_, err = m.chatbot.CreateTrigger(ctx, &pb.TriggerRequest{
			Trigger: &pb.Trigger{
				Kind:      string(enum.MessageTypeScript),
				MessageId: message.Id,
			},
			Info: &pb.TriggerInfo{
				MessageText: message.Text,
			},
		})
		if err != nil {
			return nil, err
		}
		err = m.bus.Publish(ctx, enum.Chatbot, event.WorkflowRunSubject, message)
		if err != nil {
			return nil, err
		}
	case enum.MessageTypeImage:
		data := bytes.NewReader(payload.Message.Data)
		buf := make([]byte, 1024)
		uc, err := m.storage.UploadFile(ctx)
		if err != nil {
			return nil, err
		}
		err = uc.Send(&pb.FileRequest{
			Data: &pb.FileRequest_Info{Info: &pb.FileInfo{FileType: "png"}},
		})
		if err != nil {
			return nil, err
		}
		for {
			n, err := data.Read(buf)
			if err == io.EOF {
				break
			}
			if err != nil {
				return nil, err
			}
			err = uc.Send(&pb.FileRequest{Data: &pb.FileRequest_Chuck{Chuck: buf[:n]}})
			if err != nil {
				return nil, err
			}
		}
		fileReply, err := uc.CloseAndRecv()
		if err != nil {
			return nil, err
		}
		imageMsg := pb.ImageMsg{
			Src: fileReply.Path,
		}
		p, err := json.Marshal(imageMsg)
		if err != nil {
			return nil, err
		}
		message.Payload = util.ByteToString(p)
		err = m.repo.Save(ctx, &message)
		if err != nil {
			return nil, err
		}
	default:
		// bot handle
		err = m.bus.Publish(ctx, enum.Message, event.BotHandleSubject, message)
		if err != nil {
			return nil, err
		}
	}

	return &pb.MessageReply{
		Message: &message,
	}, nil
}

func (m *Message) List(ctx context.Context, _ *pb.MessageRequest) (*pb.MessagesReply, error) {
	messages, err := m.repo.List(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.MessagesReply{Messages: messages}, nil
}

func (m *Message) ListByGroup(ctx context.Context, payload *pb.GetMessagesRequest) (*pb.GetMessagesReply, error) {
	id, _ := md.FromIncoming(ctx)
	total, messages, err := m.repo.ListByGroup(ctx, payload.GroupId, int(payload.Page), int(payload.Limit))
	if err != nil {
		return nil, err
	}

	// avatar
	var botId []int64
	for _, item := range messages {
		if item.SenderType == enum.MessageBotType && item.Sender > 0 {
			botId = append(botId, item.Sender)
		}
	}
	bots, err := m.chatbot.GetBots(ctx, &pb.BotsRequest{BotId: botId})
	if err != nil {
		return nil, err
	}
	botMap := make(map[int64]*pb.Bot)
	for i, item := range bots.Bots {
		botMap[item.Id] = bots.Bots[i]
	}

	var reply []*pb.Message
	for _, item := range messages {
		if item.UserId != id {
			continue
		}

		// avatar
		var avatar *pb.Avatar
		if item.SenderType == enum.MessageBotType {
			if v, ok := botMap[item.Sender]; ok {
				avatar = &pb.Avatar{
					Name:       v.Name,
					Src:        v.Avatar,
					Identifier: v.Identifier,
				}
			}
		}
		item.Avatar = avatar

		// covert
		direction := ""
		if item.SenderType == enum.MessageBotType || item.SenderType == enum.MessageGroupType {
			direction = enum.MessageIncomingDirection
		} else {
			direction = enum.MessageOutgoingDirection
		}
		item.Direction = direction
		item.SendTime = util.Format(item.CreatedAt)

		reply = append(reply, item)
	}

	return &pb.GetMessagesReply{
		Total:    total,
		Page:     payload.Page,
		PageSize: payload.Limit,
		Messages: reply,
	}, nil
}

func (m *Message) GetByUuid(ctx context.Context, payload *pb.MessageRequest) (*pb.GetMessageReply, error) {
	message, err := m.repo.GetByUUID(ctx, payload.Message.GetUuid())
	if err != nil {
		return nil, err
	}

	// covert
	direction := ""
	if message.SenderType == enum.MessageBotType || message.SenderType == enum.MessageGroupType {
		direction = enum.MessageIncomingDirection
	} else {
		direction = enum.MessageOutgoingDirection
	}
	message.Direction = direction
	message.SendTime = util.Format(message.CreatedAt)

	return &pb.GetMessageReply{
		Message: &message,
	}, nil
}

func (m *Message) GetById(ctx context.Context, payload *pb.MessageRequest) (*pb.GetMessageReply, error) {
	message, err := m.repo.GetByID(ctx, payload.Message.GetId())
	if err != nil {
		return nil, err
	}

	// covert
	direction := ""
	if message.SenderType == enum.MessageBotType || message.SenderType == enum.MessageGroupType {
		direction = enum.MessageIncomingDirection
	} else {
		direction = enum.MessageOutgoingDirection
	}
	message.Direction = direction
	message.SendTime = util.Format(message.CreatedAt)

	return &pb.GetMessageReply{
		Message: &message,
	}, nil
}

func (m *Message) GetBySequence(ctx context.Context, payload *pb.MessageRequest) (*pb.GetMessageReply, error) {
	message, err := m.repo.GetBySequence(ctx, payload.Message.GetUserId(), payload.Message.GetId())
	if err != nil {
		return nil, err
	}

	// covert
	direction := ""
	if message.SenderType == enum.MessageBotType || message.SenderType == enum.MessageGroupType {
		direction = enum.MessageIncomingDirection
	} else {
		direction = enum.MessageOutgoingDirection
	}
	message.Direction = direction
	message.SendTime = util.Format(message.CreatedAt)

	return &pb.GetMessageReply{
		Message: &message,
	}, nil
}

func (m *Message) LastByGroup(ctx context.Context, payload *pb.LastByGroupRequest) (*pb.LastByGroupReply, error) {
	message, err := m.repo.GetLastByGroup(ctx, payload.GroupId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &pb.LastByGroupReply{
			Message: &pb.Message{},
		}, nil
	}

	return &pb.LastByGroupReply{
		Message: &message,
	}, nil
}

func (m *Message) Save(ctx context.Context, payload *pb.MessageRequest) (*pb.MessageReply, error) {
	id, err := m.repo.Create(ctx, payload.Message)
	if err != nil {
		return nil, err
	}
	payload.Message.Id = id
	return &pb.MessageReply{Message: payload.Message}, nil
}

func (m *Message) Delete(ctx context.Context, payload *pb.MessageRequest) (*pb.TextReply, error) {
	err := m.repo.Delete(ctx, payload.Message.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.TextReply{Text: ""}, nil
}

func (m *Message) Send(ctx context.Context, payload *pb.MessageRequest) (*pb.StateReply, error) {
	if payload.Message.GetText() == "" {
		return &pb.StateReply{State: false}, nil
	}

	// push provider
	mapCmd := m.redis.HGetAll(ctx, "system:push:switch")
	for k, v := range mapCmd.Val() {
		if v == "1" {
			provider := vendors.NewPushProvider(k)
			if provider == nil {
				continue
			}
			err := provider.Send(push.Message{
				Title:   util.SubString(payload.Message.GetText(), 0, 100),
				Content: util.SubString(payload.Message.GetText(), 0, 2000),
			})
			if err != nil {
				m.logger.Error(err)
			}
		}
	}

	// push inbox
	_, err := m.repo.CreateInbox(ctx, pb.Inbox{
		UserId:     payload.Message.GetUserId(),
		Sender:     payload.Message.GetSender(),
		SenderType: payload.Message.GetSenderType(),
		Type:       payload.Message.GetType(),
		Title:      util.SubString(payload.Message.GetText(), 0, 100),
		Content:    util.SubString(payload.Message.GetText(), 0, 2000),
		Payload:    payload.Message.GetPayload(),
		Status:     enum.InboxCreate,
	})
	if err != nil {
		return nil, err
	}

	// push ws hub
	err = m.bus.Publish(ctx, enum.Message, event.MessageChannelSubject, payload.Message)
	if err != nil {
		return nil, err
	}

	return &pb.StateReply{
		State: true,
	}, nil
}

func (m *Message) Run(ctx context.Context, payload *pb.MessageRequest) (*pb.TextReply, error) {
	var reply string
	message, err := m.repo.GetByID(ctx, payload.Message.GetId())
	if err != nil {
		return nil, err
	}

	switch enum.MessageType(message.Type) {
	case enum.MessageTypeScript:
		wfReply, err := m.chatbot.RunActionScript(ctx, &pb.WorkflowRequest{Text: message.RemoveActionScriptFlag()})
		if err != nil {
			return nil, err
		}
		reply = wfReply.GetText()
	default:
		reply = "Not running"
	}

	return &pb.TextReply{
		Text: reply,
	}, nil
}

func (m *Message) Action(ctx context.Context, payload *pb.ActionRequest) (*pb.ActionReply, error) {
	// store
	id, _ := md.FromIncoming(ctx)
	message, err := m.repo.GetByID(ctx, payload.MessageId)
	if err != nil {
		return nil, err
	}
	if message.Status == enum.MessageActionedStatus {
		return &pb.ActionReply{State: true}, nil
	}
	var p pb.ActionMsg
	err = json.Unmarshal(util.StringToByte(message.Payload), &p)
	if err != nil {
		return nil, err
	}
	p.Value = payload.Value
	data, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	err = m.repo.SavePayload(ctx, payload.MessageId, util.ByteToString(data))
	if err != nil {
		return nil, err
	}

	// event
	err = m.bus.Publish(ctx, enum.Chatbot, event.BotActionSubject, pb.Message{
		UserId:     id,
		GroupId:    message.GroupId,
		Sender:     message.Sender,
		SenderType: enum.MessageBotType,
		Type:       string(enum.MessageTypeAction),
		Payload:    util.ByteToString(data),
	})
	if err != nil {
		return nil, err
	}

	return &pb.ActionReply{State: true}, nil
}

func (m *Message) Form(ctx context.Context, payload *pb.FormRequest) (*pb.FormReply, error) {
	// store
	id, _ := md.FromIncoming(ctx)
	kvMap := make(map[string]interface{})
	for _, item := range payload.Form {
		kvMap[item.Key] = item.Value
	}
	message, err := m.repo.GetByID(ctx, payload.MessageId)
	if err != nil {
		return nil, err
	}
	if message.Status == enum.MessageActionedStatus {
		return &pb.FormReply{State: true}, nil
	}
	var p pb.FormMsg
	err = json.Unmarshal(util.StringToByte(message.Payload), &p)
	if err != nil {
		return nil, err
	}
	for i, item := range p.Field {
		if v, ok := kvMap[item.Key]; ok {
			p.Field[i].Value = v
		}
	}
	data, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	err = m.repo.SavePayload(ctx, payload.MessageId, util.ByteToString(data))
	if err != nil {
		return nil, err
	}

	// event
	err = m.bus.Publish(ctx, enum.Chatbot, event.BotFormSubject, pb.Message{
		UserId:     id,
		GroupId:    message.GroupId,
		Sender:     message.Sender,
		SenderType: enum.MessageBotType,
		Type:       string(enum.MessageTypeForm),
		Payload:    util.ByteToString(data),
	})
	if err != nil {
		return nil, err
	}

	return &pb.FormReply{State: true}, nil
}

func (m *Message) ListInbox(ctx context.Context, payload *pb.InboxRequest) (*pb.InboxReply, error) {
	id, _ := md.FromIncoming(ctx)
	total, list, err := m.repo.ListInbox(ctx, id, int(payload.Page), int(payload.Limit))
	if err != nil {
		return nil, err
	}
	return &pb.InboxReply{
		Total:    total,
		Page:     payload.Page,
		PageSize: payload.Limit,
		Inbox:    list,
	}, nil
}

func (m *Message) LastInbox(ctx context.Context, _ *pb.InboxRequest) (*pb.InboxReply, error) {
	id, _ := md.FromIncoming(ctx)
	inbox, err := m.repo.LastInbox(ctx, id)
	if err != nil {
		return nil, err
	}
	return &pb.InboxReply{Inbox: []*pb.Inbox{&inbox}}, nil
}

func (m *Message) MarkSendInbox(ctx context.Context, payload *pb.InboxRequest) (*pb.InboxReply, error) {
	err := m.repo.UpdateInboxStatus(ctx, payload.InboxId, enum.InboxSend)
	if err != nil {
		return nil, err
	}
	return &pb.InboxReply{}, nil
}

func (m *Message) MarkReadInbox(ctx context.Context, payload *pb.InboxRequest) (*pb.InboxReply, error) {
	err := m.repo.UpdateInboxStatus(ctx, payload.InboxId, enum.InboxRead)
	if err != nil {
		return nil, err
	}
	return &pb.InboxReply{}, nil
}
