package listener

import (
	"context"
	"encoding/json"
	"github.com/nats-io/nats.go"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/chatbot/trigger"
	"github.com/tsundata/assistant/internal/app/chatbot/trigger/ctx"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/log"
	"go.uber.org/zap"
)

func RegisterEventHandler(bus event.Bus, logger log.Logger, middle pb.MiddleSvcClient, todo pb.TodoSvcClient, user pb.UserSvcClient) error {
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

	return nil
}
