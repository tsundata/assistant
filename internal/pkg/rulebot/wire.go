// +build wireinject

package rulebot

import (
	"github.com/google/wire"
	"github.com/tsundata/assistant/internal/app/cron/rpcclient"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/middleware/consul"
	"github.com/tsundata/assistant/internal/pkg/middleware/jaeger"
	"github.com/tsundata/assistant/internal/pkg/middleware/redis"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
	"github.com/tsundata/assistant/internal/pkg/vendors/rollbar"
)

var testProviderSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	consul.ProviderSet,
	ProviderSet,
	rpcclient.ProviderSet,
	redis.ProviderSet,
	rollbar.ProviderSet,
	rpc.ProviderSet,
	jaeger.ProviderSet,
)

func CreateRuleBot(id string) (*RuleBot, error) {
	panic(wire.Build(testProviderSet))
}
