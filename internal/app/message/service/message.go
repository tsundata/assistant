package service

import (
	"context"
	"github.com/robertkrimen/otto"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/message/bot"
	"github.com/tsundata/assistant/internal/pkg/interpreter"
	"github.com/tsundata/assistant/internal/pkg/model"
	"github.com/tsundata/assistant/internal/pkg/transports/http"
	"github.com/tsundata/assistant/internal/pkg/utils"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strings"
	"time"
)

type Message struct {
	db      *gorm.DB
	logger  *zap.Logger
	bot     *bot.Bot
	webhook string
}

func NewManage(db *gorm.DB, logger *zap.Logger, bot *bot.Bot, webhook string) *Message {
	return &Message{db: db, logger: logger, bot: bot, webhook: webhook}
}

func (m *Message) List(ctx context.Context, payload *pb.MessageRequest) (*pb.MessageList, error) {
	var messages []model.Event
	m.db.Order("created_at DESC").Limit(10).Find(&messages)

	var reply []string
	for _, item := range messages {
		reply = append(reply, item.Data.Message.Text)
	}
	return &pb.MessageList{
		Text: reply,
	}, nil
}

func (m *Message) Get(ctx context.Context, payload *pb.MessageRequest) (*pb.MessageReply, error) {
	var find model.Event
	m.db.Where("id = ?", 1).Take(&find)

	return &pb.MessageReply{
		Text: find.Data.Message.Text,
	}, nil
}

func (m *Message) Create(ctx context.Context, in *pb.MessageRequest) (*pb.MessageList, error) {
	// check uuid
	var find model.Event
	var payload model.Event
	payload.UUID = in.GetUuid()
	payload.Data.Message.Type = model.MessageTypeText
	payload.Data.Message.Text = in.GetText()
	m.db.Where("uuid = ?", in.GetUuid()).Take(&find)

	if find.ID > 0 {
		return &pb.MessageList{
			Id: int64(find.ID),
		}, nil
	}

	// parse type
	payload.Data.Message.Text = strings.TrimSpace(payload.Data.Message.Text)
	if utils.IsUrl(payload.Data.Message.Text) {
		payload.Data.Message.Type = model.MessageTypeLink
	}
	if utils.IsMessageOfAction(payload.Data.Message.Text) {
		payload.Data.Message.Type = model.MessageTypeAction
	}
	if utils.IsMessageOfScript(payload.Data.Message.Text) {
		payload.Data.Message.Type = model.MessageTypeScript
	}

	if payload.Data.Message.Type == model.MessageTypeText {
		out := m.bot.Process(payload).MessageProviderOut()
		if len(out) > 0 {
			var reply []string
			for _, item := range out {
				reply = append(reply, item.Data.Message.Text)
			}
			return &pb.MessageList{
				Text: reply,
			}, nil
		}
	}

	payload.Time = time.Now()

	// insert
	m.db.Create(&payload)

	return &pb.MessageList{
		Id: int64(payload.ID),
	}, nil
}

func (m *Message) Delete(ctx context.Context, payload *pb.MessageRequest) (*pb.MessageReply, error) {
	m.db.Where("id = ?", 1).Delete(model.Message{})
	return nil, nil
}

func (m *Message) Send(ctx context.Context, payload *pb.MessageRequest) (*pb.MessageReply, error) {
	// TODO switch service
	client := http.NewClient()
	resp, err := client.PostJSON(m.webhook, map[string]interface{}{
		"text": payload.GetText(),
	})
	if err != nil {
		return nil, err
	}

	reply := string(resp.Body())
	fasthttp.ReleaseResponse(resp)

	return &pb.MessageReply{
		Text: reply,
	}, nil
}

func (m *Message) Run(ctx context.Context, in *pb.MessageRequest) (*pb.MessageReply, error) {
	// check uuid
	var reply string
	var find model.Event
	var payload model.Event
	m.db.Where("uuid = ?", payload).Take(&find)

	if find.ID == 0 {
		//	*reply = "Not message"
		return nil, nil
	}

	switch find.Data.Message.Text {
	case model.MessageTypeAction:
		// TODO action
	case model.MessageTypeScript:
		switch utils.MessageScriptKind(find.Data.Message.Text) {
		case model.MessageScriptOfFlowscript:
			text := strings.Replace(find.Data.Message.Text, "#!script:flowscript", "", -1)
			p, err := interpreter.NewParser(interpreter.NewLexer([]rune(text)))
			if err != nil {
				m.logger.Error(err.Error())
				return nil, err
			}
			tree, err := p.Parse()
			if err != nil {
				m.logger.Error(err.Error())
				return nil, err
			}
			i := interpreter.NewInterpreter(tree)
			_, err = i.Interpret()
			if err != nil {
				m.logger.Error(err.Error())
				return nil, err
			}
			reply = i.Stdout()
		case model.MessageScriptOfJavascript:
			vm := otto.New()
			v, err := vm.Run(strings.Replace(find.Data.Message.Text, "#!script:javascript", "", -1))
			if err != nil {
				m.logger.Error(err.Error())
				return nil, err
			}
			reply = v.String()
		case model.MessageScriptOfUndefined:
			reply = "MessageScriptOfUndefined"
		default:
			reply = "MessageScriptOfUndefined"
		}
	default:
		reply = "Not running"
	}

	return &pb.MessageReply{
		Text: reply,
	}, nil
}
