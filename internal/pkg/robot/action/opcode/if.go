package opcode

import (
	"context"
	"github.com/tsundata/assistant/internal/pkg/robot/action/inside"
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

func (o *If) Run(_ context.Context, comp *inside.Component, _ []interface{}) (interface{}, error) {
	if integer, ok := comp.Value.(int64); ok {
		comp.SetContinue(integer >= 0)
		return nil, nil
	}
	if float, ok := comp.Value.(float64); ok {
		comp.SetContinue(float >= 0)
		return nil, nil
	}
	if text, ok := comp.Value.(string); ok {
		comp.SetContinue(text != "")
		return nil, nil
	}
	if boolean, ok := comp.Value.(bool); ok {
		comp.SetContinue(boolean)
		return nil, nil
	}
	if objects, ok := comp.Value.(map[string]interface{}); ok {
		comp.SetContinue(len(objects) > 0)
		return nil, nil
	}
	if arrays, ok := comp.Value.([]interface{}); ok {
		comp.SetContinue(len(arrays) > 0)
		return nil, nil
	}
	comp.SetContinue(false)
	return nil, nil
}
