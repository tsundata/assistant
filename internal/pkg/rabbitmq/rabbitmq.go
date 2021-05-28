package rabbitmq

import (
	"github.com/google/wire"
	"github.com/streadway/amqp"
	"github.com/tsundata/assistant/internal/pkg/config"
)

func New(c *config.AppConfig) (*amqp.Connection, error) {
	return amqp.Dial(c.Rabbitmq.Url)
}

var ProviderSet = wire.NewSet(New)
