package rabbitmq

import (
	"github.com/google/wire"
	"github.com/pkg/errors"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/tsundata/assistant/internal/pkg/config"
)

func New(c *config.AppConfig) (*amqp.Connection, error) {
	conn, err := amqp.Dial(c.Rabbitmq.Url)
	if err != nil {
		return nil, errors.Wrap(err, "rabbitmq server error")
	}
	return conn, nil
}

var ProviderSet = wire.NewSet(New)
