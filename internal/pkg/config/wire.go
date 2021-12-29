//go:build wireinject
// +build wireinject

package config

import (
	"github.com/google/wire"
	"github.com/tsundata/assistant/internal/pkg/middleware/etcd"
)

var testProviderSet = wire.NewSet(
	etcd.ProviderSet,
	ProviderSet,
)

func CreateAppConfig(id string) (*AppConfig, error) {
	panic(wire.Build(testProviderSet))
}
