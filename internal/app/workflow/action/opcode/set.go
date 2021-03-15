package opcode

import (
	"errors"
	"github.com/tsundata/assistant/internal/app/workflow/action/inside"
)

type Set struct{}

func NewSet() *Set {
	return &Set{}
}

func (s *Set) Run(ctx *inside.Context, params []interface{}) (interface{}, error) {
	if len(params) != 1 {
		return nil, errors.New("error params")
	}
	ctx.SetValue(params[0])
	return params[0], nil
}
