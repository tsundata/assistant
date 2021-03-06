package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/tsundata/assistant/internal/pkg/config"
)

func New(c *config.AppConfig) (*redis.Client, error) {
	r := redis.NewClient(&redis.Options{
		Addr:     c.Redis.Addr,
		Password: c.Redis.Password,
		DB:       0,
	})
	s := r.Ping(context.TODO())
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
