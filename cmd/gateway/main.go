package main

import (
	"flag"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/gateway"
	"github.com/tsundata/assistant/internal/app/gateway/controllers"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/redis"
	"github.com/tsundata/assistant/internal/pkg/transports/http"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc"
)

func CreateApp(cf string) (*app.Application, error) {
	viper, err := config.New(cf)
	if err != nil {
		return nil, err
	}
	httpOptions, err := http.NewOptions(viper)
	if err != nil {
		return nil, err
	}

	clientOptions, err := rpc.NewClientOptions(viper)
	if err != nil {
		return nil, err
	}
	subClientConn, err := rpc.NewClient(clientOptions, "subscribe")
	if err != nil {
		return nil, err
	}
	subClient := pb.NewSubscribeClient(subClientConn.CC)
	msgClientConn, err := rpc.NewClient(clientOptions, "message")
	if err != nil {
		return nil, err
	}
	msgClient := pb.NewMessageClient(msgClientConn.CC)

	redisOption, err := redis.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	r, err := redis.New(redisOption)
	if err != nil {
		return nil, err
	}

	log := logger.NewLogger()
	gatewayOptions, err := gateway.NewOptions(viper, log)
	if err != nil {
		return nil, err
	}
	gatewayController := controllers.NewGatewayController(gatewayOptions, r, log, &subClient, &msgClient)
	initControllers := controllers.CreateInitControllersFn(gatewayController)
	router := http.NewRouter(httpOptions, initControllers)
	server, err := http.New(httpOptions, router)
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

	a, err := CreateApp(*configFile)
	if err != nil {
		panic(err)
	}

	if err := a.Start(); err != nil {
		panic(err)
	}

	a.AwaitSignal()
}
