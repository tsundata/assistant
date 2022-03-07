package listener

import (
	"context"
	"encoding/json"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/chatbot/repository"
	"github.com/tsundata/assistant/internal/app/chatbot/service"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
	"github.com/tsundata/assistant/internal/pkg/robot/rulebot"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc/md"
)

func RegisterEventHandler(bus event.Bus, logger log.Logger, bot *rulebot.RuleBot, message pb.MessageSvcClient,
	middle pb.MiddleSvcClient, repo repository.ChatbotRepository, comp component.Component) error {
	ctx := context.Background()

	err := bus.Subscribe(ctx, enum.Chatbot, event.MessageHandleSubject, func(msg *event.Msg) error {
		var m pb.Message
		err := json.Unmarshal(msg.Data, &m)
		if err != nil {
			return err
		}

		chatbot := service.NewChatbot(logger, bus, repo, message, middle, bot, comp)
		_, err = chatbot.Handle(md.BuildAuthContext(m.UserId), &pb.ChatbotRequest{MessageId: m.Id})
		if err != nil {
			return err
		}
		return nil
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

		chatbot := service.NewChatbot(logger, bus, repo, message, middle, bot, comp)
		_, err = chatbot.Register(ctx, &pb.BotRequest{Bot: &b})
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
