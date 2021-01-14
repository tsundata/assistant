package middle

import (
	"errors"
	"github.com/spf13/viper"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/middle/service"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc"
	"go.etcd.io/bbolt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Options struct {
	Name   string
	webURL string
}

func NewOptions(v *viper.Viper) (*Options, error) {
	var err error
	o := new(Options)

	if err = v.UnmarshalKey("app", o); err != nil {
		return nil, errors.New("unmarshal app option error")
	}

	web := v.GetStringMapString("web")
	o.webURL = web["url"]

	return o, err
}

// FIXME rename
func NewApp(o *Options, logger *zap.Logger, rs *rpc.Server, db *bbolt.DB) (*app.Application, error) {
	// service
	mid := service.NewMiddle(db, o.webURL)
	err := rs.Register(func(gs *grpc.Server) error {
		pb.RegisterMiddleServer(gs, mid)
		return nil
	})
	if err != nil {
		return nil, err
	}

	a, err := app.New(o.Name, logger, app.RPCServerOption(rs))
	if err != nil {
		return nil, err
	}

	return a, nil
}
