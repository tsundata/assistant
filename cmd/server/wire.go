//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	server "github.com/tsundata/assistant/internal/app"
	"github.com/tsundata/assistant/internal/app/controller"
	"github.com/tsundata/assistant/internal/app/repository"
	"github.com/tsundata/assistant/internal/app/service"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/global"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/middleware/etcd"
	"github.com/tsundata/assistant/internal/pkg/middleware/influx"
	"github.com/tsundata/assistant/internal/pkg/middleware/jaeger"
	"github.com/tsundata/assistant/internal/pkg/middleware/mysql"
	"github.com/tsundata/assistant/internal/pkg/middleware/rabbitmq"
	"github.com/tsundata/assistant/internal/pkg/middleware/redis"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
	"github.com/tsundata/assistant/internal/pkg/transport/http"
	"github.com/tsundata/assistant/internal/pkg/vendors/newrelic"
	"github.com/tsundata/assistant/internal/pkg/vendors/rollbar"
)

var providerSet = wire.NewSet(
	config.ProviderSet,
	log.ProviderSet,
	http.ProviderSet,
	jaeger.ProviderSet,
	influx.ProviderSet,
	redis.ProviderSet,
	controller.ProviderSet,
	server.ProviderSet,
	rollbar.ProviderSet,
	event.ProviderSet,
	etcd.ProviderSet,
	newrelic.ProviderSet,
	rabbitmq.ProviderSet,
	service.ProviderSet,
	repository.ProviderSet,
	global.ProviderSet,
	mysql.ProviderSet,
	component.ProviderSet,
)

func CreateApp(id string) (*app.Application, error) {
	panic(wire.Build(providerSet))
}
