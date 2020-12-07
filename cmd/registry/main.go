package main

import (
	"flag"
	"github.com/tsundata/assistant/internal/app/registry"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc"
)

func CreateApp(cf string) (*app.Application, error) {
	viper, err := config.New(cf)
	if err != nil {
		return nil, err
	}
	rpcOptions, err := rpc.NewRegistryOptions(viper)
	if err != nil {
		return nil, err
	}
	log := logger.NewLogger()
	server, err := rpc.NewRegistry(rpcOptions, log, nil)

	registryOptions, err := registry.NewOptions(viper, log)
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

	a, err := CreateApp(*configFile)
	if err != nil {
		panic(err)
	}

	if err := a.Start(); err != nil {
		panic(err)
	}

	a.AwaitSignal()
}
