package agent

import (
	"github.com/google/wire"
	"github.com/tsundata/assistant/internal/app/agent/broker"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"os"
)

func NewApp(logger *logger.Logger, b broker.Runner) (*app.Application, error) {
	name := os.Getenv("APP_NAME")

	logger.Info("start agent " + name)

	a, err := app.New(name, logger)
	if err != nil {
		return nil, err
	}

	b.Run()

	return a, nil
}

var ProviderSet = wire.NewSet(NewApp)
