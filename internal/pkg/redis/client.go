package redis

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

type Options struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
}

func NewOptions(v *viper.Viper) (*Options, error) {
	var err error
	o := new(Options)
	if err = v.UnmarshalKey("redis", o); err != nil {
		return nil, errors.New("unmarshal redis option error")
	}

	return o, err
}

func New(o *Options) (*redis.Client, error) {
	r := redis.NewClient(&redis.Options{
		Addr:     o.Addr,
		Password: o.Password,
		DB:       0,
	})
	s := r.Ping(context.TODO())
	result, err := s.Result()
	if err != nil {
		return nil, err
	}
	if result != "PONG" {
		return nil, errors.New("redis conn error")
	}
	return r, nil
}

var ProviderSet = wire.NewSet(New, NewOptions)
