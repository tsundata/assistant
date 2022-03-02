package listener

import (
	"context"
	"encoding/json"
	"github.com/nats-io/nats.go"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/chatbot/repository"
	"github.com/tsundata/assistant/internal/app/chatbot/service"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/robot/rulebot"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc/md"
	"go.uber.org/zap"
)

func RegisterEventHandler(bus event.Bus, logger log.Logger, bot *rulebot.RuleBot, message pb.MessageSvcClient,
	repo repository.ChatbotRepository) error {
	ctx := context.Background()

	err := bus.Subscribe(ctx, event.MessageHandleSubject, func(msg *nats.Msg) {
		var m pb.Message
		err := json.Unmarshal(msg.Data, &m)
		if err != nil {
			logger.Error(err, zap.Any("event", event.MessageHandleSubject))
			return
		}

		chatbot := service.NewChatbot(logger, bus, repo, message, bot)
		_, err = chatbot.Handle(md.BuildAuthContext(m.UserId), &pb.ChatbotRequest{MessageId: m.Id})
		if err != nil {
			logger.Error(err, zap.Any("event", event.MessageHandleSubject))
			return
		}
	})
	if err != nil {
		return err
	}

	err = bus.Subscribe(ctx, event.BotRegisterSubject, func(msg *nats.Msg) {
		var b pb.Bot
		err := json.Unmarshal(msg.Data, &b)
		if err != nil {
			logger.Error(err, zap.Any("event", event.BotRegisterSubject))
			return
		}

		chatbot := service.NewChatbot(logger, bus, repo, message, bot)
		_, err = chatbot.Register(ctx, &pb.BotRequest{Bot: &b})
		if err != nil {
			logger.Error(err, zap.Any("event", event.BotRegisterSubject))
			return
		}
	})

	return nil
}
