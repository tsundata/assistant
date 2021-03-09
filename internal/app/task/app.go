package task

import (
	"github.com/go-redis/redis/v8"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/task/service"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc"
	"go.etcd.io/etcd/clientv3"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func NewApp(name string, logger *zap.Logger, rs *rpc.Server, etcd *clientv3.Client, rdb *redis.Client) (*app.Application, error) {
	// service
	task := service.NewTask(etcd)
	err := rs.Register(func(gs *grpc.Server) error {
		pb.RegisterTaskServer(gs, task)
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
