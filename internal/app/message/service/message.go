package service

import (
	"context"
	"github.com/robertkrimen/otto"
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
	webhook string
	db      *gorm.DB
	logger  *zap.Logger
	bot     *bot.Bot
}

func NewManage(db *gorm.DB, logger *zap.Logger, bot *bot.Bot) *Message {
	return &Message{db: db, logger: logger, bot: bot}
}

func (m *Message) List(ctx context.Context, payload *model.Event, reply *[]model.Event) error {
	var messages []model.Event
	m.db.Order("created_at DESC").Limit(10).Find(&messages)
	*reply = messages

	return nil
}

func (m *Message) View(ctx context.Context, payload *model.Event, reply *model.Event) error {
	var find model.Event
	m.db.Where("id = ?", payload.ID).Take(&find)
	*reply = find

	return nil
}

func (m *Message) Create(ctx context.Context, payload *model.Event, reply *[]model.Event) error {
	// check uuid
	var find model.Event
	m.db.Where("uuid = ?", payload.UUID).Take(&find)

	if find.ID > 0 {
		*reply = []model.Event{find}
		return nil
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
		out := m.bot.Process(*payload).MessageProviderOut()
		if len(out) > 0 {
			*reply = out
			return nil
		}
	}

	payload.Time = time.Now()

	// insert
	m.db.Create(&payload)
	*reply = []model.Event{*payload}

	return nil
}

func (m *Message) Delete(ctx context.Context, payload *model.Event, reply *model.Event) error {
	m.db.Where("id = ?", payload.ID).Delete(model.Message{})
	return nil
}

func (m *Message) SendMessage(ctx context.Context, message string, reply *string) error {
	// TODO switch service
	client := http.NewClient()
	resp, err := client.PostJSON(m.webhook, map[string]interface{}{
		"text": message,
	})
	if err != nil {
		return err
	}

	*reply = string(resp.Body())
	fasthttp.ReleaseResponse(resp)

	return nil
}

func (m *Message) Run(ctx context.Context, payload string, reply *string) error {
	// check uuid
	var find model.Event
	m.db.Where("uuid = ?", payload).Take(&find)

	if find.ID == 0 {
		*reply = "Not message"
		return nil
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
				return err
			}
			tree, err := p.Parse()
			if err != nil {
				m.logger.Error(err.Error())
				return err
			}
			i := interpreter.NewInterpreter(tree)
			_, err = i.Interpret()
			if err != nil {
				m.logger.Error(err.Error())
				return err
			}
			*reply = i.Stdout()
		case model.MessageScriptOfJavascript:
			vm := otto.New()
			v, err := vm.Run(strings.Replace(find.Data.Message.Text, "#!script:javascript", "", -1))
			if err != nil {
				m.logger.Error(err.Error())
				return err
			}
			*reply = v.String()
		case model.MessageScriptOfUndefined:
			*reply = "MessageScriptOfUndefined"
		default:
			*reply = "MessageScriptOfUndefined"
		}
	default:
		*reply = "Not running"
	}

	return nil
}
