// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package repository

import (
	"github.com/google/wire"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/global"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/middleware/etcd"
	"github.com/tsundata/assistant/internal/pkg/middleware/jaeger"
	"github.com/tsundata/assistant/internal/pkg/middleware/mysql"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc/rpcclient"
	"github.com/tsundata/assistant/internal/pkg/vendors/newrelic"
	"github.com/tsundata/assistant/internal/pkg/vendors/rollbar"
)

// Injectors from wire.go:

func CreateTodoRepository(id string) (TodoRepository, error) {
	client, err := etcd.New()
	if err != nil {
		return nil, err
	}
	appConfig := config.NewConfig(id, client)
	rollbarRollbar := rollbar.New(appConfig)
	logger := log.NewZapLogger(rollbarRollbar)
	logLogger := log.NewAppLogger(logger)
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
	conn, err := mysql.New(appConfig)
	if err != nil {
		return nil, err
	}
	todoRepository := NewMysqlTodoRepository(globalID, conn)
	return todoRepository, nil
}

// wire.go:

var testProviderSet = wire.NewSet(log.ProviderSet, config.ProviderSet, etcd.ProviderSet, ProviderSet, rollbar.ProviderSet, mysql.ProviderSet, newrelic.ProviderSet, global.ProviderSet, rpc.ProviderSet, jaeger.ProviderSet, rpcclient.ProviderSet)
