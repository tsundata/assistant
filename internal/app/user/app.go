package user

import (
	"context"
	"github.com/google/wire"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/user/listener"
	"github.com/tsundata/assistant/internal/app/user/repository"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
)

func NewApp(c *config.AppConfig, bus event.Bus, logger log.Logger, rs *rpc.Server, repo repository.UserRepository, nlpClient pb.NLPSvcClient) (*app.Application, error) {
	// event bus register
	err := listener.RegisterEventHandler(context.Background(), bus, logger, repo, nlpClient)
	if err != nil {
		return nil, err
	}

	a, err := app.New(c, app.RPCServerOption(rs))
	if err != nil {
		return nil, err
	}

	return a, nil
}

var ProviderSet = wire.NewSet(NewApp)
