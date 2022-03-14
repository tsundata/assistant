package service

import (
	"context"
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
	message.Type = enum.MessageTypeText
	message.Text = strings.TrimSpace(payload.Message.GetText())

	// check
	find, err := m.repo.GetByUUID(ctx, message.Uuid)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) && find.Id > 0 {
		return &pb.MessageReply{
			Message: &pb.Message{
				Uuid: message.Uuid,
				Type: message.Type,
				Text: message.Text,
			},
		}, nil
	}

	// parse type
	message.Text = strings.TrimSpace(message.Text)
	if util.IsUrl(message.Text) {
		message.Type = enum.MessageTypeLink
	}
	if message.IsMessageOfAction() {
		message.Type = enum.MessageTypeAction
	}

	// store
	_, err = m.repo.Create(ctx, &message)
	if err != nil {
		return nil, err
	}

	if message.Type == enum.MessageTypeAction {
		_, err = m.chatbot.CreateTrigger(ctx, &pb.TriggerRequest{
			Trigger: &pb.Trigger{
				Kind:      enum.MessageTypeAction,
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

	var reply []*pb.Message
	for _, item := range messages {
		if item.UserId != id {
			return nil, exception.ErrGrpcUnauthenticated
		}

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
		Title:      payload.Message.GetText(),
		Content:    payload.Message.GetText(),
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

	switch message.Type {
	case enum.MessageTypeAction:
		wfReply, err := m.chatbot.RunAction(ctx, &pb.WorkflowRequest{Text: message.RemoveActionFlag()})
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
	items, err := m.repo.ListByType(ctx, enum.MessageTypeAction)
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
		Type: enum.MessageTypeAction,
	})
	if err != nil {
		return nil, err
	}

	// store message
	uuid := util.UUID()
	id, err := m.repo.Create(ctx, &pb.Message{
		Uuid: uuid,
		Type: enum.MessageTypeAction,
		Text: payload.GetText(),
	})
	if err != nil {
		return nil, err
	}

	// check/create trigger
	_, err = m.chatbot.CreateTrigger(ctx, &pb.TriggerRequest{
		Trigger: &pb.Trigger{
			Kind:      enum.MessageTypeAction,
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
