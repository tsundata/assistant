package nats

import (
	"github.com/google/wire"
	"github.com/nats-io/nats.go"
	"github.com/pkg/errors"
	"github.com/tsundata/assistant/internal/pkg/config"
)

func New(c *config.AppConfig) (*nats.Conn, error) {
	nc, err := nats.Connect(c.Nats.Url)
	if err != nil {
		return nil, errors.Wrap(err, "nats server error")
	}
	return nc, nil
}

var ProviderSet = wire.NewSet(New)
