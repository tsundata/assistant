package listener

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/nats-io/nats.go"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/workflow/service"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/model"
)

func RegisterEventHandler(bus *event.Bus, rdb *redis.Client, message pb.MessageClient, logger *logger.Logger) error {
	err := bus.Subscribe(event.RunWorkflowSubject, func(msg *nats.Msg) {
		var m model.Message
		err := json.Unmarshal(msg.Data, &message)
		if err != nil {
			logger.Error(err)
			return
		}

		ctx := context.Background()
		reply, err := message.Get(ctx, &pb.MessageRequest{Id: int64(m.ID)})
		if err != nil {
			logger.Error(err)
			return
		}

		switch reply.GetType() {
		case model.MessageTypeAction:
			workflow := service.NewWorkflow(bus, rdb, nil, message, logger)
			_, err := workflow.RunAction(ctx, &pb.WorkflowRequest{Text: reply.GetText()})
			if err != nil {
				logger.Error(err)
				return
			}
		}
	})
	return err
}
