package service

import (
	"context"
	"github.com/tsundata/assistant/internal/pkg/model"
	"github.com/tsundata/assistant/internal/pkg/transports/http"
	"gorm.io/gorm"
	"io/ioutil"
)

type Message struct {
	webhook string
	db      *gorm.DB
}

func NewManage(db *gorm.DB) *Message {
	return &Message{db: db}
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
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	*reply = string(body)

	return nil
}
