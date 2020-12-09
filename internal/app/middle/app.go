package middle

import (
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"github.com/tsundata/assistant/internal/app/middle/service"
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
	webURL string
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

	web := v.GetStringMapString("web")
	o.webURL = web["url"]

	return o, err
}

func NewApp(o *Options, rs *rpc.Server) (*app.Application, error) {
	// service
	page := service.NewWeb(o.db, o.webURL)
	err := rs.Register(page, "")
	if err != nil {
		return nil, err
	}

	a, err := app.New(o.Name, o.logger, app.RPCServerOption(rs))
	if err != nil {
		return nil, err
	}

	return a, nil
}
