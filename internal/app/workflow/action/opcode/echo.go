package opcode

import (
	"context"
	"errors"
	"github.com/tsundata/assistant/api/model"
	"github.com/tsundata/assistant/internal/app/workflow/action/inside"
	"github.com/tsundata/assistant/internal/pkg/event"
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

func (o *Echo) Run(ctx context.Context, comp *inside.Component, params []interface{}) (interface{}, error) {
	if len(params) != 1 {
		return nil, errors.New("error params")
	}

	if text, ok := params[0].(string); ok {
		if comp.Bus == nil {
			return false, nil
		}
		err := comp.Bus.Publish(ctx, event.SendMessageSubject, model.Message{Text: text})
		if err != nil {
			return false, err
		}
		comp.SetValue(true)
		return true, nil
	}
	return false, nil
}
