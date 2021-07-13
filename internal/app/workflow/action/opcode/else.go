package opcode

import (
	"context"
	"github.com/tsundata/assistant/internal/app/workflow/action/inside"
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

func (o *Else) Run(_ context.Context, comp *inside.Component, _ []interface{}) (interface{}, error) {
	comp.SetContinue(!comp.Continue)
	return nil, nil
}
