package web

import (
	"github.com/google/wire"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/transports/http"
)

func NewApp(c *config.AppConfig, logger *logger.Logger, hs *http.Server) (*app.Application, error) {
	a, err := app.New(c, logger, app.HTTPServerOption(hs))

	if err != nil {
		return nil, err
	}

	return a, nil
}

var ProviderSet = wire.NewSet(NewApp)
