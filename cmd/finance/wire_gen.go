// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/tsundata/assistant/internal/app/finance"
	"github.com/tsundata/assistant/internal/app/finance/service"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/middleware/consul"
	"github.com/tsundata/assistant/internal/pkg/middleware/influx"
	"github.com/tsundata/assistant/internal/pkg/middleware/jaeger"
	"github.com/tsundata/assistant/internal/pkg/middleware/mysql"
	"github.com/tsundata/assistant/internal/pkg/middleware/redis"
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
	rollbarRollbar := rollbar.New(appConfig)
	logger := log.NewZapLogger(rollbarRollbar)
	serviceFinance := service.NewFinance()
	initServer := service.CreateInitServerFn(serviceFinance)
	configuration, err := jaeger.NewConfiguration(appConfig, logger)
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
	redisClient, err := redis.New(appConfig)
	if err != nil {
		return nil, err
	}
	server, err := rpc.NewServer(appConfig, logger, initServer, tracer, influxdb2Client, redisClient, client)
	if err != nil {
		return nil, err
	}
	application, err := finance.NewApp(appConfig, logger, server)
	if err != nil {
		return nil, err
	}
	return application, nil
}

// wire.go:

var providerSet = wire.NewSet(config.ProviderSet, log.ProviderSet, rpc.ProviderSet, jaeger.ProviderSet, influx.ProviderSet, redis.ProviderSet, finance.ProviderSet, mysql.ProviderSet, rollbar.ProviderSet, consul.ProviderSet, service.ProviderSet)
