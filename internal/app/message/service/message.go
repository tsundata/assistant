package service

import (
	"context"
	"github.com/robertkrimen/otto"
	"github.com/tsundata/assistant/internal/pkg/model"
	"github.com/tsundata/assistant/internal/pkg/transports/http"
	"github.com/tsundata/assistant/internal/pkg/utils"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strings"
)

type Message struct {
	webhook string
	db      *gorm.DB
	logger  *zap.Logger
}

func NewManage(db *gorm.DB, logger *zap.Logger) *Message {
	return &Message{db: db, logger: logger}
}

func (m *Message) List(ctx context.Context, payload *model.Message, reply *[]model.Message) error {
	var messages []model.Message
	m.db.Order("created_at DESC").Limit(10).Find(&messages)
	*reply = messages

	return nil
}

func (m *Message) View(ctx context.Context, payload *model.Message, reply *model.Message) error {
	var find model.Message
	m.db.Where("id = ?", payload.ID).Take(&find)
	*reply = find

	return nil
}

func (m *Message) Create(ctx context.Context, payload *model.Message, reply *model.Message) error {
	// check uuid
	var find model.Message
	m.db.Where("uuid = ?", payload.UUID).Take(&find)

	if find.ID > 0 {
		*reply = find
		return nil
	}

	// parse type
	payload.Content = strings.TrimSpace(payload.Content)
	if utils.IsUrl(payload.Content) {
		payload.Type = model.MessageTypeUrl
	}
	if utils.IsMessageOfAction(payload.Content) {
		payload.Type = model.MessageTypeAction
	}
	if utils.IsMessageOfScript(payload.Content) {
		payload.Type = model.MessageTypeScript
	}

	// insert
	m.db.Create(payload)
	*reply = *payload

	return nil
}

func (m *Message) Delete(ctx context.Context, payload *model.Message, reply *model.Message) error {
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
	var find model.Message
	m.db.Where("uuid = ?", payload).Take(&find)

	if find.ID == 0 {
		*reply = "Not message"
		return nil
	}

	// TODO
	switch find.Type {
	case model.MessageTypeAction:
	case model.MessageTypeScript:
		switch utils.MessageScriptKind(find.Content) {
		case model.MessageScriptOfDSL:
		case model.MessageScriptOfJavascript:
			// TODO
			vm := otto.New()
			v, err := vm.Run(strings.Replace(find.Content, "#!script:javascript", "", -1))
			if err != nil {
				m.logger.Error(err.Error())
				return err
			}
			*reply = v.String()
		case model.MessageScriptOfUndefined:
		default:
		}
	default:
		*reply = "Not running"
	}

	return nil
}
