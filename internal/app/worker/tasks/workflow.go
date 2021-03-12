package tasks

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/model"
	"github.com/tsundata/assistant/internal/pkg/utils"
	"strconv"
)

type WorkflowTask struct {
	msgClient pb.MessageClient
	wfClient  pb.WorkflowClient
}

func NewWorkflowTask(msgClient pb.MessageClient, wfClient pb.WorkflowClient) *WorkflowTask {
	return &WorkflowTask{msgClient: msgClient, wfClient: wfClient}
}

func (t *WorkflowTask) Run(data string) (bool, error) {
	var args map[string]string
	err := json.Unmarshal(utils.StringToByte(data), &args)
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

	message, err := t.msgClient.Get(context.Background(), &pb.MessageRequest{Id: id})
	if err != nil {
		return false, err
	}

	switch tp {
	case model.MessageTypeAction:
		_, err := t.wfClient.RunAction(context.Background(), &pb.WorkflowRequest{Text: message.GetText()})
		if err != nil {
			return false, err
		}
	case model.MessageTypeScript:
		_, err := t.wfClient.RunScript(context.Background(), &pb.WorkflowRequest{Text: message.GetText()})
		if err != nil {
			return false, err
		}
	default:
		return false, errors.New("error type")
	}

	return true, nil
}
