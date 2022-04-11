package opcode

import (
	"context"
	"github.com/tsundata/assistant/internal/pkg/robot/action/inside"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
	"reflect"
	"unicode/utf8"
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

func (o *Count) Run(_ context.Context, inCtx *inside.Context, _ component.Component, _ []interface{}) (interface{}, error) {
	if text, ok := inCtx.Value.(string); ok {
		result := utf8.RuneCountInString(text)
		inCtx.SetValue(result)
		return int64(result), nil
	}
	v := reflect.ValueOf(inCtx.Value)
	if v.Kind() == reflect.Slice || v.Kind() == reflect.Map {
		result := int64(v.Len())
		inCtx.SetValue(result)
		return result, nil
	}
	return int64(0), nil
}
