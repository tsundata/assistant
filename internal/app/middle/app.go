package middle

import (
	"github.com/google/wire"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
)

func NewApp(c *config.AppConfig, logger *logger.Logger, rs *rpc.Server) (*app.Application, error) {
	a, err := app.New(c, logger, app.RPCServerOption(rs))
	if err != nil {
		return nil, err
	}

	return a, nil
}

var ProviderSet = wire.NewSet(NewApp)
