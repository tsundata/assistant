package work

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/util"
	"strconv"
)

type WorkflowTask struct {
	bus      event.Bus
	message  pb.MessageSvcClient
	workflow pb.WorkflowSvcClient
}

func NewWorkflowTask(bus event.Bus, message pb.MessageSvcClient, workflow pb.WorkflowSvcClient) *WorkflowTask {
	return &WorkflowTask{bus: bus, message: message, workflow: workflow}
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

	ctx := context.Background()
	message, err := t.message.Get(ctx, &pb.MessageRequest{Message: &pb.Message{Id: id}})
	if err != nil {
		return false, err
	}

	switch tp {
	case enum.MessageTypeAction:
		_, err = t.workflow.RunAction(ctx, &pb.WorkflowRequest{Text: message.Message.GetMessage()})
		if err != nil {
			return false, err
		}
		return true, nil
	default:
		return false, errors.New("error type")
	}
}
