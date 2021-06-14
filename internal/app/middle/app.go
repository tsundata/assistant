package middle

import (
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/hashicorp/consul/api"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/middle/repository"
	"github.com/tsundata/assistant/internal/app/middle/service"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
	"google.golang.org/grpc"
)

func NewApp(c *config.AppConfig, logger *logger.Logger, rs *rpc.Server, consul *api.Client, rdb *redis.Client, repo repository.MiddleRepository) (*app.Application, error) {
	// service
	s := service.NewMiddle(consul, rdb, repo, c.Web.Url)
	err := rs.Register(func(gs *grpc.Server) error {
		pb.RegisterMiddleServer(gs, s)
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
