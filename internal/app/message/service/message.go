package service

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/model"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
	"github.com/tsundata/assistant/internal/pkg/transports/http"
	"github.com/tsundata/assistant/internal/pkg/utils"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"strings"
	"time"
)

type Message struct {
	webhook  string
	db       *sqlx.DB
	logger   *zap.Logger
	bot      *rulebot.RuleBot
	wfClient pb.WorkflowClient
}

func NewManage(db *sqlx.DB, logger *zap.Logger, bot *rulebot.RuleBot, webhook string, wfClient pb.WorkflowClient) *Message {
	return &Message{db: db, logger: logger, bot: bot, webhook: webhook, wfClient: wfClient}
}

func (m *Message) List(_ context.Context, _ *pb.MessageRequest) (*pb.MessageListReply, error) {
	var messages []model.Message
	err := m.db.Select(&messages, "SELECT * FROM `messages` ORDER BY `id` DESC")
	if err != nil {
		return nil, err
	}

	var reply []*pb.MessageItem
	for _, item := range messages {
		reply = append(reply, &pb.MessageItem{
			Uuid: item.UUID,
			Text: item.Text,
			Time: item.Time.Format("2006-01-02 15:04:05"),
		})
	}

	return &pb.MessageListReply{
		Messages: reply,
	}, nil
}

func (m *Message) Get(_ context.Context, payload *pb.MessageRequest) (*pb.TextReply, error) {
	var message model.Message
	err := m.db.Get(&message, "SELECT text FROM `messages` WHERE `uuid` = ? LIMIT 1", payload.GetUuid())
	if err != nil {
		return nil, err
	}

	return &pb.TextReply{
		Text: message.Text,
	}, nil
}

func (m *Message) Create(_ context.Context, payload *pb.MessageRequest) (*pb.MessageReply, error) {
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
		return &pb.MessageReply{
			Uuid: message.UUID,
		}, nil
	}

	// process rule
	out := m.bot.Process(message).MessageProviderOut()
	if len(out) > 0 {
		var reply []string
		for _, item := range out {
			reply = append(reply, item.Text)
		}
		return &pb.MessageReply{
			Text: reply,
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

	return &pb.MessageReply{
		Id:   id,
		Uuid: message.UUID,
	}, nil
}

func (m *Message) Delete(_ context.Context, payload *pb.MessageRequest) (*pb.TextReply, error) {
	_, err := m.db.Exec("DELETE FROM `messages` WHERE `uuid` = ?", payload.GetUuid())
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (m *Message) Send(_ context.Context, payload *pb.MessageRequest) (*pb.TextReply, error) {
	// TODO switch service
	client := http.NewClient()
	resp, err := client.PostJSON(m.webhook, map[string]interface{}{
		"text": payload.GetText(),
	})
	if err != nil {
		return nil, err
	}

	reply := utils.ByteToString(resp.Body())
	fasthttp.ReleaseResponse(resp)

	return &pb.TextReply{
		Text: reply,
	}, nil
}

func (m *Message) Run(_ context.Context, payload *pb.MessageRequest) (*pb.TextReply, error) {
	// check uuid
	var reply string
	var message model.Message

	err := m.db.Get(&message, "SELECT id FROM `messages` WHERE `uuid` = ? LIMIT 1", payload.GetUuid())
	if err != nil {
		return nil, err
	}

	if message.UUID == "" {
		return &pb.TextReply{
			Text: "Not message",
		}, nil
	}

	switch message.Text {
	case model.MessageTypeAction:
		// TODO action
	case model.MessageTypeScript:
		switch model.MessageScriptKind(message.Text) {
		case model.MessageScriptOfFlowscript:
			txt := strings.ReplaceAll(message.Text, "#!script:flowscript", "")
			r, err := m.wfClient.Run(context.Background(), &pb.WorkflowRequest{
				Text: txt,
			})
			reply = "run error"
			if err != nil {
				reply = err.Error()
			}
			if r != nil {
				reply = r.Text
			}
		case model.MessageScriptOfUndefined:
			reply = "MessageScriptOfUndefined"
		default:
			reply = "MessageScriptOfUndefined"
		}
	default:
		reply = "Not running"
	}

	return &pb.TextReply{
		Text: reply,
	}, nil
}
