package cron

import (
	"github.com/google/wire"
	"github.com/tsundata/assistant/internal/app/cron/rule"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/robot/rulebot"
)

func NewApp(c *config.AppConfig, logger log.Logger, bot *rulebot.RuleBot) (*app.Application, error) {
	// cron
	go func() {
		// Delayed loading
		//time.Sleep(1 * time.Minute)
		// load rule
		bot.SetOptions(rule.Options...)
		logger.Info("start cron rule bot")
	}()

	a, err := app.New(c, logger)
	if err != nil {
		return nil, err
	}

	return a, nil
}

var ProviderSet = wire.NewSet(NewApp)
