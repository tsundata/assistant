package listener

import (
	"context"
	"encoding/json"
	"github.com/nats-io/nats.go"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/model"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc/rpcclient"
)

func RegisterEventHandler(bus *event.Bus, client *rpc.Client, logger *logger.Logger) error {
	err := bus.Subscribe(event.RunWorkflowSubject, func(msg *nats.Msg) {
		var message model.Message
		err := json.Unmarshal(msg.Data, &message)
		if err != nil {
			logger.Error(err)
			return
		}

		reply, err := rpcclient.GetMessageClient(client).Get(context.Background(), &pb.MessageRequest{Id: int64(message.ID)})
		if err != nil {
			logger.Error(err)
			return
		}

		switch reply.GetType() {
		case model.MessageTypeAction:
			_, err := rpcclient.GetWorkflowClient(client).RunAction(context.Background(), &pb.WorkflowRequest{Text: reply.GetText()})
			if err != nil {
				logger.Error(err)
				return
			}
		}
	})
	if err != nil {
		return err
	}

	return nil
}
