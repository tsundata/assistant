// +build wireinject

package repository

import (
	"github.com/google/wire"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/middleware/consul"
	"github.com/tsundata/assistant/internal/pkg/middleware/rqlite"
	"github.com/tsundata/assistant/internal/pkg/vendors/newrelic"
	"github.com/tsundata/assistant/internal/pkg/vendors/rollbar"
)

var testProviderSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	consul.ProviderSet,
	ProviderSet,
	rollbar.ProviderSet,
	rqlite.ProviderSet,
	newrelic.ProviderSet,
)

func CreateTodoRepository(id string) (TodoRepository, error) {
	panic(wire.Build(testProviderSet))
}
