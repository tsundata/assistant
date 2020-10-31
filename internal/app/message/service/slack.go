package service

import (
	"github.com/tsundata/assistant/internal/pkg/transports/http"
	"io/ioutil"
)

type Slack struct {
	Webhook string
}

func (s *Slack) SendMessage(message string, reply *string) error {
	client := http.NewClient()
	resp, err := client.PostJSON(s.Webhook, map[string]interface{}{
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
