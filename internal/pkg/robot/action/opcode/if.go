package opcode

import (
	"context"
	"github.com/tsundata/assistant/internal/pkg/robot/action/inside"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
)

type If struct{}

func NewIf() *If {
	return &If{}
}

func (o *If) Type() int {
	return TypeCond
}

func (o *If) Doc() string {
	return "if"
}

func (o *If) Run(_ context.Context,  inCtx *inside.Context, _ component.Component, _ []interface{}) (interface{}, error) {
	if integer, ok := inCtx.Value.(int64); ok {
		inCtx.SetContinue(integer >= 0)
		return nil, nil
	}
	if float, ok := inCtx.Value.(float64); ok {
		inCtx.SetContinue(float >= 0)
		return nil, nil
	}
	if text, ok := inCtx.Value.(string); ok {
		inCtx.SetContinue(text != "")
		return nil, nil
	}
	if boolean, ok := inCtx.Value.(bool); ok {
		inCtx.SetContinue(boolean)
		return nil, nil
	}
	if objects, ok := inCtx.Value.(map[string]interface{}); ok {
		inCtx.SetContinue(len(objects) > 0)
		return nil, nil
	}
	if arrays, ok := inCtx.Value.([]interface{}); ok {
		inCtx.SetContinue(len(arrays) > 0)
		return nil, nil
	}
	inCtx.SetContinue(false)
	return nil, nil
}
