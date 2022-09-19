//go:build wireinject
// +build wireinject

package rulebot

import (
	"github.com/google/wire"
	financeService "github.com/tsundata/assistant/internal/app/bot/finance/service"
	okrRepository "github.com/tsundata/assistant/internal/app/bot/okr/repository"
	okrService "github.com/tsundata/assistant/internal/app/bot/okr/service"
	todoRepository "github.com/tsundata/assistant/internal/app/bot/todo/repository"
	todoService "github.com/tsundata/assistant/internal/app/bot/todo/service"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/global"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/middleware/etcd"
	"github.com/tsundata/assistant/internal/pkg/middleware/jaeger"
	"github.com/tsundata/assistant/internal/pkg/middleware/mysql"
	"github.com/tsundata/assistant/internal/pkg/middleware/rabbitmq"
	"github.com/tsundata/assistant/internal/pkg/middleware/redis"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc/rpcclient"
	"github.com/tsundata/assistant/internal/pkg/vendors/newrelic"
	"github.com/tsundata/assistant/internal/pkg/vendors/rollbar"
)

var testProviderSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	etcd.ProviderSet,
	ProviderSet,
	rpcclient.ProviderSet,
	redis.ProviderSet,
	rollbar.ProviderSet,
	rpc.ProviderSet,
	jaeger.ProviderSet,
	newrelic.ProviderSet,
	component.ProviderSet,
	event.ProviderSet,
	rabbitmq.ProviderSet,
	global.ProviderSet,
	mysql.ProviderSet,
	okrService.ProviderSet,
	todoRepository.ProviderSet,
	todoService.ProviderSet,
	okrRepository.ProviderSet,
	financeService.ProviderSet,
)

func CreateRuleBot(id string) (*RuleBot, error) {
	panic(wire.Build(testProviderSet))
}
