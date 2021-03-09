package cron

import (
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
	"go.uber.org/zap"
)

func NewApp(name string, logger *zap.Logger, b *rulebot.RuleBot) (*app.Application, error) {
	logger.Info("start cron bot " + b.Name())

	a, err := app.New(name, logger)
	if err != nil {
		return nil, err
	}

	return a, nil
}
