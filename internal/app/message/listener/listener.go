package listener

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/transport/http"
	"github.com/tsundata/assistant/internal/pkg/util"
	"github.com/tsundata/assistant/internal/pkg/vendors/slack"
	"github.com/valyala/fasthttp"
)

func RegisterEventHandler(bus event.Bus, config *config.AppConfig, logger log.Logger) error {
	err := bus.Subscribe(context.Background(), event.EchoSubject, func(msg *nats.Msg) {
		fmt.Println(msg)
	})
	if err != nil {
		return err
	}

	err = bus.Subscribe(context.Background(), event.SendMessageSubject, func(msg *nats.Msg) {
		var message pb.Message
		err := json.Unmarshal(msg.Data, &message)
		if err != nil {
			logger.Error(err)
			return
		}

		client := http.NewClient()
		webhook := slack.ChannelSelect(message.Channel, config.Slack.Webhook)
		resp, err := client.PostJSON(webhook, map[string]interface{}{
			"text": message.Text,
		})
		if err != nil {
			logger.Error(err)
			return
		}

		_ = util.ByteToString(resp.Body())
		fasthttp.ReleaseResponse(resp)
	})
	if err != nil {
		return err
	}

	return nil
}
