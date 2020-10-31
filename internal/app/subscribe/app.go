package subscribe

import (
	"github.com/spf13/viper"
	"github.com/tsundata/assistant/internal/app/subscribe/service"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc"
	"gorm.io/gorm"
)

type Options struct {
	Name string
	v    *viper.Viper
	db   *gorm.DB
}

func NewOptions(v *viper.Viper, db *gorm.DB) (*Options, error) {
	var err error
	o := new(Options)
	o.v = v
	o.db = db

	return o, err
}

func NewApp(o *Options, rs *rpc.Server) (*app.Application, error) {
	var subscribe service.Subscribe
	subscribe.DB = o.db
	rs.Register(&subscribe)

	a, err := app.New(o.Name, app.RpcServerOption(rs))

	if err != nil {
		return nil, err
	}

	return a, nil
}
