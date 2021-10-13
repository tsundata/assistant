// +build wireinject

package repository

import (
	"github.com/google/wire"
	"github.com/tsundata/assistant/internal/app/message/rpcclient"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/global"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/middleware/etcd"
	"github.com/tsundata/assistant/internal/pkg/middleware/jaeger"
	"github.com/tsundata/assistant/internal/pkg/middleware/mysql"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
	"github.com/tsundata/assistant/internal/pkg/vendors/newrelic"
	"github.com/tsundata/assistant/internal/pkg/vendors/rollbar"
)

var testProviderSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	etcd.ProviderSet,
	ProviderSet,
	rollbar.ProviderSet,
	mysql.ProviderSet,
	newrelic.ProviderSet,
	rpcclient.ProviderSet,
	rpc.ProviderSet,
	jaeger.ProviderSet,
	global.ProviderSet,
)

func CreateMessageRepository(id string) (MessageRepository, error) {
	panic(wire.Build(testProviderSet))
}
