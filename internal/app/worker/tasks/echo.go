package tasks

import (
	"context"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc/rpcclient"
)

type EchoTask struct {
	client *rpc.Client
}

func NewEchoTask(client *rpc.Client) *EchoTask {
	return &EchoTask{client: client}
}

func (t *EchoTask) Echo(data string) (bool, error) {
	_, err := rpcclient.GetMessageClient(t.client).Send(context.Background(), &pb.MessageRequest{Text: data})
	if err != nil {
		return false, err
	}
	return true, nil
}
