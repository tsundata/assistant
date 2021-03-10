package machinery

import (
	"errors"
	"fmt"
	"github.com/RichardKnop/machinery/v2"
	redisbackend "github.com/RichardKnop/machinery/v2/backends/redis"
	redisbroker "github.com/RichardKnop/machinery/v2/brokers/redis"
	"github.com/RichardKnop/machinery/v2/config"
	eagerlock "github.com/RichardKnop/machinery/v2/locks/eager"
	"github.com/spf13/viper"
)

type Options struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
}

func NewOptions(v *viper.Viper) (*Options, error) {
	var err error
	o := new(Options)
	if err = v.UnmarshalKey("redis", o); err != nil {
		return nil, errors.New("unmarshal machinery redis option error")
	}

	return o, err
}

func New(o *Options) (*machinery.Server, error) {
	cnf := &config.Config{
		DefaultQueue:    "assistant_tasks",
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
	broker := redisbroker.NewGR(cnf, []string{fmt.Sprintf("%s@%s", o.Password, o.Addr)}, 0)
	backend := redisbackend.NewGR(cnf, []string{fmt.Sprintf("%s@%s", o.Password, o.Addr)}, 0)
	lock := eagerlock.New()

	server := machinery.NewServer(cnf, broker, backend, lock)
	return server, nil
}
