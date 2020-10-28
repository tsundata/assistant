package gateway

import (
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/transports/http"
)

type Options struct {
	Name string
}

func NewOptions() (*Options, error) {
	var err error
	o := new(Options)

	return o, err
}

func NewApp(o *Options, hs *http.Server) (*app.Application, error) {
	a, err := app.New(o.Name, app.HttpServerOption(hs))

	if err != nil {
		return nil, err
	}

	return a, nil
}
