package listener

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/tsundata/assistant/api/enum"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/middle/service"
	"github.com/tsundata/assistant/internal/pkg/event"
)

func RegisterEventHandler(bus event.Bus, rdb *redis.Client) error {
	ctx := context.Background()

	err := bus.Subscribe(ctx, enum.Middle, event.CronRegisterSubject, func(msg *event.Msg) error {
		var m pb.CronRequest
		err := json.Unmarshal(msg.Data, &m)
		if err != nil {
			return err
		}

		middle := service.NewMiddle(nil, rdb, nil, nil)
		_, err = middle.RegisterCron(ctx, &pb.CronRequest{Text: m.Text})
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
