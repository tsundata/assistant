package listener

import (
	"context"
	"encoding/json"
	"github.com/nats-io/nats.go"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/chatbot/repository"
	"github.com/tsundata/assistant/internal/app/chatbot/service"
	"github.com/tsundata/assistant/internal/app/chatbot/trigger"
	"github.com/tsundata/assistant/internal/app/chatbot/trigger/ctx"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc/md"
	"go.uber.org/zap"
)

func RegisterEventHandler(bus event.Bus, logger log.Logger, bot *rulebot.RuleBot, message pb.MessageSvcClient,
	middle pb.MiddleSvcClient, todo pb.TodoSvcClient, user pb.UserSvcClient, repo repository.ChatbotRepository) error {
	err := bus.Subscribe(context.Background(), event.MessageTriggerSubject, func(msg *nats.Msg) {
		var m pb.Message
		err := json.Unmarshal(msg.Data, &m)
		if err != nil {
			logger.Error(err, zap.Any("event", event.MessageTriggerSubject))
			return
		}

		comp := ctx.NewComponent()
		comp.Logger = logger
		comp.Middle = middle
		comp.Todo = todo
		comp.User = user
		comp.Bus = bus
		trigger.Run(context.Background(), comp, m.Text)
	})
	if err != nil {
		return err
	}

	err = bus.Subscribe(context.Background(), event.MessageHandleSubject, func(msg *nats.Msg) {
		var m pb.Message
		err := json.Unmarshal(msg.Data, &m)
		if err != nil {
			logger.Error(err, zap.Any("event", event.MessageHandleSubject))
			return
		}

		chatbot := service.NewChatbot(logger, repo, message, middle, todo, bot)
		_, err = chatbot.Handle(md.BuildAuthContext(m.UserId), &pb.ChatbotRequest{MessageId: m.Id})
		if err != nil {
			logger.Error(err, zap.Any("event", event.MessageHandleSubject))
			return
		}
	})
	if err != nil {
		return err
	}

	return nil
}
