package opcode

import (
	"errors"
	"github.com/tsundata/assistant/internal/app/workflow/action/inside"
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

func (o *Debug) Run(ctx *inside.Context, params []interface{}) (interface{}, error) {
	if len(params) == 1 {
		if state, ok := params[0].(bool); ok {
			ctx.Debug = state
			return state, nil
		}
	} else if len(params) == 0 {
		ctx.Debug = true
		return true, nil
	} else {
		return false, errors.New("error params")
	}

	return false, nil
}
