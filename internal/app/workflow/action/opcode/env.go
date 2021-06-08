package opcode

import (
	"context"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/workflow/action/inside"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc/rpcclient"
)

type Env struct{}

func NewEnv() *Env {
	return &Env{}
}

func (o *Env) Type() int {
	return TypeCond
}

func (o *Env) Doc() string {
	return "env [string] : (nil -> string)"
}

func (o *Env) Run(ctx *inside.Context, params []interface{}) (interface{}, error) {
	if len(params) != 1 {
		return nil, nil
	}

	if ctx.Client == nil {
		return nil, nil
	}

	if text, ok := params[0].(string); ok {
		reply, err :=  rpcclient.GetMiddleClient(ctx.Client).GetSetting(context.Background(), &pb.TextRequest{Text: text})
		if err != nil {
			return nil, err
		}
		ctx.SetValue(reply.GetValue())
		return reply.GetValue(), nil
	}

	return nil, nil
}
