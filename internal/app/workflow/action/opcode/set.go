package opcode

import (
	"context"
	"github.com/tsundata/assistant/internal/app/workflow/action/inside"
	"github.com/tsundata/assistant/internal/pkg/app"
)

type Set struct{}

func NewSet() *Set {
	return &Set{}
}

func (o *Set) Type() int {
	return TypeOp
}

func (o *Set) Doc() string {
	return "set [any]... : (nil -> any)"
}

func (o *Set) Run(_ context.Context, comp *inside.Component, params []interface{}) (interface{}, error) {
	if len(params) < 1 {
		return nil, app.ErrInvalidParameter
	}
	if len(params) > 1 {
		comp.SetValue(params)
		return params, nil
	}
	comp.SetValue(params[0])
	return params[0], nil
}
