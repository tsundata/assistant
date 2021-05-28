package rollbar

import (
	"github.com/google/wire"
	"github.com/rollbar/rollbar-go"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/version"
)

type Rollbar struct {
	c *config.AppConfig
}

func New(c *config.AppConfig) *Rollbar {
	return &Rollbar{c: c}
}

func (r *Rollbar) Config() {
	rollbar.SetToken(r.c.Rollbar.Token)
	rollbar.SetEnvironment(r.c.Rollbar.Environment)
	rollbar.SetCodeVersion(version.Version)
	rollbar.SetServerRoot("github.com/tsundata/assistant")
}

var ProviderSet = wire.NewSet(New)
