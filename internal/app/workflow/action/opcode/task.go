package opcode

import (
	"context"
	"errors"
	"github.com/tsundata/assistant/api/model"
	"github.com/tsundata/assistant/internal/app/workflow/action/inside"
	"github.com/tsundata/assistant/internal/pkg/event"
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
		return false, errors.New("error params")
	}

	if comp.Bus == nil {
		return false, errors.New("error client")
	}

	if id, ok := params[0].(int64); ok {
		err := comp.Bus.Publish(ctx, event.RunWorkflowSubject, model.Message{ID: int(id)})
		if err != nil {
			return false, err
		}

		return true, nil
	}

	return false, nil
}
