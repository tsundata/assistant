package opcode

import (
	"context"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/robot/action/inside"
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

func (o *Echo) Run(ctx context.Context, comp *inside.Component, params []interface{}) (interface{}, error) {
	if len(params) != 1 {
		return nil, app.ErrInvalidParameter
	}

	if text, ok := params[0].(string); ok {
		if comp.Bus == nil {
			return false, nil
		}
		err := comp.Bus.Publish(ctx, enum.Message, event.MessageChannelSubject, pb.Message{
			UserId:  comp.Message.UserId,
			GroupId: comp.Message.GroupId,
			Text:    text,
			Type:    string(enum.MessageTypeText),
		})
		if err != nil {
			return false, err
		}
		comp.SetValue(true)
		return true, nil
	}
	return false, nil
}
