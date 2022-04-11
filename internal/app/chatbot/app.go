package chatbot

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/tsundata/assistant/api/pb"
	finance2 "github.com/tsundata/assistant/internal/app/chatbot/bot/finance"
	"github.com/tsundata/assistant/internal/app/chatbot/bot/github"
	org2 "github.com/tsundata/assistant/internal/app/chatbot/bot/org"
	system2 "github.com/tsundata/assistant/internal/app/chatbot/bot/system"
	todo2 "github.com/tsundata/assistant/internal/app/chatbot/bot/todo"
	"github.com/tsundata/assistant/internal/app/chatbot/listener"
	"github.com/tsundata/assistant/internal/app/chatbot/repository"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/robot"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
	"github.com/tsundata/assistant/internal/pkg/robot/rulebot"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
)

func NewApp(conf *config.AppConfig, bus event.Bus, rdb *redis.Client, logger log.Logger, rs *rpc.Server,
	message pb.MessageSvcClient, middle pb.MiddleSvcClient, repo repository.ChatbotRepository, bot *rulebot.RuleBot, comp component.Component,
) (*app.Application, error) {
	// event bus register
	err := listener.RegisterEventHandler(conf, bus, rdb, logger, bot, message, middle, repo, comp)
	if err != nil {
		return nil, err
	}

	// bots register
	err = robot.RegisterBot(context.Background(), bus, system2.Bot, todo2.Bot, org2.Bot, finance2.Bot, github.Bot)
	if err != nil {
		return nil, err
	}

	a, err := app.New(conf, logger, app.RPCServerOption(rs))
	if err != nil {
		return nil, err
	}

	return a, nil
}

var ProviderSet = wire.NewSet(NewApp)
