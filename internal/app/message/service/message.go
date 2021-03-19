package service

import (
	"context"
	"database/sql"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/model"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
	"github.com/tsundata/assistant/internal/pkg/transports/http"
	"github.com/tsundata/assistant/internal/pkg/utils"
	"github.com/valyala/fasthttp"
	"strings"
	"time"
)

type Message struct {
	webhook  string
	db       *sqlx.DB
	logger   *logger.Logger
	bot      *rulebot.RuleBot
	wfClient pb.WorkflowClient
}

func NewManage(db *sqlx.DB, logger *logger.Logger, bot *rulebot.RuleBot, webhook string, wfClient pb.WorkflowClient) *Message {
	return &Message{db: db, logger: logger, bot: bot, webhook: webhook, wfClient: wfClient}
}

func (m *Message) List(_ context.Context, _ *pb.MessageRequest) (*pb.MessageListReply, error) {
	var messages []model.Message
	err := m.db.Select(&messages, "SELECT * FROM `messages` WHERE `type` <> ? AND `type` <> ? ORDER BY `id` DESC",
		model.MessageTypeAction, model.MessageTypeScript)
	if err != nil {
		return nil, err
	}

	var reply []*pb.MessageItem
	for _, item := range messages {
		reply = append(reply, &pb.MessageItem{
			Uuid: item.UUID,
			Text: item.Text,
			Type: item.Type,
			Time: item.Time.Format("2006-01-02 15:04:05"),
		})
	}

	return &pb.MessageListReply{
		Messages: reply,
	}, nil
}

func (m *Message) Get(_ context.Context, payload *pb.MessageRequest) (*pb.MessageReply, error) {
	var message model.Message
	err := m.db.Get(&message, "SELECT * FROM `messages` WHERE `id` = ? LIMIT 1", payload.GetId())
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return &pb.MessageReply{
		Id:   int64(message.ID),
		Uuid: message.UUID,
		Text: message.Text,
		Type: message.Type,
		Time: message.Time.Format("2006-01-02 15:04:05"),
	}, nil
}

func (m *Message) Create(_ context.Context, payload *pb.MessageRequest) (*pb.TextsReply, error) {
	// check uuid
	var message model.Message
	message.Time = time.Now()
	message.UUID = payload.GetUuid()
	message.Type = model.MessageTypeText
	message.Text = strings.TrimSpace(payload.GetText())

	// check
	var find model.Message
	err := m.db.Get(&find, "SELECT id FROM `messages` WHERE `uuid` = ? LIMIT 1", message.UUID)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if find.ID > 0 {
		return &pb.TextsReply{
			Uuid: message.UUID,
		}, nil
	}

	// process rule
	out := m.bot.Process(message.Text).MessageProviderOut()
	if len(out) > 0 {
		return &pb.TextsReply{
			Text: out,
		}, nil
	}

	// parse type
	message.Text = strings.TrimSpace(message.Text)
	if utils.IsUrl(message.Text) {
		message.Type = model.MessageTypeLink
	}
	if model.IsMessageOfAction(message.Text) {
		message.Type = model.MessageTypeAction
	}
	if model.IsMessageOfScript(message.Text) {
		message.Type = model.MessageTypeScript
	}

	// store
	res, err := m.db.NamedExec("INSERT INTO `messages` (`uuid`, `type`, `text`, `time`) VALUES (:uuid, :type, :text, :time)", message)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	// trigger
	//triggers := trigger.Triggers()
	//wg := sync.WaitGroup{}
	//for _, item := range triggers {
	//	wg.Add(1)
	//	go func(t trigger.Trigger) {
	//		defer wg.Done()
	//		if t.Cond(message.Text) {
	//			t.Handle()
	//		}
	//	}(item)
	//}
	//wg.Done()

	return &pb.TextsReply{
		Id:   id,
		Uuid: message.UUID,
	}, nil
}

