// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/tsundata/assistant/internal/app/message"
	"github.com/tsundata/assistant/internal/app/message/repository"
	"github.com/tsundata/assistant/internal/app/message/rpcclient"
	"github.com/tsundata/assistant/internal/app/message/service"
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
	db, err := mysql.New(appConfig)
	if err != nil {
		return nil, err
	}
	messageRepository := repository.NewMysqlMessageRepository(loggerLogger, db)
	configuration, err := jaeger.NewConfiguration(appConfig, loggerLogger)
	if err != nil {
		return nil, err
	}
	tracer, err := jaeger.New(configuration)
	if err != nil {
		return nil, err
	}
	clientOptions, err := rpc.NewClientOptions(appConfig, tracer)
	if err != nil {
		return nil, err
	}
	rpcClient, err := rpc.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}
	serviceMessage := service.NewMessage(loggerLogger, appConfig, messageRepository, rpcClient)
	initServer := service.CreateInitServerFn(serviceMessage)
	influxdb2Client, err := influx.New(appConfig)
	if err != nil {
		return nil, err
	}
	redisClient, err := redis.New(appConfig)
	if err != nil {
		return nil, err
	}
	server, err := rpc.NewServer(appConfig, loggerLogger, initServer, tracer, influxdb2Client, redisClient, client)
	if err != nil {
		return nil, err
	}
	application, err := message.NewApp(appConfig, bus, loggerLogger, server, rpcClient)
	if err != nil {
		return nil, err
	}
	return application, nil
}

// wire.go:

var providerSet = wire.NewSet(config.ProviderSet, logger.ProviderSet, http.ProviderSet, rpc.ProviderSet, jaeger.ProviderSet, influx.ProviderSet, redis.ProviderSet, message.ProviderSet, mysql.ProviderSet, rollbar.ProviderSet, repository.ProviderSet, nats.ProviderSet, event.ProviderSet, consul.ProviderSet, service.ProviderSet, rpcclient.ProviderSet)
