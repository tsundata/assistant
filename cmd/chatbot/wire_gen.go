// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/tsundata/assistant/internal/app/chatbot"
	"github.com/tsundata/assistant/internal/app/chatbot/rpcclient"
	"github.com/tsundata/assistant/internal/app/chatbot/service"
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
	"github.com/tsundata/assistant/internal/pkg/rulebot"
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
	configuration, err := jaeger.NewConfiguration(appConfig, loggerLogger)
	if err != nil {
		return nil, err
	}
	tracer, err := jaeger.New(configuration)
	if err != nil {
		return nil, err
	}
	clientOptions, err := rpc.NewClientOptions(tracer)
	if err != nil {
		return nil, err
	}
	rpcClient, err := rpc.NewClient(clientOptions, client, loggerLogger)
	if err != nil {
		return nil, err
	}
	middleClient, err := rpcclient.NewMiddleClient(rpcClient)
	if err != nil {
		return nil, err
	}
	todoClient, err := rpcclient.NewTodoClient(rpcClient)
	if err != nil {
		return nil, err
	}
	redisClient, err := redis.New(appConfig)
	if err != nil {
		return nil, err
	}
	messageClient, err := rpcclient.NewMessageClient(rpcClient)
	if err != nil {
		return nil, err
	}
	subscribeClient, err := rpcclient.NewSubscribe(rpcClient)
	if err != nil {
		return nil, err
	}
	workflowClient, err := rpcclient.NewWorkflowClient(rpcClient)
	if err != nil {
		return nil, err
	}
	storageClient, err := rpcclient.NewStorageClient(rpcClient)
	if err != nil {
		return nil, err
	}
	userClient, err := rpcclient.NewUserClient(rpcClient)
	if err != nil {
		return nil, err
	}
	nlpClient, err := rpcclient.NewNLPClient(rpcClient)
	if err != nil {
		return nil, err
	}
	iContext := rulebot.NewContext(appConfig, redisClient, loggerLogger, messageClient, middleClient, subscribeClient, workflowClient, storageClient, todoClient, userClient, nlpClient)
	ruleBot := rulebot.New(iContext)
	serviceChatbot := service.NewChatbot(loggerLogger, middleClient, todoClient, ruleBot)
	initServer := service.CreateInitServerFn(serviceChatbot)
	influxdb2Client, err := influx.New(appConfig)
	if err != nil {
		return nil, err
	}
	server, err := rpc.NewServer(appConfig, loggerLogger, initServer, tracer, influxdb2Client, redisClient, client)
	if err != nil {
		return nil, err
	}
	application, err := chatbot.NewApp(appConfig, bus, loggerLogger, server, middleClient, todoClient, userClient)
	if err != nil {
		return nil, err
	}
	return application, nil
}

// wire.go:

var providerSet = wire.NewSet(config.ProviderSet, logger.ProviderSet, rpc.ProviderSet, jaeger.ProviderSet, influx.ProviderSet, redis.ProviderSet, chatbot.ProviderSet, mysql.ProviderSet, rollbar.ProviderSet, consul.ProviderSet, service.ProviderSet, rpcclient.ProviderSet, rulebot.ProviderSet, event.ProviderSet, nats.ProviderSet)
