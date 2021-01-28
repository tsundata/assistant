package service

import (
	"context"
	"encoding/json"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/model"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
	"github.com/tsundata/assistant/internal/pkg/transports/http"
	"github.com/tsundata/assistant/internal/pkg/utils"
	"github.com/valyala/fasthttp"
	"go.etcd.io/bbolt"
	"go.uber.org/zap"
	"strings"
	"time"
)

type Message struct {
	webhook  string
	db       *bbolt.DB
	logger   *zap.Logger
	bot      *rulebot.RuleBot
	wfClient pb.WorkflowClient
}

func NewManage(db *bbolt.DB, logger *zap.Logger, bot *rulebot.RuleBot, webhook string, wfClient pb.WorkflowClient) *Message {
	return &Message{db: db, logger: logger, bot: bot, webhook: webhook, wfClient: wfClient}
}

func (m *Message) List(ctx context.Context, payload *pb.MessageRequest) (*pb.MessageListReply, error) {
	tx, err := m.db.Begin(true)
	if err != nil {
		return nil, err
	}
	b := tx.Bucket(utils.StringToByte("message"))
	c := b.Cursor()
	limit := 20

	index := 0
	var reply []*pb.MessageItem
	for k, v := c.Last(); k != nil; k, v = c.Prev() {
		index++

		var m model.Message
		err := json.Unmarshal(v, &m)
		if err != nil {
			return nil, err
		}

		reply = append(reply, &pb.MessageItem{
			Uuid: m.UUID,
			Text: m.Text,
			Time: m.Time,
		})

		if index >= limit {
			break
		}
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &pb.MessageListReply{
		Messages: reply,
	}, nil
}

func (m *Message) Get(ctx context.Context, payload *pb.MessageRequest) (*pb.TextReply, error) {
	tx, err := m.db.Begin(true)
	if err != nil {
		return nil, err
	}
	b := tx.Bucket(utils.StringToByte("message"))
	v := b.Get(utils.StringToByte(payload.Uuid))
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	var find model.Message
	err = json.Unmarshal(v, &find)
	if err != nil {
		return nil, err
	}

	return &pb.TextReply{
		Text: find.Text,
	}, nil
}

func (m *Message) Create(ctx context.Context, in *pb.MessageRequest) (*pb.MessageReply, error) {
	// check uuid
	var payload model.Message
	payload.Time = time.Now().Format("2006-01-02 15:04:05")
	payload.UUID = in.GetUuid()
	payload.Type = model.MessageTypeText
	payload.Text = strings.TrimSpace(in.GetText())

	// check
	tx, err := m.db.Begin(true)
	if err != nil {
		return nil, err
	}
	b, err := tx.CreateBucketIfNotExists(utils.StringToByte("message"))
	if err != nil {
		return nil, err
	}
	v := b.Get(utils.StringToByte(payload.UUID))
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	if v != nil {
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
	err = m.db.Update(func(tx *bbolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(utils.StringToByte("message"))
		if err != nil {
			return err
		}
		data, err := json.Marshal(payload)
		if err != nil {
			return err
		}
		return b.Put(utils.StringToByte(payload.UUID), data)
	})
	if err != nil {
		return nil, err
	}

	return &pb.MessageReply{
		Uuid: payload.UUID,
	}, nil
}

func (m *Message) Delete(ctx context.Context, payload *pb.MessageRequest) (*pb.TextReply, error) {
	tx, err := m.db.Begin(true)
	if err != nil {
		return nil, err
	}
	b := tx.Bucket(utils.StringToByte("message"))
	err = b.Delete(utils.StringToByte(payload.Uuid))
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
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

	tx, err := m.db.Begin(false)
	if err != nil {
		return nil, err
	}
	b := tx.Bucket(utils.StringToByte("message"))
	v := b.Get(utils.StringToByte(payload.UUID))
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	var find model.Message
	err = json.Unmarshal(v, &find)
	if err != nil {
		return nil, err
	}

	if find.UUID == "" {
		return &pb.TextReply{
			Text: "Not message",
		}, nil
	}

	switch find.Text {
	case model.MessageTypeAction:
		// TODO action
	case model.MessageTypeScript:
		switch model.MessageScriptKind(find.Text) {
		case model.MessageScriptOfFlowscript:
			txt := strings.ReplaceAll(find.Text, "#!script:flowscript", "")
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
