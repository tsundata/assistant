// +build wireinject

package collection

import (
	"github.com/go-redis/redis/v8"
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
)

func CreateRedisClient(id string) (*redis.Client, error) {
	panic(wire.Build(testProviderSet))
}
