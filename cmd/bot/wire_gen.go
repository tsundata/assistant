// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/tsundata/assistant/internal/app/bot"
	service2 "github.com/tsundata/assistant/internal/app/bot/finance/service"
	repository2 "github.com/tsundata/assistant/internal/app/bot/org/repository"
	service3 "github.com/tsundata/assistant/internal/app/bot/org/service"
	"github.com/tsundata/assistant/internal/app/bot/todo/repository"
	"github.com/tsundata/assistant/internal/app/bot/todo/service"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/global"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/middleware/etcd"
	"github.com/tsundata/assistant/internal/pkg/middleware/influx"
	"github.com/tsundata/assistant/internal/pkg/middleware/jaeger"
	"github.com/tsundata/assistant/internal/pkg/middleware/mysql"
	"github.com/tsundata/assistant/internal/pkg/middleware/nats"
	"github.com/tsundata/assistant/internal/pkg/middleware/redis"
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
	mysqlConn, err := mysql.New(appConfig)
	if err != nil {
		return nil, err
	}
	todoRepository := repository.NewMysqlTodoRepository(globalID, mysqlConn)
	todo := service.NewTodo(bus, logLogger, todoRepository)
	finance := service2.NewFinance()
	orgRepository := repository2.NewMysqlOrgRepository(globalID, mysqlConn)
	middleSvcClient, err := rpcclient.NewMiddleClient(rpcClient)
	if err != nil {
		return nil, err
	}
	org := service3.NewOrg(orgRepository, middleSvcClient)
	initServer := bot.CreateInitServerFn(todo, finance, org)
	redisClient, err := redis.New(appConfig, newrelicApp)
	if err != nil {
		return nil, err
	}
	server, err := rpc.NewServer(appConfig, logger, logLogger, initServer, tracer, redisClient, newrelicApp)
	if err != nil {
		return nil, err
	}
	chatbotSvcClient, err := rpcclient.NewChatbotClient(rpcClient)
	if err != nil {
		return nil, err
	}
	application, err := bot.NewApp(appConfig, logLogger, server, chatbotSvcClient)
	if err != nil {
		return nil, err
	}
	return application, nil
}

// wire.go:

var providerSet = wire.NewSet(config.ProviderSet, log.ProviderSet, rpc.ProviderSet, jaeger.ProviderSet, influx.ProviderSet, redis.ProviderSet, bot.ProviderSet, rollbar.ProviderSet, etcd.ProviderSet, newrelic.ProviderSet, event.ProviderSet, repository.ProviderSet, repository2.ProviderSet, nats.ProviderSet, global.ProviderSet, mysql.ProviderSet, rpcclient.ProviderSet)
