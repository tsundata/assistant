package opcode

import (
	"context"
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
	return "debug : (nil -> bool)"
}

func (o *Debug) Run(_ context.Context, comp *inside.Component, params []interface{}) (interface{}, error) {
	comp.Debug = true
	return true, nil
}
