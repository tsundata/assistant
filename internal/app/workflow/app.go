package workflow

import (
	"github.com/jmoiron/sqlx"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/workflow/service"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc"
	"go.etcd.io/etcd/clientv3"
	"google.golang.org/grpc"
)

func NewApp(name string, logger *logger.Logger, rs *rpc.Server, etcd *clientv3.Client, db *sqlx.DB, midClient pb.MiddleClient, msgClient pb.MessageClient, taskClient pb.TaskClient) (*app.Application, error) {
	// service
	subscribe := service.NewWorkflow(etcd, db, midClient, msgClient, taskClient)
	err := rs.Register(func(gs *grpc.Server) error {
		pb.RegisterWorkflowServer(gs, subscribe)
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
