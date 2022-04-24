//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/tsundata/assistant/internal/app/chatbot"
	financeService "github.com/tsundata/assistant/internal/app/chatbot/bot/finance/service"
	orgRepository "github.com/tsundata/assistant/internal/app/chatbot/bot/org/repository"
	orgService "github.com/tsundata/assistant/internal/app/chatbot/bot/org/service"
	systemRepository "github.com/tsundata/assistant/internal/app/chatbot/bot/system/repository"
	systemService "github.com/tsundata/assistant/internal/app/chatbot/bot/system/service"
	todoRepository "github.com/tsundata/assistant/internal/app/chatbot/bot/todo/repository"
	todoService "github.com/tsundata/assistant/internal/app/chatbot/bot/todo/service"
	"github.com/tsundata/assistant/internal/app/chatbot/repository"
	"github.com/tsundata/assistant/internal/app/chatbot/service"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/global"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/middleware/etcd"
	"github.com/tsundata/assistant/internal/pkg/middleware/influx"
	"github.com/tsundata/assistant/internal/pkg/middleware/jaeger"
	"github.com/tsundata/assistant/internal/pkg/middleware/mysql"
	"github.com/tsundata/assistant/internal/pkg/middleware/rabbitmq"
	"github.com/tsundata/assistant/internal/pkg/middleware/redis"
	"github.com/tsundata/assistant/internal/pkg/robot/component"
	"github.com/tsundata/assistant/internal/pkg/robot/rulebot"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc/rpcclient"
	"github.com/tsundata/assistant/internal/pkg/vendors/newrelic"
	"github.com/tsundata/assistant/internal/pkg/vendors/rollbar"
)

var providerSet = wire.NewSet(
	config.ProviderSet,
	log.ProviderSet,
	rpc.ProviderSet,
	jaeger.ProviderSet,
	influx.ProviderSet,
	redis.ProviderSet,
	chatbot.ProviderSet,
	rollbar.ProviderSet,
	etcd.ProviderSet,
	service.ProviderSet,
	rpcclient.ProviderSet,
	rulebot.ProviderSet,
	event.ProviderSet,
	newrelic.ProviderSet,
	mysql.ProviderSet,
	repository.ProviderSet,
	global.ProviderSet,
	rabbitmq.ProviderSet,
	component.ProviderSet,
	orgService.ProviderSet,
	todoRepository.ProviderSet,
	todoService.ProviderSet,
	systemRepository.ProviderSet,
	systemService.ProviderSet,
	orgRepository.ProviderSet,
	financeService.ProviderSet,
)

func CreateApp(id string) (*app.Application, error) {
	panic(wire.Build(providerSet))
}
