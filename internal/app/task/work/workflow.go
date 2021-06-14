package work

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/model"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc/rpcclient"
	"github.com/tsundata/assistant/internal/pkg/util"
	"strconv"
)

type WorkflowTask struct {
	bus    *event.Bus
	client *rpc.Client
}

func NewWorkflowTask(bus *event.Bus, client *rpc.Client) *WorkflowTask {
	return &WorkflowTask{bus: bus, client: client}
}

func (t *WorkflowTask) Run(data string) (bool, error) {
	var args map[string]string
	err := json.Unmarshal(util.StringToByte(data), &args)
	if err != nil {
		return false, err
	}

	tp, ok := args["type"]
	if !ok {
		return false, errors.New("error arg type")
	}

	idStr, ok := args["id"]
	if !ok {
		return false, errors.New("error arg id")
	}
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return false, err
	}

	message, err := rpcclient.GetMessageClient(t.client).Get(context.Background(), &pb.MessageRequest{Id: id})
	if err != nil {
		return false, err
	}

	switch tp {
	case model.MessageTypeAction:
		_, err := rpcclient.GetWorkflowClient(t.client).RunAction(context.Background(), &pb.WorkflowRequest{Text: message.GetText()})
		if err != nil {
			return false, err
		}
		return true, nil
	default:
		return false, errors.New("error type")
	}
}
