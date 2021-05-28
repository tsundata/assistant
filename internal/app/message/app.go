package message

import (
	"github.com/google/wire"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/message/repository"
	"github.com/tsundata/assistant/internal/app/message/rules"
	"github.com/tsundata/assistant/internal/app/message/service"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc"
	"google.golang.org/grpc"
)

func NewApp(c *config.AppConfig, logger *logger.Logger, rs *rpc.Server, repo repository.MessageRepository,
	subClient pb.SubscribeClient, midClient pb.MiddleClient, msgClient pb.MessageClient,
	taskClient pb.TaskClient, wfClient pb.WorkflowClient, storageClient pb.StorageClient) (*app.Application, error) {

	b := rulebot.New(c, nil, subClient, midClient, msgClient, wfClient, taskClient, storageClient, rules.Options...)

	s := service.NewManage(logger, b, c.Slack.Webhook, repo, wfClient, msgClient, midClient)
	err := rs.Register(func(gs *grpc.Server) error {
		pb.RegisterMessageServer(gs, s)
		return nil
	})
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
