package workflow

import (
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/workflow/listener"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
)

func NewApp(
	c *config.AppConfig,
	bus event.Bus,
	logger log.Logger,
	rs *rpc.Server,
	rdb *redis.Client,
	middle pb.MiddleSvcClient,
	message pb.MessageSvcClient) (*app.Application, error) {
	// event bus register
	err := listener.RegisterEventHandler(bus, rdb, message, middle, logger)
	if err != nil {
		return nil, err
	}

	a, err := app.New(c, app.RPCServerOption(rs))
	if err != nil {
		return nil, err
	}

	return a, nil
}

var ProviderSet = wire.NewSet(NewApp)
