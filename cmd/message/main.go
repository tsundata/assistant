package main

import (
	"flag"
	"github.com/tsundata/assistant/internal/pkg/database"

	"github.com/tsundata/assistant/internal/app/message"
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

	server, err := rpc.NewServer(rpcOptions, nil)

	dbOptions, err := database.NewOptions(viper)
	db, err := database.New(dbOptions)

	appOptions, err := message.NewOptions(viper, db)
	if err != nil {
		return nil, err
	}
	application, err := message.NewApp(appOptions, server)
	if err != nil {
		return nil, err
	}
	return application, nil
}

var configFile = flag.String("f", "message.yml", "set config file which will loading")

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
