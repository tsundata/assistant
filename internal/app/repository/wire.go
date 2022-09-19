//go:build wireinject
// +build wireinject

package repository

import (
	"github.com/google/wire"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/middleware/etcd"
	"github.com/tsundata/assistant/internal/pkg/middleware/mysql"
	"github.com/tsundata/assistant/internal/pkg/vendors/newrelic"
	"github.com/tsundata/assistant/internal/pkg/vendors/rollbar"
)

var testProviderSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	etcd.ProviderSet,
	ProviderSet,
	rollbar.ProviderSet,
	mysql.ProviderSet,
	newrelic.ProviderSet,
)

func CreateMiddleRepository(id string) (MiddleRepository, error) {
	panic(wire.Build(testProviderSet))
}

func CreateChatbotRepository(id string) (ChatbotRepository, error) {
	panic(wire.Build(testProviderSet))
}

func CreateIdRepository(id string) (IdRepository, error) {
	panic(wire.Build(testProviderSet))
}

func CreateMessageRepository(id string) (MessageRepository, error) {
	panic(wire.Build(testProviderSet))
}

func CreateUserRepository(id string) (UserRepository, error) {
	panic(wire.Build(testProviderSet))
}
