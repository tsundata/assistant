package opcode

import (
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

func (o *Else) Run(ctx *inside.Context, _ []interface{}) (interface{}, error) {
	ctx.SetContinue(!ctx.Continue)
	return nil, nil
}
