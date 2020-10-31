package service

import (
	"github.com/spf13/viper"
	"github.com/tsundata/assistant/internal/pkg/transports/http"
	"io/ioutil"
)

type Slack struct {
	V *viper.Viper
}

func (s *Slack) SendMessage(message string, reply *string) error {
	slack := s.V.GetStringMap("slack")

	url := slack["webhook"].(string)
	client := http.NewClient()
	resp, err := client.PostJSON(url, map[string]interface{}{
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
