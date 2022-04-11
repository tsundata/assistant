package opcode

import (
	"context"
	"github.com/tidwall/gjson"
	"github.com/tsundata/assistant/internal/pkg/robot/action/inside"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
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

func (o *Json) Run(_ context.Context, inCtx *inside.Context, _ component.Component, _ []interface{}) (interface{}, error) {
	if text, ok := inCtx.Value.(string); ok {
		result := gjson.Parse(text).Value()
		inCtx.SetValue(result)
		return result, nil
	}
	return nil, nil
}
