// +build wireinject

package config

import (
	"github.com/google/wire"
	"github.com/tsundata/assistant/internal/pkg/middleware/consul"
)

var testProviderSet = wire.NewSet(
	consul.ProviderSet,
	ProviderSet,
)

func CreateAppConfig(id string) (*AppConfig, error) {
	panic(wire.Build(testProviderSet))
}
