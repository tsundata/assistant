package listener

import (
	"context"
	"encoding/json"
	"fmt"
	config2 "github.com/appleboy/gorush/config"
	"github.com/appleboy/gorush/core"
	"github.com/appleboy/gorush/notify"
	"github.com/go-redis/redis/v8"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/repository"
	"github.com/tsundata/assistant/internal/app/service"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
	"github.com/tsundata/assistant/internal/pkg/robot/rulebot"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc/md"
	"github.com/tsundata/assistant/internal/pkg/util"
	"go.uber.org/zap"
	"strings"
	"time"
)

func RegisterEventHandler(bus event.Bus, logger log.Logger, conf *config.AppConfig, rdb *redis.Client,
	bot *rulebot.RuleBot,
	comp component.Component,
	message service.MessageSvcClient,
	middle service.MiddleSvcClient,
	chatbot service.ChatbotSvcClient,
	storage pb.StorageSvcClient,
	messageRepo repository.MessageRepository,
	userRepo repository.UserRepository,
	chatbotRepo repository.ChatbotRepository) error {
	ctx := context.Background()
	err := bus.Subscribe(ctx, enum.Middle, event.CronRegisterSubject, func(msg *event.Msg) error {
		var m pb.CronRequest
		err := json.Unmarshal(msg.Data, &m)
		if err != nil {
			return err
		}

		_, err = middle.RegisterCron(ctx, &pb.CronRequest{Text: m.Text})
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	err = bus.Subscribe(ctx, enum.Middle, event.SubscribeRegisterSubject, func(msg *event.Msg) error {
		var m pb.SubscribeRequest
		err := json.Unmarshal(msg.Data, &m)
		if err != nil {
			return err
		}

		_, err = middle.RegisterSubscribe(ctx, &pb.SubscribeRequest{Text: m.Text})
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	err = bus.Subscribe(ctx, enum.Middle, event.CounterCreateSubject, func(msg *event.Msg) error {
		var m pb.Counter
		err := json.Unmarshal(msg.Data, &m)
		if err != nil {
			return err
		}

		ctx = md.BuildAuthContext(m.UserId)
		find, err := middle.GetCounterByFlag(ctx, &pb.CounterRequest{Counter: &m})
		if err != nil {
			return err
		}
		if find.Counter.Id > 0 {
			return nil
		}
		_, err = middle.CreateCounter(ctx, &pb.CounterRequest{Counter: &m})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	err = bus.Subscribe(ctx, enum.Middle, event.CounterIncreaseSubject, func(msg *event.Msg) error {
		var m pb.Counter
		err := json.Unmarshal(msg.Data, &m)
		if err != nil {
			return err
		}

		ctx = md.BuildAuthContext(m.UserId)
		_, err = middle.ChangeCounter(ctx, &pb.CounterRequest{Counter: &m})
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	err = bus.Subscribe(ctx, enum.Middle, event.CounterDecreaseSubject, func(msg *event.Msg) error {
		var m pb.Counter
		err := json.Unmarshal(msg.Data, &m)
		if err != nil {
			return err
		}

		ctx = md.BuildAuthContext(m.UserId)
		m.Digit = -m.Digit
		_, err = middle.ChangeCounter(ctx, &pb.CounterRequest{Counter: &m})
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	// message

	err = bus.Subscribe(ctx, enum.Message, event.EchoSubject, func(msg *event.Msg) error {
		fmt.Println(msg)
		if msg.Callback != nil {
			return bus.Publish(ctx, msg.Callback.Service, msg.Callback.Subject, pb.Message{Text: "echo"})
		}
		return nil
	})
	if err != nil {
		return err
	}

	err = bus.Subscribe(ctx, enum.Message, event.MessageSendSubject, func(msg *event.Msg) error {
		var m pb.Message
		err := json.Unmarshal(msg.Data, &m)
		if err != nil {
			return err
		}

		message := service.NewMessage(bus, logger, rdb, conf, messageRepo, chatbot, storage, middle)
		_, err = message.Send(md.BuildAuthContext(m.UserId), &pb.MessageRequest{Message: &m})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	err = bus.Subscribe(ctx, enum.Message, event.MessagePushSubject, func(msg *event.Msg) error {
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
			_, err := notify.SendNotification(&notification, &config2.ConfYaml{})
			if err != nil {
				logger.Error(err, zap.Any("event", event.MessagePushSubject))
			}
		}()
		return nil
	})
	if err != nil {
		return err
	}

	// chatbot

	err = bus.Subscribe(ctx, enum.Chatbot, event.BotHandleSubject, func(msg *event.Msg) error {
		var m pb.Message
		err := json.Unmarshal(msg.Data, &m)
		if err != nil {
			return err
		}

		chatbot := service.NewChatbot(logger, bus, rdb, chatbotRepo, message, middle, bot, comp)
		_, err = chatbot.Handle(md.BuildAuthContext(m.UserId), &pb.ChatbotRequest{MessageId: m.Id})
		return err
	})
	if err != nil {
		return err
	}

	err = bus.Subscribe(ctx, enum.Chatbot, event.BotRegisterSubject, func(msg *event.Msg) error {
		var b pb.Bot
		err := json.Unmarshal(msg.Data, &b)
		if err != nil {
			return err
		}

		chatbot := service.NewChatbot(logger, bus, rdb, chatbotRepo, message, middle, bot, comp)
		_, err = chatbot.Register(md.BuildAuthContext(enum.SuperUserID), &pb.BotRequest{Bot: &b})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	err = bus.Subscribe(ctx, enum.Chatbot, event.ScriptRunSubject, func(msg *event.Msg) error {
		var m pb.Message
		err := json.Unmarshal(msg.Data, &m)
		if err != nil {
			return err
		}

		switch enum.MessageType(m.GetType()) {
		case enum.MessageTypeScript:
			chatbot := service.NewChatbot(logger, bus, rdb, chatbotRepo, message, middle, bot, comp)
			reply, err := chatbot.RunActionScript(md.BuildAuthContext(m.UserId), &pb.WorkflowRequest{Message: &m})
			if err != nil {
				return err
			}
			if reply.GetText() != "" {
				uuid := util.UUID()
				rdb.Set(ctx, fmt.Sprintf("debug:%s", uuid), reply.GetText(), time.Hour)
				m.Text = fmt.Sprintf("DEBUG %s/debug/%s", conf.Gateway.Url, uuid)
				m.Type = string(enum.MessageTypeText)
				m.Direction = enum.MessageIncomingDirection
				m.SendTime = util.Now()
				_ = bus.Publish(ctx, enum.Message, event.MessageChannelSubject, m)
			}
			return nil
		}
		return nil
	})
	if err != nil {
		return err
	}

	err = bus.Subscribe(ctx, enum.Chatbot, event.BotActionSubject, func(msg *event.Msg) error {
		var m pb.Message
		err := json.Unmarshal(msg.Data, &m)
		if err != nil {
			return err
		}

		var p pb.ActionMsg
		err = json.Unmarshal(util.StringToByte(m.Payload), &p)
		if err != nil {
			return err
		}

		chatbot := service.NewChatbot(logger, bus, rdb, chatbotRepo, message, middle, bot, comp)
		_, err = chatbot.Action(md.BuildAuthContext(m.UserId), &pb.BotRequest{
			UserId:   m.UserId,
			GroupId:  m.GroupId,
			BotId:    m.Sender,
			ActionId: p.ID,
			Value:    p.Value,
		})
		return err
	})
	if err != nil {
		return err
	}

	err = bus.Subscribe(ctx, enum.Chatbot, event.BotFormSubject, func(msg *event.Msg) error {
		var m pb.Message
		err := json.Unmarshal(msg.Data, &m)
		if err != nil {
			return err
		}

		var p pb.FormMsg
		err = json.Unmarshal(util.StringToByte(m.Payload), &p)
		if err != nil {
			return err
		}

		var form []*pb.KV
		for _, item := range p.Field {
			if item.Value == nil {
				continue
			}
			form = append(form, &pb.KV{
				Key:   item.Key,
				Value: item.Value.(string),
			})
		}

		chatbot := service.NewChatbot(logger, bus, rdb, chatbotRepo, message, middle, bot, comp)
		_, err = chatbot.Form(md.BuildAuthContext(m.UserId), &pb.BotRequest{
			UserId:  m.UserId,
			GroupId: m.GroupId,
			BotId:   m.Sender,
			FormId:  p.ID,
			Form:    form,
		})
		return err
	})
	if err != nil {
		return err
	}

	// user

	err = bus.Subscribe(ctx, enum.User, event.RoleChangeExpSubject, func(msg *event.Msg) error {
		var role pb.Role
		err := json.Unmarshal(msg.Data, &role)
		if err != nil {
			return err
		}
		err = userRepo.ChangeRoleExp(ctx, role.UserId, role.Exp)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	err = bus.Subscribe(ctx, enum.User, event.RoleChangeAttrSubject, func(msg *event.Msg) error {
		var data pb.AttrChange
		err := json.Unmarshal(msg.Data, &data)
		if err != nil {
			return err
		}

		res, err := middle.Classifier(ctx, &pb.TextRequest{Text: data.Content})
		if err != nil {
			return err
		}

		err = userRepo.ChangeRoleAttr(ctx, data.UserId, res.Text, 1)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
