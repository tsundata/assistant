package cron

import (
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/cron/rules"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
	"os"
)

func NewApp(logger *logger.Logger, rdb *redis.Client, subClient pb.SubscribeClient,
	midClient pb.MiddleClient, msgClient pb.MessageClient, wfClient pb.WorkflowClient) (*app.Application, error) {
	name := os.Getenv("APP_NAME")

	b := rulebot.New(rdb, subClient, midClient, msgClient, wfClient, nil, rules.Options...)

	logger.Info("start cron bot " + b.Name())

	a, err := app.New(name, logger)
	if err != nil {
		return nil, err
	}

	return a, nil
}

var ProviderSet = wire.NewSet(NewApp)
