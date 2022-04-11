package opcode

import (
	"context"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/robot/action/inside"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
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

func (o *Set) Run(_ context.Context, inCtx *inside.Context, _ component.Component, params []interface{}) (interface{}, error) {
	if len(params) < 1 {
		return nil, app.ErrInvalidParameter
	}
	if len(params) > 1 {
		inCtx.SetValue(params)
		return params, nil
	}
	inCtx.SetValue(params[0])
	return params[0], nil
}
