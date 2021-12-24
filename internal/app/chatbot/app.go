package chatbot

import (
	"github.com/google/wire"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/chatbot/listener"
	"github.com/tsundata/assistant/internal/app/chatbot/repository"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/robot/rulebot"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
)

func NewApp(c *config.AppConfig, bus event.Bus, logger log.Logger, rs *rpc.Server,
	message pb.MessageSvcClient, middle pb.MiddleSvcClient, todo pb.TodoSvcClient, user pb.UserSvcClient,
	repo repository.ChatbotRepository, bot *rulebot.RuleBot,
) (*app.Application, error) {
	// event bus register
	err := listener.RegisterEventHandler(bus, logger, bot, message, middle, todo, user, repo)
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
