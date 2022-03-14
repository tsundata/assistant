package opcode

import (
	"context"
	"github.com/tidwall/gjson"
	"github.com/tsundata/assistant/internal/pkg/robot/action/inside"
)

type Json struct{}

func NewJson() *Json {
	return &Json{}
}

func (o *Json) Type() int {
	return TypeOp
}

func (o *Json) Doc() string {
	return "json : (string -> any)"
}

func (o *Json) Run(_ context.Context, comp *inside.Component, _ []interface{}) (interface{}, error) {
	if text, ok := comp.Value.(string); ok {
		result := gjson.Parse(text).Value()
		comp.SetValue(result)
		return result, nil
	}
	return nil, nil
}
