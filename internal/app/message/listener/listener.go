package listener

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/appleboy/gorush/core"
	"github.com/appleboy/gorush/notify"
	"github.com/go-redis/redis/v8"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/message/repository"
	"github.com/tsundata/assistant/internal/app/message/service"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/log"
	"go.uber.org/zap"
	"strings"
)

func RegisterEventHandler(bus event.Bus, config *config.AppConfig, logger log.Logger, redis *redis.Client,
	repo repository.MessageRepository, chatbot pb.ChatbotSvcClient, storage pb.StorageSvcClient, middle pb.MiddleSvcClient) error {
	err := bus.Subscribe(context.Background(), enum.Message, event.EchoSubject, func(msg *event.Msg) error {
		fmt.Println(msg)
		if msg.Callback != nil {
			return bus.Publish(context.Background(), msg.Callback.Service, msg.Callback.Subject, pb.Message{Text: "echo"})
		}
		return nil
	})
	if err != nil {
		return err
	}

	err = bus.Subscribe(context.Background(), enum.Message, event.MessageSendSubject, func(msg *event.Msg) error {
		var m pb.Message
		err := json.Unmarshal(msg.Data, &m)
		if err != nil {
			return err
		}

		message := service.NewMessage(bus, logger, redis, config, repo, chatbot, storage, middle)
		_, err = message.Send(context.Background(), &pb.MessageRequest{Message: &m})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	err = bus.Subscribe(context.Background(), enum.Message, event.MessagePushSubject, func(msg *event.Msg) error {
		var in pb.Notification
		err := json.Unmarshal(msg.Data, &in)
		if err != nil {
			return err
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
			ThreadID:         in.ThreadId,
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
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
