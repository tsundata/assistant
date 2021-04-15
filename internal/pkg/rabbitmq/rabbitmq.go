package rabbitmq

import (
	"errors"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

type Options struct {
	URL string `yaml:"url"`
}

func NewOptions(v *viper.Viper) (*Options, error) {
	var err error
	o := new(Options)
	if err = v.UnmarshalKey("rabbitmq", o); err != nil {
		return nil, errors.New("unmarshal etcd option error")
	}

	return o, err
}

func New(o *Options) (*amqp.Connection, error) {
	return amqp.Dial(o.URL)
}

var ProviderSet = wire.NewSet(New, NewOptions)
