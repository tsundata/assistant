package opcode

import (
	"github.com/tsundata/assistant/internal/app/workflow/action/inside"
)

type Count struct{}

func NewCount() *Count {
	return &Count{}
}

func (o *Count) Type() int {
	return TypeOp
}

func (o *Count) Doc() string {
	return "count : (any -> integer)"
}

func (o *Count) Run(ctx *inside.Context, _ []interface{}) (interface{}, error) {
	if text, ok := ctx.Value.(string); ok {
		result := len(text)
		ctx.SetValue(result)
		return int64(result), nil
	}
	if objects, ok := ctx.Value.(map[string]interface{}); ok {
		result := len(objects)
		ctx.SetValue(result)
		return int64(result), nil
	}
	if arrays, ok := ctx.Value.([]interface{}); ok {
		result := len(arrays)
		ctx.SetValue(result)
		return int64(result), nil
	}
	return int64(0), nil
}
