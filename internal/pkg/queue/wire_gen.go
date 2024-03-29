// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package queue

import (
	"github.com/RichardKnop/machinery/v2"
	"github.com/google/wire"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/middleware/etcd"
	"github.com/tsundata/assistant/internal/pkg/middleware/redis"
)

// Injectors from wire.go:

func CreateQueueServer(id string) (*machinery.Server, error) {
	client, err := etcd.New()
	if err != nil {
		return nil, err
	}
	appConfig := config.NewConfig(id, client)
	server, err := New(appConfig)
	if err != nil {
		return nil, err
	}
	return server, nil
}

// wire.go:

var testProviderSet = wire.NewSet(log.ProviderSet, config.ProviderSet, etcd.ProviderSet, redis.ProviderSet, ProviderSet)
