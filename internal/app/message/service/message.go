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

func (m *Message) List(ctx context.Context, payload *pb.MessageRequest) (*pb.MessageListReply, error) {
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

func (m *Message) Get(ctx context.Context, payload *pb.MessageRequest) (*pb.TextReply, error) {
	var message model.Message
	err := m.db.Get(&message, "SELECT text FROM `messages` WHERE `uuid` = ? LIMIT 1", payload.Uuid)
	if err != nil {
		return nil, err
	}

	return &pb.TextReply{
		Text: message.Text,
	}, nil
}

func (m *Message) Create(ctx context.Context, in *pb.MessageRequest) (*pb.MessageReply, error) {
	// check uuid
	var payload model.Message
	payload.Time = time.Now()
	payload.UUID = in.GetUuid()
	payload.Type = model.MessageTypeText
	payload.Text = strings.TrimSpace(in.GetText())

	// check
	var find model.Message
	err := m.db.Get(&find, "SELECT id FROM `messages` WHERE `uuid` = ? LIMIT 1", payload.UUID)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if find.ID > 0 {
		return &pb.MessageReply{
			Uuid: payload.UUID,
		}, nil
	}

	// parse type
	payload.Text = strings.TrimSpace(payload.Text)
	if utils.IsUrl(payload.Text) {
		payload.Type = model.MessageTypeLink
	}
	if model.IsMessageOfAction(payload.Text) {
		payload.Type = model.MessageTypeAction
	}
	if model.IsMessageOfScript(payload.Text) {
		payload.Type = model.MessageTypeScript
	}

	if payload.Type == model.MessageTypeText {
		out := m.bot.Process(payload).MessageProviderOut()
		if len(out) > 0 {
			var reply []string
			for _, item := range out {
				reply = append(reply, item.Text)
			}
			return &pb.MessageReply{
				Text: reply,
			}, nil
		}
	}

	// insert
	res, err := m.db.NamedExec("INSERT INTO `messages` (`uuid`, `type`, `text`, `time`) VALUES (:uuid, :type, :text, :time)", payload)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &pb.MessageReply{
		Id:   id,
		Uuid: payload.UUID,
	}, nil
}

func (m *Message) Delete(ctx context.Context, payload *pb.MessageRequest) (*pb.TextReply, error) {
	_, err := m.db.Exec("DELETE FROM `messages` WHERE `uuid` = ?", payload.Uuid)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (m *Message) Send(ctx context.Context, payload *pb.MessageRequest) (*pb.TextReply, error) {
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

func (m *Message) Run(ctx context.Context, in *pb.MessageRequest) (*pb.TextReply, error) {
	// check uuid
	var reply string
	var payload model.Message

	err := m.db.Get(&payload, "SELECT id FROM `messages` WHERE `uuid` = ? LIMIT 1", in.Uuid)
	if err != nil {
		return nil, err
	}

	if payload.UUID == "" {
		return &pb.TextReply{
			Text: "Not message",
		}, nil
	}

	switch payload.Text {
	case model.MessageTypeAction:
		// TODO action
	case model.MessageTypeScript:
		switch model.MessageScriptKind(payload.Text) {
		case model.MessageScriptOfFlowscript:
			txt := strings.ReplaceAll(payload.Text, "#!script:flowscript", "")
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
