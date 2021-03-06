package cron

import (
	"github.com/google/wire"
	"github.com/tsundata/assistant/internal/app/cron/rule"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
)

func NewApp(c *config.AppConfig, logger *logger.Logger, bot *rulebot.RuleBot) (*app.Application, error) {
	// load rule
	bot.SetOptions(rule.Options...)
	logger.Info("start cron rule bot")

	a, err := app.New(c, logger)
	if err != nil {
		return nil, err
	}

	return a, nil
}

var ProviderSet = wire.NewSet(NewApp)
