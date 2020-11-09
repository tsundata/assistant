package message

import (
	"errors"
	"github.com/spf13/viper"
	"github.com/tsundata/assistant/internal/app/message/service"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc"
	"log"
)

type Options struct {
	Name    string
	Webhook string
}

func NewOptions(v *viper.Viper) (*Options, error) {
	var err error
	o := new(Options)

	if err = v.UnmarshalKey("app", o); err != nil {
		return nil, errors.New("unmarshal app option error")
	}

	if err = v.UnmarshalKey("slack", o); err != nil {
		return nil, errors.New("unmarshal app option error")
	}

	log.Println("load application options success")

	return o, err
}

func NewApp(o *Options, rs *rpc.Server) (*app.Application, error) {
	slack := service.NewSlack(o.Webhook)
	err := rs.Register(slack, "")
	if err != nil {
		return nil, err
	}

	a, err := app.New(o.Name, app.RpcServerOption(rs))
	if err != nil {
		return nil, err
	}

	return a, nil
}
