package service

import (
	"context"
	"github.com/tsundata/assistant/api/proto"
	"github.com/tsundata/assistant/internal/pkg/models"
	"github.com/tsundata/assistant/internal/pkg/transports/http"
	"gorm.io/gorm"
	"io/ioutil"
)

type Message struct {
	// TODO load config
	webhook string
	db      *gorm.DB
}

func NewManage(db *gorm.DB) *Message {
	return &Message{db: db}
}

func (m *Message) List(ctx context.Context, payload *proto.Message, reply *[]proto.Message) error {
	var messages []models.Message
	m.db.Order("created_at DESC").Limit(10).Find(&messages)

	for _, item := range messages {
		*reply = append(*reply, proto.Message{
			Id:     uint64(item.ID),
			Uuid:   item.UUID,
			Input:  item.Input,
			Output: item.Output,
		})
	}

	return nil
}

func (m *Message) View(ctx context.Context, payload *proto.Message, reply *proto.Message) error {
	var find models.Message
	m.db.Where("id = ?", payload.Id).Take(&find)

	*reply = proto.Message{
		Id:     uint64(find.ID),
		Uuid:   find.UUID,
		Input:  find.Input,
		Output: find.Output,
	}
	return nil
}

func (m *Message) Create(ctx context.Context, payload *proto.Message, reply *proto.Message) error {
	// check uuid
	var find models.Message
	m.db.Where("uuid = ?", payload.Uuid).Take(&find)

	if find.ID > 0 {
		*reply = proto.Message{
			Id:   uint64(find.ID),
			Uuid: find.UUID,
		}
		return nil
	}

	// insert
	msg := models.NewMessage()
	msg.UUID = payload.Uuid
	msg.Input = payload.Input
	msg.Output = payload.Output
	msg.Remotes = models.RemoteType(payload.Remotes)
	m.db.Create(&msg)

	*reply = proto.Message{
		Id:   uint64(msg.ID),
		Uuid: msg.UUID,
	}

	return nil
}

func (m *Message) Delete(ctx context.Context, payload *proto.Message, reply *proto.Message) error {
	m.db.Where("id = ?", payload.Id).Delete(models.Message{})
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
