package machinery

import (
	"fmt"
	"github.com/RichardKnop/machinery/v2"
	amqpBackend "github.com/RichardKnop/machinery/v2/backends/amqp"
	redisBackend "github.com/RichardKnop/machinery/v2/backends/redis"
	amqpBroker "github.com/RichardKnop/machinery/v2/brokers/amqp"
	redisBroker "github.com/RichardKnop/machinery/v2/brokers/redis"
	"github.com/RichardKnop/machinery/v2/config"
	eagerLock "github.com/RichardKnop/machinery/v2/locks/eager"
	"github.com/google/wire"
	appConfig "github.com/tsundata/assistant/internal/pkg/config"
)

const DefaultQueue = "assistant_tasks"
const AMQPExchange = "machinery_exchange"
const AMQPBindingKey = "machinery_task"

func New(c *appConfig.AppConfig) (*machinery.Server, error) {
	// use rabbitmq or redis
	if c.Rabbitmq.Url != "" {
		cnf := &config.Config{
			DefaultQueue:    DefaultQueue,
			ResultsExpireIn: 3600,
			Broker:          c.Rabbitmq.Url,
			ResultBackend:   c.Rabbitmq.Url,
			AMQP: &config.AMQPConfig{
				Exchange:     AMQPExchange,
				ExchangeType: "direct",
				BindingKey:   AMQPBindingKey,
			},
		}

		server := machinery.NewServer(cnf, amqpBroker.New(cnf), amqpBackend.New(cnf), eagerLock.New())
		return server, nil
	} else {
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
}

var ProviderSet = wire.NewSet(New)
