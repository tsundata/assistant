package service

import (
	"context"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/message/repository"
	"github.com/tsundata/assistant/internal/app/message/rule"
	"github.com/tsundata/assistant/internal/app/message/trigger"
	"github.com/tsundata/assistant/internal/app/message/trigger/ctx"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/model"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
	"github.com/tsundata/assistant/internal/pkg/transport/http"
	"github.com/tsundata/assistant/internal/pkg/util"
	"github.com/valyala/fasthttp"
	"strings"
	"sync"
)

type Message struct {
	config   *config.AppConfig
	logger   *logger.Logger
	bot      *rulebot.RuleBot
	repo     repository.MessageRepository
	workflow pb.WorkflowClient
	middle   pb.MiddleClient
	todo     pb.TodoClient
}

func NewMessage(
	logger *logger.Logger,
	config *config.AppConfig,
	repo repository.MessageRepository,
	workflow pb.WorkflowClient,
	middle pb.MiddleClient,
	todo pb.TodoClient,
	bot *rulebot.RuleBot) *Message {
	return &Message{
		logger:   logger,
		bot:      bot,
		config:   config,
		repo:     repo,
		workflow: workflow,
		middle:   middle,
		todo:     todo,
	}
}

func (m *Message) List(_ context.Context, _ *pb.MessageRequest) (*pb.MessageListReply, error) {
	messages, err := m.repo.List()
	if err != nil {
		return nil, err
	}

	var reply []*pb.MessageItem
	for _, item := range messages {
		reply = append(reply, &pb.MessageItem{
			Uuid: item.UUID,
			Text: item.Text,
			Type: item.Type,
			Time: item.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return &pb.MessageListReply{
		Messages: reply,
	}, nil
}

func (m *Message) Get(_ context.Context, payload *pb.MessageRequest) (*pb.MessageReply, error) {
	message, err := m.repo.GetByID(payload.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.MessageReply{
		Id:   int64(message.ID),
		Uuid: message.UUID,
		Text: message.Text,
		Type: message.Type,
		Time: message.CreatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}

func (m *Message) Create(_ context.Context, payload *pb.MessageRequest) (*pb.TextsReply, error) {
	// check uuid
	var message model.Message
	message.UUID = payload.GetUuid()
	message.Type = model.MessageTypeText
	message.Text = strings.TrimSpace(payload.GetText())

	// check
	find, err := m.repo.GetByUUID(message.UUID)
	if err != nil {
		return nil, err
	}
	if find.ID > 0 {
		return &pb.TextsReply{
			Id:   int64(find.ID),
			Uuid: message.UUID,
		}, nil
	}

	// process rule
	if m.bot != nil {
		m.bot.SetOptions(rule.Options...)
		out := m.bot.Process(message.Text).MessageProviderOut()
		if len(out) > 0 {
			return &pb.TextsReply{
				Text: out,
			}, nil
		}
	}

	// parse type
	message.Text = strings.TrimSpace(message.Text)
	if util.IsUrl(message.Text) {
		message.Type = model.MessageTypeLink
	}
	if message.IsMessageOfAction() {
		message.Type = model.MessageTypeAction
	}

	// store
	id, err := m.repo.Create(message)
	if err != nil {
		return nil, err
	}

	// trigger
	c := ctx.NewContext()
	c.Logger = m.logger
	c.Middle = m.middle
	c.Todo = m.todo
	triggers := trigger.Triggers()
	wg := sync.WaitGroup{}
	for _, item := range triggers {
		wg.Add(1)
		go func(t trigger.Trigger) {
			defer wg.Done()
			if t.Cond(message.Text) {
				t.Handle(c)
			}
		}(item)
	}
	wg.Wait()

	return &pb.TextsReply{
		Id:   id,
		Uuid: message.UUID,
	}, nil
}

func (m *Message) Delete(_ context.Context, payload *pb.MessageRequest) (*pb.TextReply, error) {
	err := m.repo.Delete(payload.GetId())
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (m *Message) Send(_ context.Context, payload *pb.MessageRequest) (*pb.StateReply, error) {
	client := http.NewClient()
	resp, err := client.PostJSON(m.config.Slack.Webhook, map[string]interface{}{
		"text": payload.GetText(),
	})
	if err != nil {
		return nil, err
	}

	_ = util.ByteToString(resp.Body())
	fasthttp.ReleaseResponse(resp)

	return &pb.StateReply{
		State: true,
	}, nil
}

func (m *Message) Run(ctx context.Context, payload *pb.MessageRequest) (*pb.TextReply, error) {
	var reply string
	message, err := m.repo.GetByID(payload.GetId())
	if err != nil {
		return nil, err
	}

	switch message.Type {
	case model.MessageTypeAction:
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

func (m *Message) GetActionMessages(_ context.Context, _ *pb.TextRequest) (*pb.ActionReply, error) {
	var items []model.Message
	items, err := m.repo.ListByType(model.MessageTypeAction)
	if err != nil {
		return nil, err
	}

	var kvs []*pb.Action
	for _, item := range items {
		kvs = append(kvs, &pb.Action{
			Id:   int64(item.ID),
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
		Type: model.MessageTypeAction,
	})
	if err != nil {
		return nil, err
	}

	// store message
	uuid, err := util.GenerateUUID()
	if err != nil {
		return nil, err
	}
	id, err := m.repo.Create(model.Message{
		UUID: uuid,
		Type: model.MessageTypeAction,
		Text: payload.GetText(),
	})
	if err != nil {
		return nil, err
	}

	// check/create trigger
	_, err = m.workflow.CreateTrigger(ctx, &pb.TriggerRequest{
		Kind:        model.MessageTypeAction,
		MessageId:   id,
		MessageText: payload.GetText(),
	})
	if err != nil {
		return nil, err
	}

	return &pb.StateReply{State: true}, nil
}

func (m *Message) DeleteWorkflowMessage(ctx context.Context, payload *pb.MessageRequest) (*pb.StateReply, error) {
	err := m.repo.Delete(payload.GetId())
	if err != nil {
		return nil, err
	}

	_, err = m.workflow.DeleteTrigger(ctx, &pb.TriggerRequest{MessageId: payload.GetId()})
	if err != nil {
		return nil, err
	}

	return &pb.StateReply{State: true}, nil
}
