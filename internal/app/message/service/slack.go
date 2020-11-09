package service

import (
	"context"
	"github.com/tsundata/assistant/api/proto"
	"github.com/tsundata/assistant/internal/pkg/transports/http"
	"io/ioutil"
	"log"
)

type Slack struct {
	webhook string
}

func NewSlack(webhook string) *Slack {
	return &Slack{webhook: webhook}
}

func (s *Slack) SendMessage(message string, reply *string) error {
	client := http.NewClient()
	resp, err := client.PostJSON(s.webhook, map[string]interface{}{
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

func (s *Slack) Open(ctx context.Context, payload *proto.Detail, reply *proto.Detail) error {
	log.Println(payload)

	*reply = proto.Detail{
		Id:          1,
		Name:        "out =====>  slack",
		Price:       1000,
		CreatedTime: nil,
	}

	return nil
}
