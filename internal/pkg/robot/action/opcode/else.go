package opcode

import (
	"context"
	"github.com/tsundata/assistant/internal/pkg/robot/action/inside"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
)

type Else struct{}

func NewElse() *Else {
	return &Else{}
}

func (o *Else) Type() int {
	return TypeCond
}

func (o *Else) Doc() string {
	return "else"
}

func (o *Else) Run(_ context.Context, inCtx *inside.Context, _ component.Component, _ []interface{}) (interface{}, error) {
	inCtx.SetContinue(!inCtx.Continue)
	return nil, nil
}
