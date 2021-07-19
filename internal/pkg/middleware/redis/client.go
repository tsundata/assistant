package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/newrelic/go-agent/v3/integrations/nrredis-v8"
	"github.com/pkg/errors"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/vendors/newrelic"
)

func New(c *config.AppConfig, nr *newrelic.App) (*redis.Client, error) {
	opt := &redis.Options{
		Addr:     c.Redis.Addr,
		Password: c.Redis.Password,
		DB:       0,
	}
	r := redis.NewClient(opt)
	r.AddHook(nrredis.NewHook(opt))

	nxt := nr.StartTransaction("redis/ping")
	defer nxt.End()
	ctx := nr.NewContext(context.Background(), nxt)
	s := r.Ping(ctx)
	result, err := s.Result()
	if err != nil {
		return nil, errors.Wrap(err, "redis server error")
	}
	if result != "PONG" {
		return nil, errors.New("redis conn error")
	}
	return r, nil
}

var ProviderSet = wire.NewSet(New)
