package config

import (
	"fmt"
	"github.com/google/wire"
	"github.com/hashicorp/consul/api"
	"gopkg.in/yaml.v3"
	"os"
)

type AppConfig struct {
	Name string `json:"name"`

	Http    Http    `json:"http"`
	Rpc     Rpc     `json:"rpc"`
	Web     Web     `json:"web"`
	Gateway Gateway `json:"gateway"`
	Plugin  Plugin  `json:"plugin"`
	Storage Storage `json:"storage"`

	Mysql    Mysql    `json:"mysql"`
	Redis    Redis    `json:"redis"`
	Influx   Influx   `json:"influx"`
	Rabbitmq Rabbitmq `json:"rabbitmq"`
	Jaeger   Jaeger   `json:"jaeger"`
	Nats     Nats     `json:"nats"`

	Slack    Slack    `json:"slack"`
	Rollbar  Rollbar  `json:"rollbar"`
	Telegram Telegram `json:"telegram"`
}

func NewConfig(consul *api.Client) *AppConfig {
	kv := consul.KV()
	configNamespace := os.Getenv("CONFIG_NAMESPACE")
	if configNamespace == "" {
		panic("config namespace error")
	}
	pair, _, err := kv.Get(fmt.Sprintf("config/%s", configNamespace), nil)
	if err != nil {
		panic(err)
	}

	// Watch todo

	var xc AppConfig
	err = yaml.Unmarshal(pair.Value, &xc)
	if err != nil {
		panic(err)
	}

	return &xc
}

var ProviderSet = wire.NewSet(NewConfig)
