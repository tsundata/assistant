package queue

import (
	"github.com/RichardKnop/machinery/v2"
	amqpBackend "github.com/RichardKnop/machinery/v2/backends/amqp"
	amqpBroker "github.com/RichardKnop/machinery/v2/brokers/amqp"
	"github.com/RichardKnop/machinery/v2/config"
	eagerLock "github.com/RichardKnop/machinery/v2/locks/eager"
	"github.com/google/wire"
	appConfig "github.com/tsundata/assistant/internal/pkg/config"
)

const DefaultQueue = "queue_tasks"
const DefaultExchange = "queue_exchange"

func New(c *appConfig.AppConfig) (*machinery.Server, error) {
	cnf := &config.Config{
		Broker:          c.Rabbitmq.Url,
		ResultBackend:   c.Rabbitmq.Url,
		DefaultQueue:    DefaultQueue,
		ResultsExpireIn: 3600,
		AMQP: &config.AMQPConfig{
			Exchange:         DefaultExchange,
			ExchangeType:     "direct",
			QueueDeclareArgs: nil,
			QueueBindingArgs: nil,
			BindingKey:       DefaultQueue,
			PrefetchCount:    1,
			AutoDelete:       false,
		},
	}
	broker := amqpBroker.New(cnf)
	backend := amqpBackend.New(cnf)
	lock := eagerLock.New()

	server := machinery.NewServer(cnf, broker, backend, lock)
	return server, nil
}

var ProviderSet = wire.NewSet(New)
