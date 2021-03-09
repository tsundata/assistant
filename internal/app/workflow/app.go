package workflow

import (
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/workflow/service"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc"
	"go.etcd.io/etcd/clientv3"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func NewApp(name string, logger *zap.Logger, rs *rpc.Server, etcd *clientv3.Client, midClient pb.MiddleClient, msgClient pb.MessageClient) (*app.Application, error) {
	// service
	subscribe := service.NewWorkflow(etcd, midClient, msgClient)
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
