// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/tsundata/assistant/internal/app/user"
	"github.com/tsundata/assistant/internal/app/user/repository"
	"github.com/tsundata/assistant/internal/app/user/service"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/middleware/consul"
	"github.com/tsundata/assistant/internal/pkg/middleware/influx"
	"github.com/tsundata/assistant/internal/pkg/middleware/jaeger"
	"github.com/tsundata/assistant/internal/pkg/middleware/mysql"
	"github.com/tsundata/assistant/internal/pkg/middleware/nats"
	"github.com/tsundata/assistant/internal/pkg/middleware/redis"
	"github.com/tsundata/assistant/internal/pkg/transport/http"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
	"github.com/tsundata/assistant/internal/pkg/vendors/rollbar"
)

// Injectors from wire.go:

func CreateApp(id string) (*app.Application, error) {
	client, err := consul.New()
	if err != nil {
		return nil, err
	}
	appConfig := config.NewConfig(id, client)
	conn, err := nats.New(appConfig)
	if err != nil {
		return nil, err
	}
	bus := event.NewBus(conn)
	rollbarRollbar := rollbar.New(appConfig)
	loggerLogger := logger.NewLogger(rollbarRollbar)
	redisClient, err := redis.New(appConfig)
	if err != nil {
		return nil, err
	}
	db, err := mysql.New(appConfig)
	if err != nil {
		return nil, err
	}
	userRepository := repository.NewMysqlUserRepository(loggerLogger, db)
	serviceUser := service.NewUser(redisClient, userRepository)
	initServer := service.CreateInitServerFn(serviceUser)
	configuration, err := jaeger.NewConfiguration(appConfig, loggerLogger)
	if err != nil {
		return nil, err
	}
	tracer, err := jaeger.New(configuration)
	if err != nil {
		return nil, err
	}
	influxdb2Client, err := influx.New(appConfig)
	if err != nil {
		return nil, err
	}
	server, err := rpc.NewServer(appConfig, loggerLogger, initServer, tracer, influxdb2Client, redisClient, client)
	if err != nil {
		return nil, err
	}
	application, err := user.NewApp(appConfig, bus, loggerLogger, server, userRepository)
	if err != nil {
		return nil, err
	}
	return application, nil
}

// wire.go:

var providerSet = wire.NewSet(config.ProviderSet, logger.ProviderSet, http.ProviderSet, rpc.ProviderSet, jaeger.ProviderSet, influx.ProviderSet, redis.ProviderSet, user.ProviderSet, mysql.ProviderSet, rollbar.ProviderSet, consul.ProviderSet, repository.ProviderSet, event.ProviderSet, nats.ProviderSet, service.ProviderSet)
