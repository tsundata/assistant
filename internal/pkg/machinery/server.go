package machinery

import (
	"errors"
	"github.com/RichardKnop/machinery/v2"
	amqpBackend "github.com/RichardKnop/machinery/v2/backends/amqp"
	amqpBroker "github.com/RichardKnop/machinery/v2/brokers/amqp"
	"github.com/RichardKnop/machinery/v2/config"
	eagerlock "github.com/RichardKnop/machinery/v2/locks/eager"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

const DefaultQueue = "assistant_tasks"
const AMQPExchange = "machinery_exchange"
const AMQPBindingKey = "machinery_task"

type Options struct {
	URL string `yaml:"url"`
}

func NewOptions(v *viper.Viper) (*Options, error) {
	var err error
	o := new(Options)
	if err = v.UnmarshalKey("rabbitmq", o); err != nil {
		return nil, errors.New("unmarshal machinery redis option error")
	}

	return o, err
}

func New(o *Options) (*machinery.Server, error) {
	cnf := &config.Config{
		DefaultQueue:    DefaultQueue,
		ResultsExpireIn: 3600,
		Broker:          o.URL,
		ResultBackend:   o.URL,
		AMQP: &config.AMQPConfig{
			Exchange:     AMQPExchange,
			ExchangeType: "direct",
			BindingKey:   AMQPBindingKey,
		},
	}

	server := machinery.NewServer(cnf, amqpBroker.New(cnf), amqpBackend.New(cnf), eagerlock.New())
	return server, nil
}

var ProviderSet = wire.NewSet(New, NewOptions)
