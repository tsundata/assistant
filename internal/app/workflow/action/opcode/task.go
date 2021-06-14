package opcode

import (
	"errors"
	"github.com/tsundata/assistant/internal/app/workflow/action/inside"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/model"
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

func (o *Task) Run(ctx *inside.Context, params []interface{}) (interface{}, error) {
	if len(params) != 1 {
		return false, errors.New("error params")
	}

	if ctx.Bus == nil {
		return false, errors.New("error client")
	}

	if id, ok := params[0].(int64); ok {
		err := ctx.Bus.Publish(event.RunWorkflowSubject, model.Message{ID: int(id)})
		if err != nil {
			return false, err
		}

		return true, nil
	}

	return false, nil
}
