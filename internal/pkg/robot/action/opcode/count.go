package opcode

import (
	"context"
	"github.com/tsundata/assistant/internal/pkg/robot/action/inside"
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

func (o *Count) Run(_ context.Context, comp *inside.Component, _ []interface{}) (interface{}, error) {
	if text, ok := comp.Value.(string); ok {
		result := len(text)
		comp.SetValue(result)
		return int64(result), nil
	}
	v := reflect.ValueOf(comp.Value)
	if v.Kind() == reflect.Slice || v.Kind() == reflect.Map {
		result := int64(v.Len())
		comp.SetValue(result)
		return result, nil
	}
	return int64(0), nil
}
