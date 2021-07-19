package main

import (
	"github.com/tsundata/assistant/internal/app/agent"
	"github.com/tsundata/assistant/internal/app/agent/broker"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/middleware/consul"
	"github.com/tsundata/assistant/internal/pkg/middleware/influx"
	"github.com/tsundata/assistant/internal/pkg/vendors/rollbar"
)

func CreateApp() (*app.Application, error) {
	c, err := consul.New()
	if err != nil {
		return nil, err
	}
	appConfig := config.NewConfig(app.ServerAgent, c)

	r := rollbar.New(appConfig)

	zap := log.NewZapLogger(r)
	l := log.NewAppLogger(zap)

	i, err := influx.New(appConfig)
	if err != nil {
		return nil, err
	}

	br, err := broker.NewBroker(appConfig, l, i)
	if err != nil {
		return nil, err
	}

	b, err := broker.NewServer(br)
	if err != nil {
		return nil, err
	}

	application, err := agent.NewApp(appConfig, l, b)
	if err != nil {
		return nil, err
	}
	return application, nil
}

func main() {
	a, err := CreateApp()
	if err != nil {
		panic(err)
	}

	if err := a.Start(); err != nil {
		panic(err)
	}

	a.AwaitSignal()
}
