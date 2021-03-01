package opcode

import "github.com/tsundata/assistant/internal/app/workflow/action/context"

type Count struct{}

func NewCount() *Count {
	return &Count{}
}

func (c Count) Run(ctx *context.Context, _ []interface{}) (interface{}, error) {
	if text, ok := ctx.Value.(string); ok {
		return len(text), nil
	}
	return 0, nil
}
