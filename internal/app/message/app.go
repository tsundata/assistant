package message

import (
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/message/listener"
	"github.com/tsundata/assistant/internal/app/message/repository"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
)

func NewApp(c *config.AppConfig, bus event.Bus, logger log.Logger, redis *redis.Client, rs *rpc.Server, repo repository.MessageRepository,
	chatbot pb.ChatbotSvcClient, storage pb.StorageSvcClient, middle pb.MiddleSvcClient) (*app.Application, error) {
	// event bus register
	err := listener.RegisterEventHandler(bus, c, logger, redis, repo, chatbot, storage, middle)
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
