package nats

import (
	"github.com/google/wire"
	"github.com/nats-io/nats.go"
	"github.com/tsundata/assistant/internal/pkg/config"
)

func New(c *config.AppConfig) (*nats.Conn, error) {
	nc, err := nats.Connect(c.Nats.Url)
	if err != nil {
		return nil, err
	}
	return nc, nil
}

var ProviderSet = wire.NewSet(New)
