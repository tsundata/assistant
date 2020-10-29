package main

import (
	"flag"

	"github.com/tsundata/assistant/internal/app/registry"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc"
)

func CreateApp(cf string) (*app.Application, error) {
	rpcOptions, err := rpc.NewRegistryOptions()
	// FIXME
	rpcOptions.Port = 7001
	if err != nil {
		return nil, err
	}

	server, err := rpc.NewRegistry(rpcOptions, nil)

	registryOptions, err := registry.NewOptions()
	if err != nil {
		return nil, err
	}
	application, err := registry.NewApp(registryOptions, server)
	if err != nil {
		return nil, err
	}
	return application, nil
}

var configFile = flag.String("f", "registry.yml", "set config file which will loading")

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
