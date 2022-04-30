package opcode

import (
	"context"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/robot/action/inside"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
)

type Increase struct{}

func NewIncrease() *Increase {
	return &Increase{}
}

func (o *Increase) Type() int {
	return TypeOp
}

func (o *Increase) Doc() string {
	return "increase [string] : (nil -> bool)"
}

func (o *Increase) Run(ctx context.Context, inCtx *inside.Context, comp component.Component, params []interface{}) (interface{}, error) {
	if flag, ok := params[0].(string); ok {
		err := comp.GetBus().Publish(ctx, enum.Middle, event.CounterIncreaseSubject, pb.Counter{
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
