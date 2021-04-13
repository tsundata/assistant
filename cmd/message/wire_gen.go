// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/tsundata/assistant/internal/app/message"
	"github.com/tsundata/assistant/internal/app/message/rpcclients"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/database"
	"github.com/tsundata/assistant/internal/pkg/etcd"
	"github.com/tsundata/assistant/internal/pkg/influx"
	"github.com/tsundata/assistant/internal/pkg/jaeger"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/redis"
	"github.com/tsundata/assistant/internal/pkg/transports/http"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc"
	"github.com/tsundata/assistant/internal/pkg/vendors/rollbar"
)

// Injectors from wire.go:

func CreateApp(cf string) (*app.Application, error) {
	viper, err := config.New(cf)
	if err != nil {
		return nil, err
	}
	options, err := message.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	rollbarOptions, err := rollbar.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	rollbarRollbar := rollbar.New(rollbarOptions)
	loggerLogger := logger.NewLogger(rollbarRollbar)
	rpcOptions, err := rpc.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	configuration, err := jaeger.NewConfiguration(viper, loggerLogger)
	if err != nil {
		return nil, err
	}
	tracer, err := jaeger.New(configuration)
	if err != nil {
		return nil, err
	}
	etcdOptions, err := etcd.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	client, err := etcd.New(etcdOptions)
	if err != nil {
		return nil, err
	}
	influxOptions, err := influx.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	influxdb2Client, err := influx.New(influxOptions)
	if err != nil {
		return nil, err
	}
	redisOptions, err := redis.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	redisClient, err := redis.New(redisOptions)
	if err != nil {
		return nil, err
	}
	server, err := rpc.NewServer(rpcOptions, loggerLogger, tracer, client, influxdb2Client, redisClient)
	if err != nil {
		return nil, err
	}
	databaseOptions, err := database.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	db, err := database.New(databaseOptions)
	if err != nil {
		return nil, err
	}
	clientOptions, err := rpc.NewClientOptions(viper, tracer)
	if err != nil {
		return nil, err
	}
	rpcClient, err := rpc.NewClient(clientOptions, client)
	if err != nil {
		return nil, err
	}
	subscribeClient, err := rpcclients.NewSubscribeClient(rpcClient)
	if err != nil {
		return nil, err
	}
	middleClient, err := rpcclients.NewMiddleClient(rpcClient)
	if err != nil {
		return nil, err
	}
	messageClient, err := rpcclients.NewMessageClient(rpcClient)
	if err != nil {
		return nil, err
	}
	taskClient, err := rpcclients.NewTaskClient(rpcClient)
	if err != nil {
		return nil, err
	}
	workflowClient, err := rpcclients.NewWorkflowClient(rpcClient)
	if err != nil {
		return nil, err
	}
	application, err := message.NewApp(options, loggerLogger, server, db, subscribeClient, middleClient, messageClient, taskClient, workflowClient)
	if err != nil {
		return nil, err
	}
	return application, nil
}

// wire.go:

var providerSet = wire.NewSet(config.ProviderSet, logger.ProviderSet, http.ProviderSet, rpc.ProviderSet, jaeger.ProviderSet, etcd.ProviderSet, influx.ProviderSet, rpcclients.ProviderSet, redis.ProviderSet, message.ProviderSet, database.ProviderSet, rollbar.ProviderSet)
