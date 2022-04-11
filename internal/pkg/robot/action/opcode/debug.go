package opcode

import (
	"context"
	"github.com/tsundata/assistant/internal/pkg/robot/action/inside"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
)

type Debug struct{}

func NewDebug() *Debug {
	return &Debug{}
}

func (o *Debug) Type() int {
	return TypeOp
}

func (o *Debug) Doc() string {
	return "debug : (nil -> bool)"
}

func (o *Debug) Run(_ context.Context, inCtx *inside.Context, _ component.Component, params []interface{}) (interface{}, error) {
	inCtx.Debug = true
	return true, nil
}
