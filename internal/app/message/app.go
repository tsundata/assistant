package message

import (
	"errors"
	"github.com/spf13/viper"
	"github.com/tsundata/assistant/internal/app/message/service"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Options struct {
	Name    string
	Webhook string
	db      *gorm.DB
}

func NewOptions(v *viper.Viper, db *gorm.DB) (*Options, error) {
	var err error
	o := new(Options)
	o.db = db

	if err = v.UnmarshalKey("app", o); err != nil {
		return nil, errors.New("unmarshal app option error")
	}

	if err = v.UnmarshalKey("slack", o); err != nil {
		return nil, errors.New("unmarshal app option error")
	}

	return o, err
}

func NewApp(o *Options, logger *zap.Logger, rs *rpc.Server) (*app.Application, error) {
	message := service.NewManage(o.db)
	err := rs.Register(message, "")
	if err != nil {
		return nil, err
	}

	a, err := app.New(o.Name, logger, app.RpcServerOption(rs))
	if err != nil {
		return nil, err
	}

	return a, nil
}
