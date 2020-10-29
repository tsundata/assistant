package main

import (
	"flag"

	"github.com/tsundata/assistant/internal/app/message"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc"
)

func CreateApp(cf string) (*app.Application, error) {
	rpcOptions, err := rpc.NewServerOptions()
	// FIXME
	rpcOptions.RegistryAddr = "http://127.0.0.1:7001/_rpc_/registry"
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
