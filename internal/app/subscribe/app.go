package subscribe

import (
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"github.com/tsundata/assistant/internal/app/subscribe/service"
	"github.com/tsundata/assistant/internal/app/subscribe/spider"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Options struct {
	Name   string
	db     *gorm.DB
	redis  *redis.Client
	logger *zap.Logger
}

func NewOptions(v *viper.Viper, db *gorm.DB, logger *zap.Logger, redis *redis.Client) (*Options, error) {
	var err error
	o := new(Options)
	o.db = db
	o.redis = redis
	o.logger = logger

	if err = v.UnmarshalKey("app", o); err != nil {
		return nil, errors.New("unmarshal app option error")
	}

	return o, err
}

func NewApp(o *Options, s *spider.Spider, rs *rpc.Server) (*app.Application, error) {
	// spider cron
	s.Cron()

	// service
	subscribe := service.NewSubscribe(o.db)
	err := rs.Register(subscribe, "")
	if err != nil {
		return nil, err
	}

	a, err := app.New(o.Name, o.logger, app.RPCServerOption(rs))
	if err != nil {
		return nil, err
	}

	return a, nil
}
