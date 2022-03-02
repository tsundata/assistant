//go:build wireinject
// +build wireinject

package event

import (
	"github.com/google/wire"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/middleware/etcd"
	natsMiddle "github.com/tsundata/assistant/internal/pkg/middleware/nats"
	"github.com/tsundata/assistant/internal/pkg/middleware/rabbitmq"
)

var testProviderSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	etcd.ProviderSet,
	ProviderSet,
	natsMiddle.ProviderSet,
	rabbitmq.ProviderSet,
)

func CreateRabbitmq(id string) (*amqp.Connection, error) {
	panic(wire.Build(testProviderSet))
}
