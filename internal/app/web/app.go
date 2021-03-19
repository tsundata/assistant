package web

import (
	"errors"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/transports/http"
	"os"
)

type Options struct {
	Name string
	URL  string
}

func NewOptions(v *viper.Viper) (*Options, error) {
	var err error
	o := new(Options)

	o.Name = os.Getenv("APP_NAME")

	if err = v.UnmarshalKey("web", o); err != nil {
		return nil, errors.New("unmarshal web option error")
	}

	return o, err
}

func NewApp(o *Options, logger *logger.Logger, hs *http.Server) (*app.Application, error) {
	a, err := app.New(o.Name, logger, app.HTTPServerOption(hs))

	if err != nil {
		return nil, err
	}

	return a, nil
}

var ProviderSet = wire.NewSet(NewApp, NewOptions)
