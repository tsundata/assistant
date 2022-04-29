package system

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
	err := comp.GetBus().Subscribe(ctx, enum.Chatbot, event.SystemCounterCreateSubject, func(msg *event.Msg) error {
		var m pb.Counter
		err := json.Unmarshal(msg.Data, &m)
		if err != nil {
			return err
		}

		ctx = md.BuildAuthContext(m.UserId)
		find, err := comp.System().GetCounterByFlag(ctx, &pb.CounterRequest{Counter: &m})
		if err != nil {
			return err
		}
		if find.Counter.Id > 0 {
			return nil
		}
		_, err = comp.System().CreateCounter(ctx, &pb.CounterRequest{Counter: &m})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	err = comp.GetBus().Subscribe(ctx, enum.Chatbot, event.SystemCounterIncreaseSubject, func(msg *event.Msg) error {
		var m pb.Counter
		err := json.Unmarshal(msg.Data, &m)
		if err != nil {
			return err
		}

		ctx = md.BuildAuthContext(m.UserId)
		_, err = comp.System().ChangeCounter(ctx, &pb.CounterRequest{Counter: &m})
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	err = comp.GetBus().Subscribe(ctx, enum.Chatbot, event.SystemCounterDecreaseSubject, func(msg *event.Msg) error {
		var m pb.Counter
		err := json.Unmarshal(msg.Data, &m)
		if err != nil {
			return err
		}

		ctx = md.BuildAuthContext(m.UserId)
		m.Digit = -m.Digit
		_, err = comp.System().ChangeCounter(ctx, &pb.CounterRequest{Counter: &m})
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
