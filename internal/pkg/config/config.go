package config

import (
	"encoding/json"
	"github.com/google/wire"
	"github.com/micro-in-cn/XConf/pkg/client/source"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/util/log"
	"os"
)

type AppConfig struct {
	c config.Config

	Name string `json:"name"`

	Http    Http    `json:"http"`
	Rpc     Rpc     `json:"rpc"`
	Web     Web     `json:"web"`
	Gateway Gateway `json:"gateway"`
	Plugin  Plugin  `json:"plugin"`
	Storage Storage `json:"storage"`

	Mysql    Mysql    `json:"mysql"`
	Redis    Redis    `json:"redis"`
	Etcd     Etcd     `json:"etcd"`
	Influx   Influx   `json:"influx"`
	Rabbitmq Rabbitmq `json:"rabbitmq"`
	Jaeger   Jaeger   `json:"jaeger"`
	Nats     Nats     `json:"nats"`

	Slack    Slack    `json:"slack"`
	Rollbar  Rollbar  `json:"rollbar"`
	Telegram Telegram `json:"telegram"`
}

func NewConfig() *AppConfig {
	xconfCluster := os.Getenv("XCONF_CLUSTER")
	xconfNamespace := os.Getenv("XCONF_NAMESPACE")
	xconfURL := os.Getenv("XCONF_URL")
	if xconfCluster == "" || xconfNamespace == "" || xconfURL == "" {
		panic("error xconf")
	}

	c, err := config.NewConfig(
		config.WithSource(
			source.NewSource(
				"assistant",
				xconfCluster,
				xconfNamespace,
				source.WithURL(xconfURL),
			),
		),
	)
	if err != nil {
		panic(err)
	}
	log.Info("read: ", string(c.Get().Bytes()))

	// Watch
	w, err := c.Watch()
	if err != nil {
		panic(err)
	}

	var xc AppConfig
	xc.c = c
	err = json.Unmarshal(c.Get().Bytes(), &xc)
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			v, err := w.Next()
			if err != nil {
				log.Info(err)
			}

			err = json.Unmarshal(v.Bytes(), &xc)
			if err != nil {
				log.Info(err)
			}
		}
	}()

	return &xc
}

var ProviderSet = wire.NewSet(NewConfig)
