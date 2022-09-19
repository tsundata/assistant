package okr

import (
	"context"
	"encoding/json"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc/md"
)

var eventHandler = func(comp component.Component) error {
	ctx := context.Background()
	err := comp.GetBus().Subscribe(ctx, enum.Chatbot, event.OkrValueSubject, func(msg *event.Msg) error {
		var m pb.OkrValue
		err := json.Unmarshal(msg.Data, &m)
		if err != nil {
			return err
		}

		ctx = md.BuildAuthContext(m.UserId)

		reply, err := comp.Okr().GetKeyResultsByTag(ctx, &pb.KeyResultRequest{Tag: m.Tag})
		if err != nil {
			return err
		}

		for _, item := range reply.Result {
			_, err = comp.Okr().CreateKeyResultValue(ctx, &pb.KeyResultValueRequest{
				KeyResultSequence: item.Sequence,
				Value:             m.Value,
			})
			comp.GetLogger().Error(err)
		}

		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
