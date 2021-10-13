package config

import (
	"context"
	"fmt"
	"github.com/appleboy/gorush/config"
	"github.com/google/wire"
	"github.com/tsundata/assistant/internal/pkg/middleware/etcd"
	"github.com/tsundata/assistant/internal/pkg/util"
	clientv3 "go.etcd.io/etcd/client/v3"
	"gopkg.in/yaml.v2"
	"strings"
	"sync"
	"time"
)

type AppConfig struct {
	kv   *etcd.Client
	once sync.Once

	ID   string
	Name string `json:"name" yaml:"name"`

	SvcAddr SvcAddr `json:"svc_addr" yaml:"svc_addr"`
	Http    Http    `json:"http" yaml:"http"`
	Rpc     Rpc     `json:"rpc" yaml:"rpc"`
	Web     Web     `json:"web" yaml:"web"`
	Gateway Gateway `json:"gateway" yaml:"gateway"`
	Storage Storage `json:"storage" yaml:"storage"`
	Jwt     Jwt     `json:"jwt" yaml:"jwt"`

	Mysql  Mysql  `json:"mysql" yaml:"mysql"`
	Redis  Redis  `json:"redis" yaml:"redis"`
	Influx Influx `json:"influx" yaml:"influx"`
	Jaeger Jaeger `json:"jaeger" yaml:"jaeger"`
	Nats   Nats   `json:"nats" yaml:"nats"`

	Slack    Slack    `json:"slack" yaml:"slack"`
	Rollbar  Rollbar  `json:"rollbar" yaml:"rollbar"`
	Telegram Telegram `json:"telegram" yaml:"telegram"`
	Newrelic Newrelic `json:"newrelic" yaml:"newrelic"`

	// Notification
	config.ConfYaml
}

func NewConfig(id string) *AppConfig {
	var xc AppConfig
	xc.Name = id

	uuid := util.UUID()
	xc.ID = uuid

	xc.Rpc = Rpc{
		Host: "0.0.0.0",
		Port: 6012,
	}

	return &xc
}

func (c *AppConfig) readConfig() {
	// common
	resp, err := c.kv.Get(context.Background(), "config/common")
	if err != nil {
		panic(err)
	}
	var value []byte
	for _, ev := range resp.Kvs {
		value = ev.Value
	}

	err = yaml.Unmarshal(value, &c)
	if err != nil {
		panic(err)
	}

	// app
	resp, err = c.kv.Get(context.Background(), fmt.Sprintf("config/%s", c.Name))
	if err != nil {
		panic(err)
	}
	if len(resp.Kvs) > 0 {
		for _, ev := range resp.Kvs {
			value = ev.Value
		}
		err = yaml.Unmarshal(value, &c)
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

func (c *AppConfig) GetConfig(ctx context.Context, key string) (string, error) {
	resp, err := c.kv.Get(ctx, fmt.Sprintf("config/%s", key))
	if err != nil {
		return "", err
	}
	if len(resp.Kvs) > 0 {
		var value []byte
		for _, ev := range resp.Kvs {
			value = ev.Value
		}
		return util.ByteToString(value), nil
	}
	return "", nil
}

func (c *AppConfig) GetSetting(ctx context.Context, key string) (string, error) {
	resp, err := c.kv.Get(ctx, fmt.Sprintf("setting/%s", key))
	if err != nil {
		return "", err
	}
	if len(resp.Kvs) > 0 {
		var value []byte
		for _, ev := range resp.Kvs {
			value = ev.Value
		}
		return util.ByteToString(value), nil
	}
	return "", nil
}

func (c *AppConfig) SetSetting(ctx context.Context, key, value string) error {
	_, err := c.kv.Put(ctx, fmt.Sprintf("setting/%s", key), value)
	return err
}

func (c *AppConfig) GetSettings(ctx context.Context) (map[string]string, error) {
	resp, err := c.kv.Get(ctx, "setting", clientv3.WithPrefix(), clientv3.WithSort(clientv3.SortByKey, clientv3.SortAscend))
	if err != nil {
		return nil, err
	}
	result := make(map[string]string)
	for _, ev := range resp.Kvs {
		result[strings.ReplaceAll(util.ByteToString(ev.Key), "setting/", "")] = util.ByteToString(ev.Value)
	}
	return result, nil
}

var ProviderSet = wire.NewSet(NewConfig)
