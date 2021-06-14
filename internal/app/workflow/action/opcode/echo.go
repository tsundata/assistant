package opcode

import (
	"errors"
	"github.com/tsundata/assistant/internal/app/workflow/action/inside"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/model"
)

type Echo struct{}

func NewEcho() *Echo {
	return &Echo{}
}

func (o *Echo) Type() int {
	return TypeOp
}

func (o *Echo) Doc() string {
	return "echo [any] : (nil -> bool)"
}

func (o *Echo) Run(ctx *inside.Context, params []interface{}) (interface{}, error) {
	if len(params) != 1 {
		return nil, errors.New("error params")
	}

	if text, ok := params[0].(string); ok {
		if ctx.Bus == nil {
			return false, nil
		}
		err := ctx.Bus.Publish(event.SendMessageSubject, model.Message{Text: text})
		if err != nil {
			return false, err
		}
		ctx.SetValue(true)
		return true, nil
	}
	return false, nil
}
