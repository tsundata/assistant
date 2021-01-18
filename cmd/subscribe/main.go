package main

import (
	"flag"
	"github.com/tsundata/assistant/internal/app/subscribe"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/etcd"
	"github.com/tsundata/assistant/internal/pkg/influx"
	"github.com/tsundata/assistant/internal/pkg/jaeger"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc"
)

func CreateApp(cf string) (*app.Application, error) {
	viper, err := config.New(cf)
	if err != nil {
		return nil, err
	}
	rpcOptions, err := rpc.NewServerOptions(viper)
	if err != nil {
		return nil, err
	}
	log := logger.NewLogger()

	t, err := jaeger.NewConfiguration(viper, log)
	if err != nil {
		return nil, err
	}
	j, err := jaeger.New(t)
	if err != nil {
		return nil, err
	}

	etcdOption, err := etcd.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	e, err := etcd.New(etcdOption)
	if err != nil {
		return nil, err
	}

	influxOptions, err := influx.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	i, err := influx.New(influxOptions)
	if err != nil {
		return nil, err
	}

	server, err := rpc.NewServer(rpcOptions, log, j, e, i)
	if err != nil {
		return nil, err
	}

	appOptions, err := subscribe.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	application, err := subscribe.NewApp(appOptions, log, server, e)
	if err != nil {
		return nil, err
	}
	return application, nil
}

var configFile = flag.String("f", "subscribe.yml", "set config file which will loading")

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
