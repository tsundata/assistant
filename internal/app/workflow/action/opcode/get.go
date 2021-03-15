package opcode

import (
	"errors"
	"github.com/go-resty/resty/v2"
	"github.com/tsundata/assistant/internal/app/workflow/action/inside"
	"github.com/tsundata/assistant/internal/pkg/utils"
	"time"
)

type Get struct{}

func NewGet() *Get {
	return &Get{}
}

func (o *Get) Run(ctx *inside.Context, params []interface{}) (interface{}, error) {
	if len(params) != 1 {
		return nil, errors.New("error params")
	}
	if text, ok := params[0].(string); ok {
		if utils.IsUrl(text) {
			client := resty.New()
			client.SetTimeout(time.Minute)
			resp, err := client.R().Get(text)
			if err != nil {
				return nil, err
			}
			result := utils.ByteToString(resp.Body())
			ctx.SetValue(result)
			return result, nil
		}
	}
	return nil, nil
}
