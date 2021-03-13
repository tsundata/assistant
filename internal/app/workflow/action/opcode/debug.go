package opcode

import (
	"errors"
	"github.com/tsundata/assistant/internal/app/workflow/action/inside"
)

type Debug struct{}

func NewDebug() *Debug {
	return &Debug{}
}

func (d *Debug) Run(ctx *inside.Context, params []interface{}) (interface{}, error) {
	if len(params) != 1 {
		return false, errors.New("error params")
	}

	if state, ok := params[0].(bool); ok {
		ctx.Debug = state
		return true, nil
	}
	return false, nil
}
