package config

import (
	"fmt"
	"github.com/google/wire"
	"github.com/hashicorp/consul/api"
	"github.com/tsundata/assistant/internal/pkg/util"
	"gopkg.in/yaml.v2"
	"strings"
	"sync"
	"time"
)

type AppConfig struct {
	kv   *api.KV
	once sync.Once

	ID   string
	Name string `json:"name"`

	Http    Http    `json:"http"`
	Rpc     Rpc     `json:"rpc"`
	Web     Web     `json:"web"`
	Gateway Gateway `json:"gateway"`
	Storage Storage `json:"storage"`

	Mysql    Mysql    `json:"mysql"`
	Rqlite   Rqlite   `json:"rqlite"`
	Redis    Redis    `json:"redis"`
	Influx   Influx   `json:"influx"`
	Rabbitmq Rabbitmq `json:"rabbitmq"`
	Jaeger   Jaeger   `json:"jaeger"`
	Nats     Nats     `json:"nats"`

	Slack    Slack    `json:"slack"`
	Rollbar  Rollbar  `json:"rollbar"`
	Telegram Telegram `json:"telegram"`
}

func NewConfig(id string, consul *api.Client) *AppConfig {
	kv := consul.KV()
	var xc AppConfig
	xc.kv = kv
	xc.Name = id

	uuid, err := util.GenerateUUID()
	if err != nil {
		panic(err)
	}
	xc.ID = uuid

	xc.readConfig()
	go xc.watch()

	return &xc
}

func (c *AppConfig) readConfig() {
	// common
	pair, _, err := c.kv.Get("config/common", nil)
	if err != nil {
		panic(err)
	}
	if pair == nil {
		panic("pair nil")
	}
	err = yaml.Unmarshal(pair.Value, &c)
	if err != nil {
		panic(err)
	}

	// app
	pair, _, err = c.kv.Get(fmt.Sprintf("config/%s", c.Name), nil)
	if err != nil {
		panic(err)
	}
	if pair != nil {
		err = yaml.Unmarshal(pair.Value, &c)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (c *AppConfig) watch() {
	c.once.Do(func() {
		tick := time.NewTicker(10 * time.Second)
		for range tick.C {
			c.readConfig()
		}
	})
}

func (c *AppConfig) GetConfig(key string) (string, error) {
	result, _, err := c.kv.Get(fmt.Sprintf("config/%s", key), nil)
	if err != nil {
		return "", err
	}
	if result != nil {
		return util.ByteToString(result.Value), nil
	}
	return "", nil
}

func (c *AppConfig) GetSetting(key string) (string, error) {
	result, _, err := c.kv.Get(fmt.Sprintf("setting/%s", key), nil)
	if err != nil {
		return "", err
	}
	if result != nil {
		return util.ByteToString(result.Value), nil
	}
	return "", nil
}

func (c *AppConfig) SetSetting(key, value string) error {
	_, err := c.kv.Put(&api.KVPair{
		Key:   fmt.Sprintf("setting/%s", key),
		Value: util.StringToByte(value),
	}, nil)
	return err
}

func (c *AppConfig) GetSettings() (map[string]string, error) {
	kvs, _, err := c.kv.List("setting", nil)
	if err != nil {
		return nil, err
	}
	result := make(map[string]string)
	for _, ev := range kvs {
		result[strings.ReplaceAll(ev.Key, "setting/", "")] = util.ByteToString(ev.Value)
	}
	return result, nil
}

var ProviderSet = wire.NewSet(NewConfig)
