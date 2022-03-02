package bot

import (
	"context"
	"github.com/google/wire"
	"github.com/tsundata/assistant/internal/app/bot/finance"
	service3 "github.com/tsundata/assistant/internal/app/bot/finance/service"
	"github.com/tsundata/assistant/internal/app/bot/org"
	service2 "github.com/tsundata/assistant/internal/app/bot/org/service"
	_ "github.com/tsundata/assistant/internal/app/bot/plugin"
	"github.com/tsundata/assistant/internal/app/bot/todo"
	"github.com/tsundata/assistant/internal/app/bot/todo/service"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/event"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/robot/bot"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
)

func NewApp(c *config.AppConfig, logger log.Logger, rs *rpc.Server, bus event.Bus) (*app.Application, error) {
	a, err := app.New(c, logger, app.RPCServerOption(rs))
	if err != nil {
		return nil, err
	}

	// bots register
	err = bot.RegisterBot(context.Background(), bus, todo.Bot, org.Bot, finance.Bot)
	if err != nil {
		return nil, err
	}

	return a, nil
}

var ProviderSet = wire.NewSet(NewApp, CreateInitServerFn,
	service.NewTodo, service2.NewOrg, service3.NewFinance)
