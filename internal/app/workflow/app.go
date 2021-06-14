package workflow

import (
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/workflow/listener"
	"github.com/tsundata/assistant/internal/app/workflow/repository"
	"github.com/tsundata/assistant/internal/app/workflow/service"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
	"google.golang.org/grpc"
)

func NewApp(c *config.AppConfig, bus *event.Bus, logger *logger.Logger, rs *rpc.Server, rdb *redis.Client, repo repository.WorkflowRepository,
	client *rpc.Client) (*app.Application, error) {

	// event bus register
	err := listener.RegisterEventHandler(bus, client)
	if err != nil {
		return nil, err
	}

	// service
	s := service.NewWorkflow(bus, rdb, repo, client)
	err = rs.Register(func(gs *grpc.Server) error {
		pb.RegisterWorkflowServer(gs, s)
		return nil
	})
	if err != nil {
		return nil, err
	}

	a, err := app.New(c, logger, app.RPCServerOption(rs))
	if err != nil {
		return nil, err
	}

	return a, nil
}

var ProviderSet = wire.NewSet(NewApp)