func (m *Message) Delete(_ context.Context, payload *pb.MessageRequest) (*pb.TextReply, error) {
	_, err := m.db.Exec("DELETE FROM `messages` WHERE `id` = ?", payload.GetId())
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (m *Message) Send(_ context.Context, payload *pb.MessageRequest) (*pb.StateReply, error) {
	// TODO switch service
	client := http.NewClient()
	resp, err := client.PostJSON(m.webhook, map[string]interface{}{
		"text": payload.GetText(),
	})
	if err != nil {
		return nil, err
	}

	_ = utils.ByteToString(resp.Body())
	fasthttp.ReleaseResponse(resp)

	return &pb.StateReply{
		State: true,
	}, nil
}

func (m *Message) Run(ctx context.Context, payload *pb.MessageRequest) (*pb.TextReply, error) {
	var reply string
	var message model.Message
	err := m.db.Get(&message, "SELECT * FROM `messages` WHERE `id` = ? LIMIT 1", payload.GetId())
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	switch message.Type {
	case model.MessageTypeAction:
		wfReply, err := m.wfClient.RunAction(ctx, &pb.WorkflowRequest{Text: model.RemoveActionFlag(message.Text)})
		if err != nil {
			return nil, err
		}
		reply = wfReply.GetText()
	case model.MessageTypeScript:
		switch model.MessageScriptKind(message.Text) {
		case model.MessageScriptOfFlowscript:
			wfReply, err := m.wfClient.RunScript(ctx, &pb.WorkflowRequest{Text: message.Text})
			if err != nil {
				return nil, err
			}
			reply = wfReply.GetText()
		default:
			reply = model.MessageScriptOfUndefined
		}
	default:
		reply = "Not running"
	}

	return &pb.TextReply{
		Text: reply,
	}, nil
}

func (m *Message) GetScriptMessages(_ context.Context, _ *pb.TextRequest) (*pb.ScriptsReply, error) {
	var items []model.Message
	err := m.db.Select(&items, "SELECT * FROM `messages` WHERE `type` = ? ORDER BY `id` DESC", model.MessageTypeScript)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	var kvs []*pb.Script
	for _, item := range items {
		kvs = append(kvs, &pb.Script{
			Id:   int64(item.ID),
			Text: item.Text,
		})
	}

	return &pb.ScriptsReply{
		Items: kvs,
	}, nil
}

func (m *Message) CreateScriptMessage(ctx context.Context, payload *pb.TextRequest) (*pb.StateReply, error) {
	if payload.GetText() == "" {
		return &pb.StateReply{State: false}, nil
	}

	// check syntax
	_, err := m.wfClient.SyntaxCheck(ctx, &pb.WorkflowRequest{
		Text: payload.GetText(),
		Type: model.MessageTypeScript,
	})
	if err != nil {
		return nil, err
	}

	// store message
	uuid, err := utils.GenerateUUID()
	if err != nil {
		return nil, err
	}
	result, err := m.db.Exec("INSERT INTO `messages` (`uuid`, `type`, `text`, `time`) VALUES (?, ?, ?, ?)",
		uuid, model.MessageTypeScript, payload.GetText(), time.Now())
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	// check/store trigger
	_, err = m.wfClient.CreateTrigger(ctx, &pb.TriggerRequest{
		Kind:        model.MessageTypeScript,
		MessageId:   id,
		MessageText: payload.GetText(),
	})
	if err != nil {
		return nil, err
	}

	return &pb.StateReply{State: true}, nil
}

func (m *Message) GetActionMessages(_ context.Context, _ *pb.TextRequest) (*pb.ActionReply, error) {
	var items []model.Message
	err := m.db.Select(&items, "SELECT * FROM `messages` WHERE `type` = ? ORDER BY `id` DESC", model.MessageTypeAction)
	if err != nil && err != sql.ErrNoRows {
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
	_, err := m.wfClient.SyntaxCheck(ctx, &pb.WorkflowRequest{
		Text: payload.GetText(),
		Type: model.MessageTypeAction,
	})
	if err != nil {
		return nil, err
	}

	// store message
	uuid, err := utils.GenerateUUID()
	if err != nil {
		return nil, err
	}
	result, err := m.db.Exec("INSERT INTO `messages` (`uuid`, `type`, `text`, `time`) VALUES (?, ?, ?, ?)",
		uuid, model.MessageTypeAction, payload.GetText(), time.Now())
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	// check/create trigger
	_, err = m.wfClient.CreateTrigger(ctx, &pb.TriggerRequest{
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
	_, err := m.db.Exec("DELETE FROM `messages` WHERE `id` = ?", payload.GetId())
	if err != nil {
		return nil, err
	}

	_, err = m.wfClient.DeleteTrigger(ctx, &pb.TriggerRequest{MessageId: payload.GetId()})
	if err != nil {
		return nil, err
	}

	return &pb.StateReply{State: true}, nil
}

var ProviderSet = wire.NewSet(NewManage)
