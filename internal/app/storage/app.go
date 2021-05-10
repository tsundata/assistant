package workflow

import (
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/storage/service"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc"
	"go.etcd.io/etcd/clientv3"
	"google.golang.org/grpc"
	"os"
)

type Options struct {
	Name string
	Path string
}

func NewOptions(v *viper.Viper) (*Options, error) {
	var err error
	o := new(Options)

	o.Name = os.Getenv("APP_NAME")

	if err = v.UnmarshalKey("storage", o); err != nil {
		return nil, errors.New("unmarshal app option error")
	}

	return o, err
}

func NewApp(o *Options, logger *logger.Logger, rs *rpc.Server, etcd *clientv3.Client, db *sqlx.DB, rdb *redis.Client) (*app.Application, error) {
	// service
	s := service.NewStorage(o.Path, etcd, db, rdb)
	err := rs.Register(func(gs *grpc.Server) error {
		pb.RegisterStorageServer(gs, s)
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

var ProviderSet = wire.NewSet(NewApp, NewOptions)