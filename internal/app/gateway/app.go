package gateway

import (
	"errors"
	"github.com/spf13/viper"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/transports/http"
	"go.uber.org/zap"
)

type Options struct {
	Name         string
	Token        string
	Verification string
	Signing      string
	logger       *zap.Logger
}

func NewOptions(v *viper.Viper, logger *zap.Logger) (*Options, error) {
	var err error
	o := new(Options)
	o.logger = logger


	if err = v.UnmarshalKey("app", o); err != nil {
		return nil, errors.New("unmarshal app option error")
	}

	if err = v.UnmarshalKey("slack", o); err != nil {
		return nil, errors.New("unmarshal app option error")
	}

	return o, err
}

func NewApp(o *Options, hs *http.Server) (*app.Application, error) {
	a, err := app.New(o.Name, o.logger, app.HTTPServerOption(hs))

	if err != nil {
		return nil, err
	}

	return a, nil
}
