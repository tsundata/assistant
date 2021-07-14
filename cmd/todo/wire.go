// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/tsundata/assistant/internal/app/todo"
	"github.com/tsundata/assistant/internal/app/todo/repository"
	"github.com/tsundata/assistant/internal/app/todo/service"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/middleware/consul"
	"github.com/tsundata/assistant/internal/pkg/middleware/influx"
	"github.com/tsundata/assistant/internal/pkg/middleware/jaeger"
	"github.com/tsundata/assistant/internal/pkg/middleware/mysql"
	"github.com/tsundata/assistant/internal/pkg/middleware/nats"
	"github.com/tsundata/assistant/internal/pkg/middleware/redis"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
	"github.com/tsundata/assistant/internal/pkg/vendors/rollbar"
)

var providerSet = wire.NewSet(
	config.ProviderSet,
	log.ProviderSet,
	rpc.ProviderSet,
	jaeger.ProviderSet,
	influx.ProviderSet,
	redis.ProviderSet,
	todo.ProviderSet,
	mysql.ProviderSet,
	rollbar.ProviderSet,
	consul.ProviderSet,
	repository.ProviderSet,
	event.ProviderSet,
	nats.ProviderSet,
	service.ProviderSet,
)

func CreateApp(id string) (*app.Application, error) {
	panic(wire.Build(providerSet))
}
