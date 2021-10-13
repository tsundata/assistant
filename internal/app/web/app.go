package web

import (
	"github.com/google/wire"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/transport/http"
)

func NewApp(c *config.AppConfig, logger log.Logger, hs *http.Server) (*app.Application, error) {
	a, err := app.New(c, app.HTTPServerOption(hs))

	if err != nil {
		return nil, err
	}

	return a, nil
}

var ProviderSet = wire.NewSet(NewApp)
