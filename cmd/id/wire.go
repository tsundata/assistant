//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/tsundata/assistant/internal/app/id"
	"github.com/tsundata/assistant/internal/app/id/repository"
	"github.com/tsundata/assistant/internal/app/id/service"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/middleware/etcd"
	"github.com/tsundata/assistant/internal/pkg/middleware/influx"
	"github.com/tsundata/assistant/internal/pkg/middleware/jaeger"
	"github.com/tsundata/assistant/internal/pkg/middleware/mysql"
	"github.com/tsundata/assistant/internal/pkg/middleware/redis"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
	"github.com/tsundata/assistant/internal/pkg/vendors/newrelic"
	"github.com/tsundata/assistant/internal/pkg/vendors/rollbar"
)

var providerSet = wire.NewSet(
	config.ProviderSet,
	log.ProviderSet,
	rpc.ProviderSet,
	jaeger.ProviderSet,
	influx.ProviderSet,
	redis.ProviderSet,
	id.ProviderSet,
	rollbar.ProviderSet,
	etcd.ProviderSet,
	service.ProviderSet,
	newrelic.ProviderSet,
	repository.ProviderSet,
	mysql.ProviderSet,
)

func CreateApp(id string) (*app.Application, error) {
	panic(wire.Build(providerSet))
}
