package bot

import (
	"github.com/google/wire"
	service3 "github.com/tsundata/assistant/internal/app/bot/finance/service"
	service2 "github.com/tsundata/assistant/internal/app/bot/org/service"
	"github.com/tsundata/assistant/internal/app/bot/todo/service"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
)

func NewApp(c *config.AppConfig, logger log.Logger, rs *rpc.Server) (*app.Application, error) {
	a, err := app.New(c, logger, app.RPCServerOption(rs))
	if err != nil {
		return nil, err
	}

	return a, nil
}

var ProviderSet = wire.NewSet(NewApp, CreateInitServerFn,
	service.NewTodo, service2.NewOrg, service3.NewFinance)
