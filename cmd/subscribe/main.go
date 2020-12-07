package main

import (
	"flag"
	"github.com/tsundata/assistant/internal/pkg/database"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/redis"

	"github.com/tsundata/assistant/internal/app/subscribe"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
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
	server, err := rpc.NewServer(rpcOptions, log, nil)

	dbOptions, err := database.NewOptions(viper)
	db, err := database.New(dbOptions)

	redisOption, err := redis.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	r, err := redis.New(redisOption)
	if err != nil {
		return nil, err
	}

	appOptions, err := subscribe.NewOptions(viper, db, log, r)
	if err != nil {
		return nil, err
	}
	application, err := subscribe.NewApp(appOptions, server)
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
