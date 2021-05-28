package etcd

import (
	"github.com/google/wire"
	"github.com/tsundata/assistant/internal/pkg/config"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func New(c *config.AppConfig) (*clientv3.Client, error) {
	return clientv3.New(clientv3.Config{
		Endpoints:   []string{c.Etcd.Url},
		Username:    c.Etcd.Username,
		Password:    c.Etcd.Password,
		DialTimeout: time.Minute,
	})
}

var ProviderSet = wire.NewSet(New)
