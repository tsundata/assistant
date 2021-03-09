package task

import (
	"github.com/RichardKnop/machinery/v2"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/task/service"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func NewApp(name string, logger *zap.Logger, rs *rpc.Server, ms *machinery.Server) (*app.Application, error) {
	// service
	task := service.NewTask(ms)
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
