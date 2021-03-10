package agent

import (
	"github.com/tsundata/assistant/internal/app/agent/broker"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/logger"
)

func NewApp(name string, logger *logger.Logger, b broker.Runner) (*app.Application, error) {
	logger.Info("start agent " + name)

	a, err := app.New(name, logger)
	if err != nil {
		return nil, err
	}

	b.Run()

	return a, nil
}
