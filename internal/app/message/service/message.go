package service

import (
	"context"
	"github.com/tsundata/assistant/api/proto"
	"github.com/tsundata/assistant/internal/pkg/transports/http"
	"io/ioutil"
)

type Message struct {
	// TODO load config
	webhook string
}

func NewManage() *Message {
	return &Message{}
}

// TODO
func (m *Message) List(ctx context.Context, payload *proto.Detail, reply *proto.Detail) error {
	return nil
}

// TODO
func (m *Message) View(ctx context.Context, payload *proto.Detail, reply *proto.Detail) error {
	return nil
}

// TODO
func (m *Message) Create(ctx context.Context, payload *proto.Detail, reply *proto.Detail) error {
	return nil
}

// TODO
func (m *Message) Delete(ctx context.Context, payload *proto.Detail, reply *proto.Detail) error {
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
