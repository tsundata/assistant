package message

import (
	"github.com/google/wire"
	"github.com/tsundata/assistant/internal/app/message/listener"
	"github.com/tsundata/assistant/internal/app/message/rule"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
)

func NewApp(c *config.AppConfig, bus *event.Bus, logger *logger.Logger, rs *rpc.Server, client *rpc.Client) (*app.Application, error) {

	// event bus register
	err := listener.RegisterEventHandler(bus, c, logger)
	if err != nil {
		return nil, err
	}

	// rule bot
	_ = rulebot.New(&rulebot.Context{Conf: c, Client: client, Logger: logger}, rule.Options...) // fixme

	// rpc server
	a, err := app.New(c, logger, app.RPCServerOption(rs))
	if err != nil {
		return nil, err
	}

	return a, nil
}

var ProviderSet = wire.NewSet(NewApp)
