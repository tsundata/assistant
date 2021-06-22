package message

import (
	"github.com/google/wire"
	"github.com/tsundata/assistant/internal/app/message/listener"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
)

func NewApp(c *config.AppConfig, bus *event.Bus, logger *logger.Logger, rs *rpc.Server) (*app.Application, error) {
	// event bus register
	err := listener.RegisterEventHandler(bus, c, logger)
	if err != nil {
		return nil, err
	}

	// rpc server
	a, err := app.New(c, logger, app.RPCServerOption(rs))
	if err != nil {
		return nil, err
	}

	return a, nil
}

var ProviderSet = wire.NewSet(NewApp)
