package main

import (
	"flag"
	"github.com/tsundata/assistant/api/pb"
	"github.com/tsundata/assistant/internal/app/subscribe"
	"github.com/tsundata/assistant/internal/app/subscribe/spider"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/database"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/redis"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc"
	"time"
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
	if err != nil {
		return nil, err
	}

	dbOptions, err := database.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	db, err := database.New(dbOptions)
	if err != nil {
		return nil, err
	}

	redisOption, err := redis.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	rdb, err := redis.New(redisOption)
	if err != nil {
		return nil, err
	}

	clientOptions, err := rpc.NewClientOptions(viper, rpc.WithTimeout(30*time.Second))
	if err != nil {
		return nil, err
	}
	msgClientConn, err := rpc.NewClient(clientOptions, "message")
	if err != nil {
		return nil, err
	}
	msgClient := pb.NewMessageClient(msgClientConn.CC)
	midClientConn, err := rpc.NewClient(clientOptions, "middle")
	if err != nil {
		return nil, err
	}
	midClient := pb.NewMiddleClient(midClientConn.CC)

	s := spider.New(rdb, &msgClient, &midClient)
	appOptions, err := subscribe.NewOptions(viper, db, log)
	if err != nil {
		return nil, err
	}
	application, err := subscribe.NewApp(appOptions, s, server)
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
