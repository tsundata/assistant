package spider

import (
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/spider/crawler"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/log"
)

func NewApp(c *config.AppConfig, rdb *redis.Client, bus event.Bus, logger log.Logger,
	middle pb.MiddleSvcClient, message pb.MessageSvcClient, user pb.UserSvcClient) (*app.Application, error) {
	// spider
	go func() {
		s := crawler.New()
		s.SetService(c, rdb, bus, logger, middle, message, user)
		err := s.LoadRule()
		if err != nil {
			logger.Error(err)
			return
		}
		s.Daemon()
	}()

	a, err := app.New(c, logger)
	if err != nil {
		return nil, err
	}

	return a, nil
}

var ProviderSet = wire.NewSet(NewApp)
