// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/tsundata/assistant/internal/app/web"
	"github.com/tsundata/assistant/internal/app/web/controllers"
	"github.com/tsundata/assistant/internal/app/web/rpcclients"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/middleware/consul"
	"github.com/tsundata/assistant/internal/pkg/middleware/etcd"
	"github.com/tsundata/assistant/internal/pkg/middleware/influx"
	"github.com/tsundata/assistant/internal/pkg/middleware/jaeger"
	"github.com/tsundata/assistant/internal/pkg/middleware/redis"
	"github.com/tsundata/assistant/internal/pkg/transports/http"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc"
	"github.com/tsundata/assistant/internal/pkg/vendors/rollbar"
)

var providerSet = wire.NewSet(
	config.ProviderSet,
	logger.ProviderSet,
	http.ProviderSet,
	rpc.ProviderSet,
	jaeger.ProviderSet,
	etcd.ProviderSet,
	influx.ProviderSet,
	rpcclients.ProviderSet,
	redis.ProviderSet,
	controllers.ProviderSet,
	web.ProviderSet,
	rollbar.ProviderSet,
	consul.ProviderSet,
)

func CreateApp() (*app.Application, error) {
	panic(wire.Build(providerSet))
}
