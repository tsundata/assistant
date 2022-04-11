package listener

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/chatbot/repository"
	"github.com/tsundata/assistant/internal/app/chatbot/service"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
	"github.com/tsundata/assistant/internal/pkg/robot/rulebot"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc/md"
	"github.com/tsundata/assistant/internal/pkg/util"
	"time"
)

func RegisterEventHandler(conf *config.AppConfig, bus event.Bus, rdb *redis.Client, logger log.Logger, bot *rulebot.RuleBot, message pb.MessageSvcClient,
	middle pb.MiddleSvcClient, repo repository.ChatbotRepository, comp component.Component) error {
	ctx := context.Background()

	err := bus.Subscribe(ctx, enum.Chatbot, event.BotHandleSubject, func(msg *event.Msg) error {
		var m pb.Message
		err := json.Unmarshal(msg.Data, &m)
		if err != nil {
			return err
		}

		chatbot := service.NewChatbot(logger, bus, rdb, repo, message, middle, bot, comp)
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

		chatbot := service.NewChatbot(logger, bus, rdb, repo, message, middle, bot, comp)
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
			chatbot := service.NewChatbot(logger, bus, rdb, repo, message, middle, bot, comp)
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

		chatbot := service.NewChatbot(logger, bus, rdb, repo, message, middle, bot, comp)
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

		chatbot := service.NewChatbot(logger, bus, rdb, repo, message, middle, bot, comp)
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

	return nil
}
