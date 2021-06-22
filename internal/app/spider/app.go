package spider

import (
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/spider/crawler"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"time"
)

func NewApp(
	c *config.AppConfig,
	rdb *redis.Client,
	logger *logger.Logger,
	subscribe pb.SubscribeClient,
	middle pb.MiddleClient,
	message pb.MessageClient) (*app.Application, error) {

	// spider
	go func() {
		// Delayed loading
		time.Sleep(10 * time.Second)
		s := crawler.New()
		s.SetService(c, rdb, logger, subscribe, middle, message)
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
