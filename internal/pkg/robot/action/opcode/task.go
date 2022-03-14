package opcode

import (
	"context"
	"errors"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/robot/action/inside"
)

type Task struct{}

func NewTask() *Task {
	return &Task{}
}

func (o *Task) Type() int {
	return TypeOp
}

func (o *Task) Doc() string {
	return "task [integer] : (nil -> bool)"
}

func (o *Task) Run(ctx context.Context, comp *inside.Component, params []interface{}) (interface{}, error) {
	if len(params) != 1 {
		return false, app.ErrInvalidParameter
	}

	if comp.Bus == nil {
		return false, errors.New("error client")
	}

	if id, ok := params[0].(int64); ok {
		// get message
		message, err := comp.MessageClient.GetById(ctx, &pb.MessageRequest{Message: &pb.Message{Id: id}})
		if err != nil {
			return nil, err
		}

		err = comp.Bus.Publish(ctx, enum.Chatbot, event.WorkflowRunSubject, message.Message)
		if err != nil {
			return false, err
		}

		return true, nil
	}

	return false, nil
}
