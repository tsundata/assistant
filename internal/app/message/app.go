package message

import (
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc"
)

type Options struct {
	Name string
}

func NewOptions() (*Options, error) {
	var err error
	o := new(Options)

	return o, err
}

func NewApp(o *Options, rs *rpc.Server) (*app.Application, error) {
	a, err := app.New(o.Name, app.RpcServerOption(rs))

	if err != nil {
		return nil, err
	}

	return a, nil
}
