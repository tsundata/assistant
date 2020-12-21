package main

import (
	"flag"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/message"
	"github.com/tsundata/assistant/internal/app/message/bot"
	"github.com/tsundata/assistant/internal/app/message/plugins"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/database"
	"github.com/tsundata/assistant/internal/pkg/jaeger"
	"github.com/tsundata/assistant/internal/pkg/logger"
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

	server, err := rpc.NewServer(rpcOptions, log, j, nil)
	if err != nil {
		return nil, err
	}

	clientOptions, err := rpc.NewClientOptions(viper, j)
	if err != nil {
		return nil, err
	}
	subClientConn, err := rpc.NewClient(clientOptions, "subscribe")
	if err != nil {
		return nil, err
	}
	subClient := pb.NewSubscribeClient(subClientConn.CC)
	midClientConn, err := rpc.NewClient(clientOptions, "middle")
	if err != nil {
		return nil, err
	}
	midClient := pb.NewMiddleClient(midClientConn.CC)

	dbOptions, err := database.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	db, err := database.New(dbOptions)
	if err != nil {
		return nil, err
	}
	appOptions, err := message.NewOptions(viper, db, log)
	if err != nil {
		return nil, err
	}
	b := bot.New("ts", viper, &subClient, &midClient, plugins.Options...)
	application, err := message.NewApp(appOptions, server, b)
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
