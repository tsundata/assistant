// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/tsundata/assistant/internal/app/id"
	"github.com/tsundata/assistant/internal/app/id/service"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/middleware/jaeger"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
)

var providerSet = wire.NewSet(
	config.ProviderSet,
	jaeger.ProviderSet,
	rpc.ProviderSet,
	id.ProviderSet,
	service.ProviderSet,
)

func CreateApp(id string) (*app.Application, error) {
	panic(wire.Build(providerSet))
}
