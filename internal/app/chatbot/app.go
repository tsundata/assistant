package chatbot

import (
	"context"
	"github.com/google/wire"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/bot/finance"
	"github.com/tsundata/assistant/internal/app/bot/org"
	"github.com/tsundata/assistant/internal/app/bot/todo"
	"github.com/tsundata/assistant/internal/app/chatbot/listener"
	"github.com/tsundata/assistant/internal/app/chatbot/repository"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/robot"
	"github.com/tsundata/assistant/internal/pkg/robot/rulebot"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
)

func NewApp(c *config.AppConfig, bus event.Bus, logger log.Logger, rs *rpc.Server,
	message pb.MessageSvcClient, repo repository.ChatbotRepository, bot *rulebot.RuleBot,
) (*app.Application, error) {
	// event bus register
	err := listener.RegisterEventHandler(bus, logger, bot, message, repo)
	if err != nil {
		return nil, err
	}

	// bots register
	err = robot.RegisterBot(context.Background(), bus, todo.Bot, org.Bot, finance.Bot)
	if err != nil {
		return nil, err
	}

	a, err := app.New(c, logger, app.RPCServerOption(rs))
	if err != nil {
		return nil, err
	}

	return a, nil
}

var ProviderSet = wire.NewSet(NewApp)
