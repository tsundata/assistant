package opcode

import (
	"context"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/robot/action/inside"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
)

type Counter struct{}

func NewCounter() *Counter {
	return &Counter{}
}

func (o *Counter) Type() int {
	return TypeOp
}

func (o *Counter) Doc() string {
	return "counter [string] : (nil -> bool)"
}

func (o *Counter) Run(ctx context.Context, inCtx *inside.Context, comp component.Component, params []interface{}) (interface{}, error) {
	if flag, ok := params[0].(string); ok {
		err := comp.GetBus().Publish(ctx, enum.Middle, event.CounterCreateSubject, pb.Counter{
			UserId: inCtx.Message.GetUserId(),
			Flag:   flag,
			Digit:  1,
		})
		if err != nil {
			return nil, err
		}
	}
	inCtx.SetValue(true)
	return true, nil
}
