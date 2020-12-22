package spider

import (
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/spider/crawler"
	"github.com/tsundata/assistant/internal/app/spider/subscribe"
	"github.com/tsundata/assistant/internal/pkg/app"
	"go.uber.org/zap"
	"time"
)

type Options struct {
	Name   string
	logger *zap.Logger
}

func NewOptions(v *viper.Viper, logger *zap.Logger) (*Options, error) {
	var err error
	o := new(Options)
	o.logger = logger

	if err = v.UnmarshalKey("app", o); err != nil {
		return nil, errors.New("unmarshal app option error")
	}

	return o, err
}

func NewApp(o *Options, rdb *redis.Client, msgClient pb.MessageClient, midClient pb.MiddleClient, subClient pb.SubscribeClient) (*app.Application, error) {
	go func() {
		// FIXME
		time.Sleep(10 * time.Second)
		s := crawler.New(rdb, o.logger, msgClient, midClient, subClient)
		s.Register(subscribe.Rules)
		s.Daemon()
	}()

	a, err := app.New(o.Name, o.logger)
	if err != nil {
		return nil, err
	}

	return a, nil
}
