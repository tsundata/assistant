// +build wireinject

package event

import (
	"github.com/google/wire"
	"github.com/nats-io/nats.go"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/middleware/consul"
	natsMiddle "github.com/tsundata/assistant/internal/pkg/middleware/nats"
)

var testProviderSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	consul.ProviderSet,
	ProviderSet,
	natsMiddle.ProviderSet,
)

func CreateNats(id string) (*nats.Conn, error) {
	panic(wire.Build(testProviderSet))
}
