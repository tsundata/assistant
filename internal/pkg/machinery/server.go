package machinery

import (
	"errors"
	"fmt"
	"github.com/RichardKnop/machinery/v2"
	amqpBackend "github.com/RichardKnop/machinery/v2/backends/amqp"
	redisBackend "github.com/RichardKnop/machinery/v2/backends/redis"
	amqpBroker "github.com/RichardKnop/machinery/v2/brokers/amqp"
	redisBroker "github.com/RichardKnop/machinery/v2/brokers/redis"
	"github.com/RichardKnop/machinery/v2/config"
	eagerLock "github.com/RichardKnop/machinery/v2/locks/eager"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

const DefaultQueue = "assistant_tasks"
const AMQPExchange = "machinery_exchange"
const AMQPBindingKey = "machinery_task"

type Options struct {
	// amqp
	URL string `yaml:"url"`

	// redis
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
}

func NewOptions(v *viper.Viper) (*Options, error) {
	var err error
	o := new(Options)
	if err = v.UnmarshalKey("rabbitmq", o); err != nil {
		return nil, errors.New("unmarshal machinery redis option error")
	}
	if err = v.UnmarshalKey("redis", o); err != nil {
		return nil, errors.New("unmarshal machinery redis option error")
	}

	return o, err
}

func New(o *Options) (*machinery.Server, error) {
	// use rabbitmq or redis
	if o.URL != "" {
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

		server := machinery.NewServer(cnf, amqpBroker.New(cnf), amqpBackend.New(cnf), eagerLock.New())
		return server, nil
	} else {
		cnf := &config.Config{
			DefaultQueue:    DefaultQueue,
			ResultsExpireIn: 3600,
			Redis: &config.RedisConfig{
				MaxIdle:                3,
				IdleTimeout:            240,
				ReadTimeout:            15,
				WriteTimeout:           15,
				ConnectTimeout:         15,
				NormalTasksPollPeriod:  1000,
				DelayedTasksPollPeriod: 500,
			},
		}
		broker := redisBroker.NewGR(cnf, []string{fmt.Sprintf("%s@%s", o.Password, o.Addr)}, 0)
		backend := redisBackend.NewGR(cnf, []string{fmt.Sprintf("%s@%s", o.Password, o.Addr)}, 0)
		lock := eagerLock.New()

		server := machinery.NewServer(cnf, broker, backend, lock)
		return server, nil
	}
}

var ProviderSet = wire.NewSet(New, NewOptions)
