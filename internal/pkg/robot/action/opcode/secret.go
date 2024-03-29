package opcode

import (
	"context"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/robot/action/inside"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
)

type Secret struct{}

func NewSecret() *Secret {
	return &Secret{}
}

func (o *Secret) Type() int {
	return TypeCond
}

func (o *Secret) Doc() string {
	return "secret [string] : (nil -> any)"
}

func (o *Secret) Run(ctx context.Context, inCtx *inside.Context, comp component.Component, params []interface{}) (interface{}, error) {
	if len(params) != 1 {
		return nil, nil
	}

	if comp.Middle() == nil {
		return nil, nil
	}

	if text, ok := params[0].(string); ok {
		reply, err := comp.Middle().GetCredential(ctx, &pb.CredentialRequest{Name: text})
		if err != nil {
			return nil, err
		}
		result := make(map[string]string)
		for _, item := range reply.GetContent() {
			result[item.Key] = item.Value
		}
		inCtx.SetCredential(result)
		return result, nil
	}

	return nil, nil
}
