package opcode

import (
	"github.com/tsundata/assistant/internal/app/workflow/action/inside"
	"reflect"
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
	v := reflect.ValueOf(ctx.Value)
	if v.Kind() == reflect.Slice || v.Kind() == reflect.Map {
		result := int64(v.Len())
		ctx.SetValue(result)
		return result, nil
	}
	return int64(0), nil
}
