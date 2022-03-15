package opcode

import (
	"context"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/robot/action/inside"
)

type Debug struct{}

func NewDebug() *Debug {
	return &Debug{}
}

func (o *Debug) Type() int {
	return TypeOp
}

func (o *Debug) Doc() string {
	return "debug [bool]? : (nil -> bool)"
}

func (o *Debug) Run(_ context.Context, comp *inside.Component, params []interface{}) (interface{}, error) {
	if len(params) == 1 {
		if state, ok := params[0].(bool); ok {
			comp.Debug = state
			return state, nil
		}
	} else if len(params) == 0 {
		comp.Debug = true
		return true, nil
	} else {
		return false, app.ErrInvalidParameter
	}

	return false, nil
}