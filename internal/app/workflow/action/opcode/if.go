package opcode

import "github.com/tsundata/assistant/internal/app/workflow/action/inside"

type If struct{}

func NewIf() *If {
	return &If{}
}

func (e *If) Run(ctx *inside.Context, _ []interface{}) (interface{}, error) {
	if integer, ok := ctx.Value.(int64); ok {
		ctx.SetContinue(integer >= 0)
		return nil, nil
	}
	if float, ok := ctx.Value.(float64); ok {
		ctx.SetContinue(float >= 0)
		return nil, nil
	}
	if text, ok := ctx.Value.(string); ok {
		ctx.SetContinue(text != "")
		return nil, nil
	}
	if boolean, ok := ctx.Value.(bool); ok {
		ctx.SetContinue(boolean)
		return nil, nil
	}
	if objects, ok := ctx.Value.(map[string]interface{}); ok {
		ctx.SetContinue(len(objects) > 0)
		return nil, nil
	}
	if arrays, ok := ctx.Value.([]interface{}); ok {
		ctx.SetContinue(len(arrays) > 0)
		return nil, nil
	}
	ctx.SetContinue(false)
	return nil, nil
}
