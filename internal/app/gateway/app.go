package gateway

import (
	"errors"
	"github.com/spf13/viper"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/transports/http"
)

type Options struct {
	Token        string
	Verification string
	Signing      string
}

func NewOptions(v *viper.Viper) (*Options, error) {
	var err error
	o := new(Options)

	if err = v.UnmarshalKey("slack", o); err != nil {
		return nil, errors.New("unmarshal app option error")
	}

	return o, err
}

func NewApp(name string, logger *logger.Logger, hs *http.Server) (*app.Application, error) {
	a, err := app.New(name, logger, app.HTTPServerOption(hs))

	if err != nil {
		return nil, err
	}

	return a, nil
}
