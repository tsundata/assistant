package main

import (
	"flag"

	"github.com/tsundata/assistant/internal/app/gateway"
	"github.com/tsundata/assistant/internal/app/gateway/controllers"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/transports/http"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc"
)

func CreateApp(cf string) (*app.Application, error) {
	httpOptions, err := http.NewOptions()
	// FIXME
	httpOptions.Port = 7000
	httpOptions.Host = "0.0.0.0"
	if err != nil {
		return nil, err
	}
	clientOptions, err := rpc.NewClientOptions()
	if err != nil {
		return nil, err
	}
	// FIXME
	clientOptions.Registry = "http://127.0.0.1:7001/_rpc_/registry"
	client, err := rpc.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}
	gatewayController := controllers.NewGatewayController(client)
	initControllers := controllers.CreateInitControllersFn(gatewayController)
	engine := http.NewRouter(httpOptions, initControllers)
	server, err := http.New(httpOptions, engine)
	if err != nil {
		return nil, err
	}
	gatewayOptions, err := gateway.NewOptions()
	if err != nil {
		return nil, err
	}
	application, err := gateway.NewApp(gatewayOptions, server)
	if err != nil {
		return nil, err
	}
	return application, nil
}

var configFile = flag.String("f", "gateway.yml", "set config file which will loading")

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
