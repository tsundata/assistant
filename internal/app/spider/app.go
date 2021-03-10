package spider

import (
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/spider/crawler"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"time"
)

type Options struct {
	Path string
}

func NewOptions(v *viper.Viper) (*Options, error) {
	var err error
	o := new(Options)

	if err = v.UnmarshalKey("plugin", o); err != nil {
		return nil, errors.New("unmarshal plugin option error")
	}

	return o, err
}

func NewApp(name string, o *Options, rdb *redis.Client, logger *logger.Logger,
	msgClient pb.MessageClient, midClient pb.MiddleClient, subClient pb.SubscribeClient) (*app.Application, error) {
	go func() {
		// Delayed loading
		time.Sleep(10 * time.Second)
		s := crawler.New()
		s.SetService(rdb, logger, msgClient, midClient, subClient)
		err := s.LoadRule(o.Path)
		if err != nil {
			logger.Error(err)
			return
		}
		s.Daemon()
	}()

	a, err := app.New(name, logger)
	if err != nil {
		return nil, err
	}

	return a, nil
}
