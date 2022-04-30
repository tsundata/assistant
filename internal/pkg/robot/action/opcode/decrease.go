package opcode

import (
	"context"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/robot/action/inside"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
)

type Decrease struct{}

func NewDecrease() *Decrease {
	return &Decrease{}
}

func (o *Decrease) Type() int {
	return TypeOp
}

func (o *Decrease) Doc() string {
	return "decrease [string] : (nil -> bool)"
}

func (o *Decrease) Run(ctx context.Context, inCtx *inside.Context, comp component.Component, params []interface{}) (interface{}, error) {
	if flag, ok := params[0].(string); ok {
		err := comp.GetBus().Publish(ctx, enum.Middle, event.CounterDecreaseSubject, pb.Counter{
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
