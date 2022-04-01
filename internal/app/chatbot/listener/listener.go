package listener

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/chatbot/repository"
	"github.com/tsundata/assistant/internal/app/chatbot/service"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
	"github.com/tsundata/assistant/internal/pkg/robot/rulebot"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc/md"
	"github.com/tsundata/assistant/internal/pkg/util"
)

func RegisterEventHandler(bus event.Bus, rdb *redis.Client, logger log.Logger, bot *rulebot.RuleBot, message pb.MessageSvcClient,
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

	err = bus.Subscribe(ctx, enum.Chatbot, event.WorkflowRunSubject, func(msg *event.Msg) error {
		var m pb.Message
		err := json.Unmarshal(msg.Data, &m)
		if err != nil {
			return err
		}

		switch enum.MessageType(m.GetType()) {
		case enum.MessageTypeScript:
			chatbot := service.NewChatbot(logger, bus, rdb, repo, message, middle, bot, comp)
			_, err = chatbot.RunActionScript(ctx, &pb.WorkflowRequest{Message: &m})
			return err
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
			form = append(form, &pb.KV{
				Key:   item.Key,
				Value: item.Value.(string),
			})
		}

		chatbot := service.NewChatbot(logger, bus, rdb, repo, message, middle, bot, comp)
		_, err = chatbot.Form(md.BuildAuthContext(m.UserId), &pb.BotRequest{
			UserId:   m.UserId,
			GroupId:  m.GroupId,
			BotId:    m.Sender,
			FormId: p.ID,
			Form:     form,
		})
		return err
	})
	if err != nil {
		return err
	}

	return nil
}
