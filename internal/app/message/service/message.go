package service

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/message/repository"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc/exception"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc/md"
	"github.com/tsundata/assistant/internal/pkg/util"
	"gorm.io/gorm"
	"strings"
)

type Message struct {
	bus     event.Bus
	config  *config.AppConfig
	logger  log.Logger
	repo    repository.MessageRepository
	chatbot pb.ChatbotSvcClient
}

func NewMessage(
	bus event.Bus,
	logger log.Logger,
	config *config.AppConfig,
	repo repository.MessageRepository,
	chatbot pb.ChatbotSvcClient) *Message {
	return &Message{
		bus:     bus,
		logger:  logger,
		config:  config,
		repo:    repo,
		chatbot: chatbot,
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
	message.Type = string(enum.MessageTypeText)
	message.Text = strings.TrimSpace(payload.Message.GetText())
	if payload.Message.GetPayload() != "" {
		message.Payload = payload.Message.GetPayload()
	} else {
		message.Payload = "{}"
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

	if enum.MessageType(message.Type) == enum.MessageTypeScript {
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
	} else {
		// bot handle
		err = m.bus.Publish(ctx, enum.Message, event.MessageHandleSubject, message)
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
	group, err := m.chatbot.GetGroupId(ctx, &pb.UuidRequest{Uuid: payload.GroupUuid})
	if err != nil {
		return nil, err
	}
	total, messages, err := m.repo.ListByGroup(ctx, group.Id, int(payload.Page), int(payload.Limit))
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
			return nil, exception.ErrGrpcUnauthenticated
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
	_, err := m.repo.Create(ctx, payload.Message)
	if err != nil {
		return nil, err
	}
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

	// todo send message

	// push inbox
	_, err := m.repo.CreateInbox(ctx, pb.Inbox{
		UserId:     payload.Message.GetUserId(),
		Sender:     payload.Message.GetSender(),
		SenderType: payload.Message.GetSenderType(),
		Type:       payload.Message.GetType(),
		Title:      util.SubString(payload.Message.GetText(), 0, 100),
		Content:    util.SubString(payload.Message.GetText(), 0, 2000),
		Payload:    payload.Message.GetPayload(),
	})
	if err != nil {
		return nil, err
	}

	// setting

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
		wfReply, err := m.chatbot.RunAction(ctx, &pb.WorkflowRequest{Text: message.RemoveActionScriptFlag()})
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

func (m *Message) GetActionMessages(ctx context.Context, _ *pb.TextRequest) (*pb.ActionReply, error) {
	var items []*pb.Message
	items, err := m.repo.ListByType(ctx, string(enum.MessageTypeScript))
	if err != nil {
		return nil, err
	}

	var kvs []*pb.Action
	for _, item := range items {
		kvs = append(kvs, &pb.Action{
			Id:   item.Id,
			Text: item.Text,
		})
	}

	return &pb.ActionReply{
		Items: kvs,
	}, nil
}

func (m *Message) CreateActionMessage(ctx context.Context, payload *pb.TextRequest) (*pb.StateReply, error) {
	if payload.GetText() == "" {
		return &pb.StateReply{State: false}, nil
	}

	// check syntax
	_, err := m.chatbot.SyntaxCheck(ctx, &pb.WorkflowRequest{
		Text: payload.GetText(),
		Type: string(enum.MessageTypeScript),
	})
	if err != nil {
		return nil, err
	}

	// store message
	uuid := util.UUID()
	id, err := m.repo.Create(ctx, &pb.Message{
		Uuid: uuid,
		Type: string(enum.MessageTypeScript),
		Text: payload.GetText(),
	})
	if err != nil {
		return nil, err
	}

	// check/create trigger
	_, err = m.chatbot.CreateTrigger(ctx, &pb.TriggerRequest{
		Trigger: &pb.Trigger{
			Kind:      string(enum.MessageTypeScript),
			MessageId: id,
		},
		Info: &pb.TriggerInfo{
			MessageText: payload.GetText(),
		},
	})
	if err != nil {
		return nil, err
	}

	return &pb.StateReply{State: true}, nil
}

func (m *Message) DeleteWorkflowMessage(ctx context.Context, payload *pb.MessageRequest) (*pb.StateReply, error) {
	err := m.repo.Delete(ctx, payload.Message.GetId())
	if err != nil {
		return nil, err
	}

	_, err = m.chatbot.DeleteTrigger(ctx, &pb.TriggerRequest{Trigger: &pb.Trigger{MessageId: payload.Message.GetId()}})
	if err != nil {
		return nil, err
	}

	return &pb.StateReply{State: true}, nil
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
