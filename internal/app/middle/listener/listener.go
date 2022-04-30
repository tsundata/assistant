package listener

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/middle/repository"
	"github.com/tsundata/assistant/internal/app/middle/service"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/global"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc/md"
)

func RegisterEventHandler(bus event.Bus, rdb *redis.Client, locker *global.Locker, repo repository.MiddleRepository) error {
	ctx := context.Background()
	middle := service.NewMiddle(nil, rdb, locker, repo, nil)
	err := bus.Subscribe(ctx, enum.Middle, event.CronRegisterSubject, func(msg *event.Msg) error {
		var m pb.CronRequest
		err := json.Unmarshal(msg.Data, &m)
		if err != nil {
			return err
		}

		_, err = middle.RegisterCron(ctx, &pb.CronRequest{Text: m.Text})
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	err = bus.Subscribe(ctx, enum.Middle, event.SubscribeRegisterSubject, func(msg *event.Msg) error {
		var m pb.SubscribeRequest
		err := json.Unmarshal(msg.Data, &m)
		if err != nil {
			return err
		}

		_, err = middle.RegisterSubscribe(ctx, &pb.SubscribeRequest{Text: m.Text})
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	err = bus.Subscribe(ctx, enum.Middle, event.CounterCreateSubject, func(msg *event.Msg) error {
		var m pb.Counter
		err := json.Unmarshal(msg.Data, &m)
		if err != nil {
			return err
		}

		ctx = md.BuildAuthContext(m.UserId)
		find, err := middle.GetCounterByFlag(ctx, &pb.CounterRequest{Counter: &m})
		if err != nil {
			return err
		}
		if find.Counter.Id > 0 {
			return nil
		}
		_, err = middle.CreateCounter(ctx, &pb.CounterRequest{Counter: &m})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	err = bus.Subscribe(ctx, enum.Middle, event.CounterIncreaseSubject, func(msg *event.Msg) error {
		var m pb.Counter
		err := json.Unmarshal(msg.Data, &m)
		if err != nil {
			return err
		}

		ctx = md.BuildAuthContext(m.UserId)
		_, err = middle.ChangeCounter(ctx, &pb.CounterRequest{Counter: &m})
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	err = bus.Subscribe(ctx, enum.Middle, event.CounterDecreaseSubject, func(msg *event.Msg) error {
		var m pb.Counter
		err := json.Unmarshal(msg.Data, &m)
		if err != nil {
			return err
		}

		ctx = md.BuildAuthContext(m.UserId)
		m.Digit = -m.Digit
		_, err = middle.ChangeCounter(ctx, &pb.CounterRequest{Counter: &m})
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
