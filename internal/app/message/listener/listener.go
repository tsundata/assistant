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
	"go.uber.org/zap"
)

func RegisterEventHandler(bus event.Bus, config *config.AppConfig, logger log.Logger) error {
	err := bus.Subscribe(context.Background(), event.EchoSubject, func(msg *nats.Msg) {
		fmt.Println(msg)
	})
	if err != nil {
		return err
	}

	err = bus.Subscribe(context.Background(), event.MessageSendSubject, func(msg *nats.Msg) {
		var m pb.Message
		err := json.Unmarshal(msg.Data, &m)
		if err != nil {
			logger.Error(err, zap.Any("event", event.MessageSendSubject))
			return
		}

		client := http.NewClient()
		webhook := slack.ChannelSelect(m.Channel, config.Slack.Webhook)
		resp, err := client.PostJSON(webhook, map[string]interface{}{
			"text": m.Text,
		})
		if err != nil {
			logger.Error(err, zap.Any("event", event.MessageSendSubject))
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
