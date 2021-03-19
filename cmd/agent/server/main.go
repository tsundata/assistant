package main

import (
	"flag"
	"github.com/tsundata/assistant/internal/app/agent"
	"github.com/tsundata/assistant/internal/app/agent/broker"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/influx"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/vendors/rollbar"
)

func CreateApp(cf string) (*app.Application, error) {
	viper, err := config.New(cf)
	if err != nil {
		return nil, err
	}

	rollbarOptions, err := rollbar.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	r := rollbar.New(rollbarOptions)

	log := logger.NewLogger(r)

	influxOptions, err := influx.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	i, err := influx.New(influxOptions)
	if err != nil {
		return nil, err
	}

	br, err := broker.NewBroker(viper, log, i)
	if err != nil {
		return nil, err
	}

	b, err := broker.NewServer(br)
	if err != nil {
		return nil, err
	}

	application, err := agent.NewApp(log, b)
	if err != nil {
		return nil, err
	}
	return application, nil
}

var configFile = flag.String("f", "agent.yml", "set config file which will loading")

func main() {
	flag.Parse()

	a, err := CreateApp(*configFile)
	if err != nil {
		panic(err)
	}

	if err := a.Start(); err != nil {
		panic(err)
	}

	a.AwaitSignal()
}
