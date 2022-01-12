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
	"github.com/tsundata/assistant/internal/pkg/util"
	"gorm.io/gorm"
	"strings"
)

type Message struct {
	bus      event.Bus
	config   *config.AppConfig
	logger   log.Logger
	repo     repository.MessageRepository
	workflow pb.WorkflowSvcClient
	chatbot  pb.ChatbotSvcClient
}

func NewMessage(
	bus event.Bus,
	logger log.Logger,
	config *config.AppConfig,
	repo repository.MessageRepository,
	workflow pb.WorkflowSvcClient,
	chatbot pb.ChatbotSvcClient) *Message {
	return &Message{
		bus:      bus,
		logger:   logger,
		config:   config,
		repo:     repo,
		workflow: workflow,
		chatbot:  chatbot,
	}
}

func (m *Message) List(ctx context.Context, _ *pb.MessageRequest) (*pb.MessagesReply, error) {
	messages, err := m.repo.List(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.MessagesReply{Messages: messages}, nil
}

func (m *Message) ListByGroup(ctx context.Context, payload *pb.GetMessagesRequest) (*pb.GetMessagesReply, error) {
	group, err := m.chatbot.GetGroupId(ctx, &pb.UuidRequest{Uuid: payload.GroupUuid})
	if err != nil {
		return nil, err
	}
	messages, err := m.repo.ListByGroup(ctx, group.Id, int(payload.Page), int(payload.Limit))
	if err != nil {
		return nil, err
	}

	var reply []*pb.MessageItem
	for _, item := range messages {
		reply = append(reply, &pb.MessageItem{
			Uuid:    item.Uuid,
			Message: item.Text,
			Type:    item.Type,
		})
	}

	return &pb.GetMessagesReply{
		Messages: reply,
	}, nil
}

func (m *Message) Get(ctx context.Context, payload *pb.MessageRequest) (*pb.GetMessageReply, error) {
	message, err := m.repo.GetByID(ctx, payload.Message.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.GetMessageReply{
		Message: &pb.MessageItem{
			Uuid:     message.Uuid,
			Message:  message.Text,
			Type:     message.Type,
			SendTime: "", // todo
		},
	}, nil
}

func (m *Message) LastByGroup(ctx context.Context, payload *pb.LastByGroupRequest) (*pb.LastByGroupReply, error) {
	message, err := m.repo.GetLastByGroup(ctx, payload.GroupId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &pb.LastByGroupReply{
			Message: &pb.MessageItem{},
		}, nil
	}

	return &pb.LastByGroupReply{
		Message: &pb.MessageItem{
			Uuid:     message.Uuid,
			Message:  message.Text,
			Type:     message.Type,
			SendTime: "", // todo
		},
	}, nil
}

func (m *Message) Create(ctx context.Context, payload *pb.MessageRequest) (*pb.MessageReply, error) {
	// check uuid
	var message pb.Message
	message.Uuid = payload.Message.GetUuid()
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

	// bot handle
	err = m.bus.Publish(ctx, event.MessageHandleSubject, message)
	if err != nil {
		return nil, err
	}

	return &pb.MessageReply{
		Message: &message,
	}, nil
}

func (m *Message) Delete(ctx context.Context, payload *pb.MessageRequest) (*pb.TextReply, error) {
	err := m.repo.Delete(ctx, payload.Message.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.TextReply{Text: ""}, nil
}

func (m *Message) Send(_ context.Context, payload *pb.MessageRequest) (*pb.StateReply, error) {
	if payload.Message.GetText() == "" {
		return &pb.StateReply{State: false}, nil
	}

	// todo send message

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
		wfReply, err := m.workflow.RunAction(ctx, &pb.WorkflowRequest{Text: message.RemoveActionFlag()})
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
	_, err := m.workflow.SyntaxCheck(ctx, &pb.WorkflowRequest{
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
	_, err = m.workflow.CreateTrigger(ctx, &pb.TriggerRequest{
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

	_, err = m.workflow.DeleteTrigger(ctx, &pb.TriggerRequest{Trigger: &pb.Trigger{MessageId: payload.Message.GetId()}})
	if err != nil {
		return nil, err
	}

	return &pb.StateReply{State: true}, nil
}
