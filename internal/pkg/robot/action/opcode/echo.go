package opcode

import (
	"context"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/robot/action/inside"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
	"github.com/tsundata/assistant/internal/pkg/util"
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

func (o *Echo) Run(ctx context.Context, inCtx *inside.Context, comp component.Component, params []interface{}) (interface{}, error) {
	if len(params) != 1 {
		return nil, app.ErrInvalidParameter
	}

	if text, ok := params[0].(string); ok {
		if comp.GetBus() == nil {
			return false, nil
		}
		err := comp.GetBus().Publish(ctx, enum.Message, event.MessageChannelSubject, pb.Message{
			UserId:    inCtx.Message.GetUserId(),
			GroupId:   inCtx.Message.GetGroupId(),
			Text:      text,
			Type:      string(enum.MessageTypeText),
			Sequence:  inCtx.Message.GetSequence(),
			SendTime:  util.Now(),
		})
		if err != nil {
			return false, err
		}
		inCtx.SetValue(true)
		return true, nil
	}
	return false, nil
}
