// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

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
	"github.com/tsundata/assistant/internal/pkg/middleware/rqlite"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
	"github.com/tsundata/assistant/internal/pkg/vendors/newrelic"
	"github.com/tsundata/assistant/internal/pkg/vendors/rollbar"
)

// Injectors from wire.go:

func CreateApp(id string) (*app.Application, error) {
	client, err := consul.New()
	if err != nil {
		return nil, err
	}
	appConfig := config.NewConfig(id, client)
	rollbarRollbar := rollbar.New(appConfig)
	logger := log.NewZapLogger(rollbarRollbar)
	logLogger := log.NewAppLogger(logger)
	conn, err := nats.New(appConfig)
	if err != nil {
		return nil, err
	}
	newrelicApp, err := newrelic.New(appConfig, logger)
	if err != nil {
		return nil, err
	}
	bus := event.NewNatsBus(conn, newrelicApp)
	rqliteConn, err := rqlite.New(appConfig, newrelicApp)
	if err != nil {
		return nil, err
	}
	todoRepository := repository.NewRqliteTodoRepository(logLogger, rqliteConn)
	serviceTodo := service.NewTodo(bus, logLogger, todoRepository)
	initServer := service.CreateInitServerFn(serviceTodo)
	configuration, err := jaeger.NewConfiguration(appConfig, logLogger)
	if err != nil {
		return nil, err
	}
	tracer, err := jaeger.New(configuration)
	if err != nil {
		return nil, err
	}
	redisClient, err := redis.New(appConfig, newrelicApp)
	if err != nil {
		return nil, err
	}
	server, err := rpc.NewServer(appConfig, logger, logLogger, initServer, tracer, redisClient, client, newrelicApp)
	if err != nil {
		return nil, err
	}
	application, err := todo.NewApp(appConfig, logLogger, server)
	if err != nil {
		return nil, err
	}
	return application, nil
}

// wire.go:

var providerSet = wire.NewSet(config.ProviderSet, log.ProviderSet, rpc.ProviderSet, jaeger.ProviderSet, influx.ProviderSet, redis.ProviderSet, todo.ProviderSet, mysql.ProviderSet, rollbar.ProviderSet, consul.ProviderSet, repository.ProviderSet, event.ProviderSet, nats.ProviderSet, service.ProviderSet, rqlite.ProviderSet, newrelic.ProviderSet)
