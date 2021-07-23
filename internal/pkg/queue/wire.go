// +build wireinject

package queue

import (
	"github.com/RichardKnop/machinery/v2"
	"github.com/google/wire"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/middleware/etcd"
	redisMiddle "github.com/tsundata/assistant/internal/pkg/middleware/redis"
)

var testProviderSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	etcd.ProviderSet,
	redisMiddle.ProviderSet,
	ProviderSet,
)

func CreateQueueServer(id string) (*machinery.Server, error) {
	panic(wire.Build(testProviderSet))
}
