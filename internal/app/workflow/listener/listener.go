package listener

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/workflow/service"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/log"
)

func RegisterEventHandler(bus event.Bus, rdb *redis.Client, message pb.MessageSvcClient, middle pb.MiddleSvcClient, logger log.Logger) error {
	err := bus.Subscribe(context.Background(), enum.Workflow, event.WorkflowRunSubject, func(msg *event.Msg) error {
		var m pb.Message
		err := json.Unmarshal(msg.Data, &m)
		if err != nil {
			return err
		}

		ctx := context.Background()
		reply, err := message.Get(ctx, &pb.MessageRequest{Message: &pb.Message{Id: m.Id}})
		if err != nil {
			return err
		}

		switch reply.Message.GetType() {
		case enum.MessageTypeAction:
			workflow := service.NewWorkflow(bus, rdb, nil, message, middle, logger)
			_, err := workflow.RunAction(ctx, &pb.WorkflowRequest{Text: reply.Message.GetText()})
			if err != nil {
				return err
			}
		}
		return nil
	})
	return err
}
