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

	if comp.Bus == nil || comp.MessageClient == nil {
		return false, errors.New("error client")
	}

	if sequence, ok := params[0].(int64); ok {
		// get message
		message, err := comp.MessageClient.GetBySequence(ctx, &pb.MessageRequest{Message: &pb.Message{Sequence: sequence}})
		if err != nil {
			return nil, err
		}

		// run script
		err = comp.Bus.Publish(ctx, enum.Chatbot, event.ScriptRunSubject, message.Message)
		if err != nil {
			return false, err
		}

		return true, nil
	}

	return false, nil
}
