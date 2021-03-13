package opcode

import (
	"github.com/tsundata/assistant/internal/app/workflow/action/inside"
)

type Count struct{}

func NewCount() *Count {
	return &Count{}
}

func (c *Count) Run(ctx *inside.Context, _ []interface{}) (interface{}, error) {
	if text, ok := ctx.Value.(string); ok {
		result := len(text)
		ctx.SetValue(result)
		return result, nil
	}
	return 0, nil
}
