package main

import (
	"flag"
	"github.com/tsundata/assistant/internal/app/middle"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/database"
	"github.com/tsundata/assistant/internal/pkg/etcd"
	"github.com/tsundata/assistant/internal/pkg/influx"
	"github.com/tsundata/assistant/internal/pkg/jaeger"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc"
)

func CreateApp(name, cf string) (*app.Application, error) {
	viper, err := config.New(cf)
	if err != nil {
		return nil, err
	}
	rpcOptions, err := rpc.NewOptions(viper)
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
	in, err := influx.New(influxOptions)
	if err != nil {
		return nil, err
	}

	server, err := rpc.NewServer(rpcOptions, log, j, e, in)
	if err != nil {
		return nil, err
	}

	dbOptions, err := database.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	db, err := database.New(dbOptions)
	if err != nil {
		return nil, err
	}

	appOptions, err := middle.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	application, err := middle.NewApp(name, appOptions, log, server, db, e)
	if err != nil {
		return nil, err
	}
	return application, nil
}

var appName = flag.String("n", "appName", "set app name")
var configFile = flag.String("f", "rpc.yml", "set config file which will loading")

func main() {
	flag.Parse()

	a, err := CreateApp(*appName, *configFile)
	if err != nil {
		panic(err)
	}

	if err := a.Start(); err != nil {
		panic(err)
	}

	a.AwaitSignal()
}
