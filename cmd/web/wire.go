// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/tsundata/assistant/internal/app/web"
	"github.com/tsundata/assistant/internal/app/web/controller"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/middleware/consul"
	"github.com/tsundata/assistant/internal/pkg/middleware/influx"
	"github.com/tsundata/assistant/internal/pkg/middleware/jaeger"
	"github.com/tsundata/assistant/internal/pkg/middleware/redis"
	"github.com/tsundata/assistant/internal/pkg/sdk"
	"github.com/tsundata/assistant/internal/pkg/transport/http"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
	"github.com/tsundata/assistant/internal/pkg/vendors/rollbar"
)

var providerSet = wire.NewSet(
	config.ProviderSet,
	log.ProviderSet,
	http.ProviderSet,
	rpc.ProviderSet,
	jaeger.ProviderSet,
	influx.ProviderSet,
	redis.ProviderSet,
	controller.ProviderSet,
	web.ProviderSet,
	rollbar.ProviderSet,
	consul.ProviderSet,
	sdk.ProviderSet,
)

func CreateApp(id string) (*app.Application, error) {
	panic(wire.Build(providerSet))
}
