package agent

import (
	"github.com/google/wire"
	"github.com/tsundata/assistant/internal/app/agent/broker"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/logger"
)

func NewApp(c *config.AppConfig, logger *logger.Logger, b broker.Runner) (*app.Application, error) {
	logger.Info("start agent " + c.Name)

	a, err := app.New(c, logger)
	if err != nil {
		return nil, err
	}

	b.Run()

	return a, nil
}

var ProviderSet = wire.NewSet(NewApp)
