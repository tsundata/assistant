package message

import (
	"github.com/google/wire"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/message/listener"
	"github.com/tsundata/assistant/internal/app/message/repository"
	"github.com/tsundata/assistant/internal/app/message/rules"
	"github.com/tsundata/assistant/internal/app/message/service"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc"
	"google.golang.org/grpc"
)

func NewApp(c *config.AppConfig, bus *event.Bus, logger *logger.Logger, rs *rpc.Server,
	repo repository.MessageRepository, client *rpc.Client) (*app.Application, error) {

	// event bus register
	err := listener.RegisterEventHandler(bus)
	if err != nil {
		return nil, err
	}

	// rule bot
	bot := rulebot.New(c, nil, client, rules.Options...)

	// rpc service
	s := service.NewManage(logger, bot, c.Slack.Webhook, repo, client)
	err = rs.Register(func(gs *grpc.Server) error {
		pb.RegisterMessageServer(gs, s)
		return nil
	})
	if err != nil {
		return nil, err
	}

	// rpc server
	a, err := app.New(c, logger, app.RPCServerOption(rs))
	if err != nil {
		return nil, err
	}

	return a, nil
}

var ProviderSet = wire.NewSet(NewApp)
