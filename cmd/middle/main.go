package main

import (
	"flag"
	"github.com/tsundata/assistant/internal/app/middle"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/database"
	"github.com/tsundata/assistant/internal/pkg/jaeger"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/redis"
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

	server, err := rpc.NewServer(rpcOptions, log, j)
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

	redisOption, err := redis.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	r, err := redis.New(redisOption)
	if err != nil {
		return nil, err
	}

	appOptions, err := middle.NewOptions(viper, db, log, r)
	if err != nil {
		return nil, err
	}
	application, err := middle.NewApp(appOptions, server)
	if err != nil {
		return nil, err
	}
	return application, nil
}

var configFile = flag.String("f", "middle.yml", "set config file which will loading")

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
