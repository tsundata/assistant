// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/tsundata/assistant/internal/app/nlp"
	"github.com/tsundata/assistant/internal/app/nlp/service"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/middleware/etcd"
	"github.com/tsundata/assistant/internal/pkg/middleware/influx"
	"github.com/tsundata/assistant/internal/pkg/middleware/jaeger"
	"github.com/tsundata/assistant/internal/pkg/middleware/redis"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
	"github.com/tsundata/assistant/internal/pkg/vendors/newrelic"
	"github.com/tsundata/assistant/internal/pkg/vendors/rollbar"
)

// Injectors from wire.go:

func CreateApp(id string) (*app.Application, error) {
	appConfig := config.NewConfig(id)
	rollbarRollbar := rollbar.New(appConfig)
	logger := log.NewZapLogger(rollbarRollbar)
	logLogger := log.NewAppLogger(logger)
	serviceNLP := service.NewNLP(appConfig, logLogger)
	initServer := service.CreateInitServerFn(serviceNLP)
	server, err := rpc.NewServer(appConfig, initServer)
	if err != nil {
		return nil, err
	}
	application, err := nlp.NewApp(appConfig, logLogger, server)
	if err != nil {
		return nil, err
	}
	return application, nil
}

// wire.go:

var providerSet = wire.NewSet(config.ProviderSet, log.ProviderSet, rpc.ProviderSet, jaeger.ProviderSet, influx.ProviderSet, redis.ProviderSet, nlp.ProviderSet, rollbar.ProviderSet, etcd.ProviderSet, service.ProviderSet, newrelic.ProviderSet)
