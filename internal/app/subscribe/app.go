package subscribe

import (
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/subscribe/service"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
	"google.golang.org/grpc"
)

func NewApp(c *config.AppConfig, logger *logger.Logger, rs *rpc.Server, rdb *redis.Client) (*app.Application, error) {
	// service
	s := service.NewSubscribe(rdb)
	err := rs.Register(func(gs *grpc.Server) error {
		pb.RegisterSubscribeServer(gs, s)
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
