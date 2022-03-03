// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/tsundata/assistant/internal/app/chatbot"
	"github.com/tsundata/assistant/internal/app/chatbot/repository"
	"github.com/tsundata/assistant/internal/app/chatbot/service"
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
	"github.com/tsundata/assistant/internal/pkg/robot/rulebot"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc/rpcclient"
	"github.com/tsundata/assistant/internal/pkg/vendors/newrelic"
	"github.com/tsundata/assistant/internal/pkg/vendors/rollbar"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

func CreateApp(id string) (*app.Application, error) {
	client, err := etcd.New()
	if err != nil {
		return nil, err
	}
	appConfig := config.NewConfig(id, client)
	connection, err := rabbitmq.New(appConfig)
	if err != nil {
		return nil, err
	}
	rollbarRollbar := rollbar.New(appConfig)
	logger := log.NewZapLogger(rollbarRollbar)
	logLogger := log.NewAppLogger(logger)
	bus := event.NewNatsBus(connection, logLogger)
	configuration, err := jaeger.NewConfiguration(appConfig, logLogger)
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
	rpcClient, err := rpc.NewClient(clientOptions, appConfig, logLogger)
	if err != nil {
		return nil, err
	}
	idSvcClient, err := rpcclient.NewIdClient(rpcClient)
	if err != nil {
		return nil, err
	}
	globalID := global.NewID(appConfig, idSvcClient)
	locker := global.NewLocker(client)
	conn, err := mysql.New(appConfig)
	if err != nil {
		return nil, err
	}
	chatbotRepository := repository.NewMysqlChatbotRepository(globalID, locker, conn)
	messageSvcClient, err := rpcclient.NewMessageClient(rpcClient)
	if err != nil {
		return nil, err
	}
	newrelicApp, err := newrelic.New(appConfig, logger)
	if err != nil {
		return nil, err
	}
	redisClient, err := redis.New(appConfig, newrelicApp)
	if err != nil {
		return nil, err
	}
	middleSvcClient, err := rpcclient.NewMiddleClient(rpcClient)
	if err != nil {
		return nil, err
	}
	workflowSvcClient, err := rpcclient.NewWorkflowClient(rpcClient)
	if err != nil {
		return nil, err
	}
	storageSvcClient, err := rpcclient.NewStorageClient(rpcClient)
	if err != nil {
		return nil, err
	}
	todoSvcClient, err := rpcclient.NewTodoClient(rpcClient)
	if err != nil {
		return nil, err
	}
	userSvcClient, err := rpcclient.NewUserClient(rpcClient)
	if err != nil {
		return nil, err
	}
	nlpSvcClient, err := rpcclient.NewNLPClient(rpcClient)
	if err != nil {
		return nil, err
	}
	orgSvcClient, err := rpcclient.NewOrgClient(rpcClient)
	if err != nil {
		return nil, err
	}
	financeSvcClient, err := rpcclient.NewFinanceClient(rpcClient)
	if err != nil {
		return nil, err
	}
	iComponent := rulebot.NewComponent(appConfig, redisClient, logLogger, messageSvcClient, middleSvcClient, workflowSvcClient, storageSvcClient, todoSvcClient, userSvcClient, nlpSvcClient, orgSvcClient, financeSvcClient)
	ruleBot := rulebot.New(iComponent)
	serviceChatbot := service.NewChatbot(logLogger, bus, chatbotRepository, messageSvcClient, ruleBot)
	initServer := service.CreateInitServerFn(serviceChatbot)
	server, err := rpc.NewServer(appConfig, logger, logLogger, initServer, tracer, redisClient, newrelicApp)
	if err != nil {
		return nil, err
	}
	application, err := chatbot.NewApp(appConfig, bus, logLogger, server, messageSvcClient, chatbotRepository, ruleBot)
	if err != nil {
		return nil, err
	}
	return application, nil
}

// wire.go:

var providerSet = wire.NewSet(config.ProviderSet, log.ProviderSet, rpc.ProviderSet, jaeger.ProviderSet, influx.ProviderSet, redis.ProviderSet, chatbot.ProviderSet, rollbar.ProviderSet, etcd.ProviderSet, service.ProviderSet, rpcclient.ProviderSet, rulebot.ProviderSet, event.ProviderSet, newrelic.ProviderSet, mysql.ProviderSet, repository.ProviderSet, global.ProviderSet, rabbitmq.ProviderSet)
