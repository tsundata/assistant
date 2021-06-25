// +build wireinject

package queue

import (
	"github.com/RichardKnop/machinery/v2"
	"github.com/google/wire"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/middleware/consul"
	redisMiddle "github.com/tsundata/assistant/internal/pkg/middleware/redis"
)

var testProviderSet = wire.NewSet(
	logger.ProviderSet,
	config.ProviderSet,
	consul.ProviderSet,
	redisMiddle.ProviderSet,
	ProviderSet,
)

func CreateQueueServer(id string) (*machinery.Server, error) {
	panic(wire.Build(testProviderSet))
}
