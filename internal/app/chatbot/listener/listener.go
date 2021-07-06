package listener

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/chatbot/trigger"
	"github.com/tsundata/assistant/internal/app/chatbot/trigger/ctx"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/model"
	"sync"
)

func RegisterEventHandler(bus *event.Bus, logger *logger.Logger, middle pb.MiddleClient, todo pb.TodoClient) error {
	err := bus.Subscribe(event.MessageTriggerSubject, func(msg *nats.Msg) {
		var message model.Message
		err := json.Unmarshal(msg.Data, &message)
		if err != nil {
			logger.Error(err)
			return
		}

		c := ctx.NewContext()
		c.Logger = logger
		c.Middle = middle
		c.Todo = todo
		triggers := trigger.Triggers()
		wg := sync.WaitGroup{}
		for _, item := range triggers {
			wg.Add(1)
			go func(t trigger.Trigger) {
				defer wg.Done()
				if t.Cond(message.Text) {
					t.Handle(c)
				}
			}(item)
		}
		wg.Wait()
	})
	if err != nil {
		return err
	}

	return nil
}
