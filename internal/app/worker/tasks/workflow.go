package tasks

import (
	"context"
	"fmt"
	"github.com/tsundata/assistant/api/pb"
	"log"
)

type WorkflowTask struct {
	msgClient pb.MessageClient
}

func NewWorkflowTask(msgClient pb.MessageClient) *WorkflowTask {
	return &WorkflowTask{msgClient: msgClient}
}

func (t *WorkflowTask) Run(args ...int64) (int64, error) {
	sum := int64(0)
	for _, arg := range args {
		sum += arg
	}

	_, err := t.msgClient.Send(context.Background(), &pb.MessageRequest{Text: fmt.Sprintf("Sum: %d", sum)})
	if err != nil {
		log.Println(err)
	}

	return sum, nil
}
