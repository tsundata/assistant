package opcode

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/robot/action/inside"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
	"github.com/tsundata/assistant/internal/pkg/util"
	"time"
)

type Get struct{}

func NewGet() *Get {
	return &Get{}
}

func (o *Get) Type() int {
	return TypeOp
}

func (o *Get) Doc() string {
	return "get [any] : (nil -> any)"
}

func (o *Get) Run(_ context.Context, inCtx *inside.Context, _ component.Component, params []interface{}) (interface{}, error) {
	if len(params) != 1 {
		return nil, app.ErrInvalidParameter
	}
	if text, ok := params[0].(string); ok {
		if util.IsUrl(text) {
			client := resty.New()
			client.SetTimeout(time.Minute)
			resp, err := client.R().Get(text)
			if err != nil {
				return nil, err
			}
			result := util.ByteToString(resp.Body())
			inCtx.SetValue(result)
			return result, nil
		}
	}
	return nil, nil
}
