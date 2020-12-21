package main

import (
	"flag"
	"github.com/tsundata/assistant/internal/app/web"
	"github.com/tsundata/assistant/internal/app/web/controllers"
	"github.com/tsundata/assistant/internal/app/web/rpcclients"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/jaeger"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/transports/http"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc"
)

func CreateApp(cf string) (*app.Application, error) {
	log := logger.NewLogger()
	viper, err := config.New(cf)
	if err != nil {
		return nil, err
	}
	httpOptions, err := http.NewOptions(viper)
	if err != nil {
		return nil, err
	}

	t, err := jaeger.NewConfiguration(viper, log)
	if err != nil {
		return nil, err
	}
	j, err := jaeger.New(t)
	if err != nil {
		return nil, err
	}

	clientOptions, err := rpc.NewClientOptions(viper, j)
	if err != nil {
		return nil, err
	}
	client, err := rpc.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}
	midClient, err := rpcclients.NewMiddleClient(client)
	if err != nil {
		return nil, err
	}

	webOptions, err := web.NewOptions(viper, log)
	if err != nil {
		return nil, err
	}
	webController := controllers.NewWebController(webOptions, log, &midClient)
	initControllers := controllers.CreateInitControllersFn(webController)
	router := http.NewRouter(httpOptions, initControllers)
	server, err := http.New(httpOptions, router)
	if err != nil {
		return nil, err
	}
	application, err := web.NewApp(webOptions, server)
	if err != nil {
		return nil, err
	}
	return application, nil
}

var configFile = flag.String("f", "web.yml", "set config file which will loading")

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
