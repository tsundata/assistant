package cron

import (
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/tsundata/assistant/internal/app/cron/rule"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
)

func NewApp(c *config.AppConfig, logger *logger.Logger, rdb *redis.Client, client *rpc.Client) (*app.Application, error) {

	b := rulebot.New(&rulebot.Context{Conf: c, RDB: rdb, Client: client}, rule.Options...)

	logger.Info("start cron bot " + b.Name())

	a, err := app.New(c, logger)
	if err != nil {
		return nil, err
	}

	return a, nil
}

var ProviderSet = wire.NewSet(NewApp)
