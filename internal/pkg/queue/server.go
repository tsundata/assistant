package queue

import (
	"fmt"
	"github.com/RichardKnop/machinery/v2"
	redisBackend "github.com/RichardKnop/machinery/v2/backends/redis"
	redisBroker "github.com/RichardKnop/machinery/v2/brokers/redis"
	"github.com/RichardKnop/machinery/v2/config"
	eagerLock "github.com/RichardKnop/machinery/v2/locks/eager"
	"github.com/google/wire"
	appConfig "github.com/tsundata/assistant/internal/pkg/config"
)

const DefaultQueue = "assistant_tasks"

func New(c *appConfig.AppConfig) (*machinery.Server, error) {
	cnf := &config.Config{
		DefaultQueue:    DefaultQueue,
		ResultsExpireIn: 3600,
		Redis: &config.RedisConfig{
			MaxIdle:                3,
			IdleTimeout:            240,
			ReadTimeout:            15,
			WriteTimeout:           15,
			ConnectTimeout:         15,
			NormalTasksPollPeriod:  1000,
			DelayedTasksPollPeriod: 500,
		},
	}
	broker := redisBroker.NewGR(cnf, []string{fmt.Sprintf("%s@%s", c.Redis.Password, c.Redis.Addr)}, 0)
	backend := redisBackend.NewGR(cnf, []string{fmt.Sprintf("%s@%s", c.Redis.Password, c.Redis.Addr)}, 0)
	lock := eagerLock.New()

	server := machinery.NewServer(cnf, broker, backend, lock)
	return server, nil
}

var ProviderSet = wire.NewSet(New)
