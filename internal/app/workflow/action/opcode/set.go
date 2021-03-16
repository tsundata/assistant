package opcode

import (
	"errors"
	"github.com/tsundata/assistant/internal/app/workflow/action/inside"
)

type Set struct{}

func NewSet() *Set {
	return &Set{}
}

func (o *Set) Type() int {
	return TypeOp
}

func (o *Set) Doc() string {
	return "set [any] : (nil -> any)"
}

func (o *Set) Run(ctx *inside.Context, params []interface{}) (interface{}, error) {
	if len(params) != 1 {
		return nil, errors.New("error params")
	}
	ctx.SetValue(params[0])
	return params[0], nil
}
