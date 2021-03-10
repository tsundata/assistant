package middle

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/middle/service"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc"
	"go.etcd.io/etcd/clientv3"
	"google.golang.org/grpc"
)

type Options struct {
	URL string
}

func NewOptions(v *viper.Viper) (*Options, error) {
	var err error
	o := new(Options)

	if err = v.UnmarshalKey("web", o); err != nil {
		return nil, errors.New("unmarshal app option error")
	}

	return o, err
}

func NewApp(name string, o *Options, logger *logger.Logger, rs *rpc.Server, db *sqlx.DB, etcd *clientv3.Client) (*app.Application, error) {
	// service
	mid := service.NewMiddle(db, etcd, o.URL)
	err := rs.Register(func(gs *grpc.Server) error {
		pb.RegisterMiddleServer(gs, mid)
		return nil
	})
	if err != nil {
		return nil, err
	}

	a, err := app.New(name, logger, app.RPCServerOption(rs))
	if err != nil {
		return nil, err
	}

	return a, nil
}
