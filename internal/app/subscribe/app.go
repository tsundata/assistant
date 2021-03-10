package subscribe

import (
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/subscribe/service"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc"
	"go.etcd.io/etcd/clientv3"
	"google.golang.org/grpc"
)

func NewApp(name string, logger *logger.Logger, rs *rpc.Server, etcd *clientv3.Client) (*app.Application, error) {
	// service
	subscribe := service.NewSubscribe(etcd)
	err := rs.Register(func(gs *grpc.Server) error {
		pb.RegisterSubscribeServer(gs, subscribe)
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
