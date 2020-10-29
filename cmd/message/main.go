package main

import (
	"flag"

	"github.com/tsundata/assistant/internal/app/message"
	"github.com/tsundata/assistant/internal/config"
	"github.com/tsundata/assistant/internal/pkg/app"
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

	messageOptions, err := message.NewOptions()
	if err != nil {
		return nil, err
	}
	application, err := message.NewApp(messageOptions, server)
	if err != nil {
		return nil, err
	}
	return application, nil
}

var configFile = flag.String("f", "message.yml", "set config file which will loading")

func main() {
	flag.Parse()

	app, err := CreateApp(*configFile)
	if err != nil {
		panic(err)
	}

	if err := app.Start(); err != nil {
		panic(err)
	}

	app.AwaitSignal()
}
