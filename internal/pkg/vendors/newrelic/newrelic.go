package newrelic

import (
	"context"
	"github.com/google/wire"
	"github.com/newrelic/go-agent/v3/integrations/nrzap"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/tsundata/assistant/internal/pkg/config"
	"go.uber.org/zap"
)

type App struct {
	nr *newrelic.Application
}

func New(c *config.AppConfig, zap *zap.Logger) (*App, error) {
	var nr *newrelic.Application
	var err error
	if c.Newrelic.Name == "" || c.Newrelic.License == "" {
		nr, _ = newrelic.NewApplication()
	} else {
		nr, err = newrelic.NewApplication(
			newrelic.ConfigAppName(c.Newrelic.Name),
			newrelic.ConfigLicense(c.Newrelic.License),
			newrelic.ConfigDistributedTracerEnabled(true),
			nrzap.ConfigLogger(zap.Named("newrelic")),
		)
		if err != nil {
			return nil, err
		}
	}

	return &App{nr: nr}, nil
}

func (a *App) Application() *newrelic.Application {
	return a.nr
}

func (a *App) StartTransaction(name string) *newrelic.Transaction {
	return a.nr.StartTransaction(name)
}

func (a *App) NewContext(ctx context.Context, nxt *newrelic.Transaction) context.Context {
	return newrelic.NewContext(ctx, nxt)
}

var ProviderSet = wire.NewSet(New)
