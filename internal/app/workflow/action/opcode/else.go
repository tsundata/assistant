package opcode

import (
	"github.com/tsundata/assistant/internal/app/workflow/action/inside"
)

type Else struct{}

func NewElse() *Else {
	return &Else{}
}

func (e *Else) Run(ctx *inside.Context, _ []interface{}) (interface{}, error) {
	ctx.SetContinue(!ctx.Continue)
	return nil, nil
}
