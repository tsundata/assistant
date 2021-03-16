package opcode

import (
	"github.com/tidwall/gjson"
	"github.com/tsundata/assistant/internal/app/workflow/action/inside"
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

func (o *Json) Run(ctx *inside.Context, _ []interface{}) (interface{}, error) {
	if text, ok := ctx.Value.(string); ok {
		result := gjson.Parse(text).Value()
		ctx.SetValue(result)
		return result, nil
	}
	return nil, nil
}
