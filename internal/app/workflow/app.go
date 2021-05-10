package workflow

import (
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/workflow/repository"
	"github.com/tsundata/assistant/internal/app/workflow/service"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc"
	"go.etcd.io/etcd/clientv3"
	"google.golang.org/grpc"
	"os"
)

func NewApp(logger *logger.Logger, rs *rpc.Server, etcd *clientv3.Client, rdb *redis.Client, repo repository.WorkflowRepository,
	midClient pb.MiddleClient, msgClient pb.MessageClient, taskClient pb.TaskClient) (*app.Application, error) {
	name := os.Getenv("APP_NAME")

	// service
	s := service.NewWorkflow(etcd, rdb, repo, midClient, msgClient, taskClient)
	err := rs.Register(func(gs *grpc.Server) error {
		pb.RegisterWorkflowServer(gs, s)
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

var ProviderSet = wire.NewSet(NewApp)
