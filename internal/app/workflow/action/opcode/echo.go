package opcode

import (
	"context"
	"errors"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/workflow/action/inside"
)

type Echo struct{}

func NewEcho() *Echo {
	return &Echo{}
}

func (o *Echo) Type() int {
	return TypeOp
}

func (o *Echo) Doc() string {
	return "echo [any] : (nil -> bool)"
}

func (o *Echo) Run(ctx *inside.Context, params []interface{}) (interface{}, error) {
	if len(params) != 1 {
		return nil, errors.New("error params")
	}

	if text, ok := params[0].(string); ok {
		if ctx.MsgClient == nil {
			return false, nil
		}
		state, err := ctx.MsgClient.Send(context.Background(), &pb.MessageRequest{Text: text})
		if err != nil {
			return false, err
		}
		ctx.SetValue(state.GetState())
		return state.GetState(), nil
	}
	return false, nil
}
