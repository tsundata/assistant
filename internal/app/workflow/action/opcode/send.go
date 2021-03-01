package opcode

import (
	cContext "context"
	"errors"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/workflow/action/context"
)

type Send struct{}

func NewSend() *Send {
	return &Send{}
}

func (c *Send) Run(ctx *context.Context, params []interface{}) (interface{}, error) {
	if len(params) != 1 {
		return nil, errors.New("error params")
	}

	if text, ok := params[0].(string); ok {
		if ctx.MsgClient == nil {
			return false, nil
		}
		state, err := ctx.MsgClient.Send(cContext.Background(), &pb.MessageRequest{Text: text})
		if err != nil {
			return false, err
		}
		ctx.SetValue(state.GetState())
		return state.GetState(), nil
	}
	return false, nil
}
