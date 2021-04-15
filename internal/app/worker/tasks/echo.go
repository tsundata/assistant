package tasks

import (
	"context"
	"github.com/tsundata/assistant/api/pb"
)

type EchoTask struct {
	msgClient pb.MessageClient
	wfClient  pb.WorkflowClient
}

func NewEchoTask(msgClient pb.MessageClient, wfClient pb.WorkflowClient) *EchoTask {
	return &EchoTask{msgClient: msgClient, wfClient: wfClient}
}

func (t *EchoTask) Echo(data string) (bool, error) {
	_, err := t.msgClient.Send(context.Background(), &pb.MessageRequest{Text: data})
	if err != nil {
		return false, err
	}
	return true, nil
}
