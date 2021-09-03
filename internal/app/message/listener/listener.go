package listener

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/appleboy/gorush/core"
	"github.com/appleboy/gorush/notify"
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
	"strings"
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

	err = bus.Subscribe(context.Background(), event.MessagePushSubject, func(msg *nats.Msg) {
		var in pb.Notification
		err := json.Unmarshal(msg.Data, &in)
		if err != nil {
			logger.Error(err, zap.Any("event", event.MessagePushSubject))
			return
		}

		badge := int(in.Badge)
		notification := notify.PushNotification{
			Platform:         int(in.Platform),
			Tokens:           in.Tokens,
			Message:          in.Message,
			Title:            in.Title,
			Topic:            in.Topic,
			APIKey:           in.Key,
			Category:         in.Category,
			Sound:            in.Sound,
			ContentAvailable: in.ContentAvailable,
			ThreadID:         in.ThreadID,
			MutableContent:   in.MutableContent,
			Image:            in.Image,
			Priority:         strings.ToLower(in.GetPriority().String()),
		}

		if badge > 0 {
			notification.Badge = &badge
		}

		if in.Topic != "" && in.Platform == core.PlatFormAndroid {
			notification.To = in.Topic
		}

		if in.Alert != nil {
			notification.Alert = notify.Alert{
				Title:        in.Alert.Title,
				Body:         in.Alert.Body,
				Subtitle:     in.Alert.Subtitle,
				Action:       in.Alert.Action,
				ActionLocKey: in.Alert.Action,
				LaunchImage:  in.Alert.LaunchImage,
				LocArgs:      in.Alert.LocArgs,
				LocKey:       in.Alert.LocKey,
				TitleLocArgs: in.Alert.TitleLocArgs,
				TitleLocKey:  in.Alert.TitleLocKey,
			}
		}

		go func() {
			_, err := notify.SendNotification(&notification, &config.ConfYaml)
			if err != nil {
				logger.Error(err, zap.Any("event", event.MessagePushSubject))
			}
		}()
	})

	return nil
}
