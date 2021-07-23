package etcd

import (
	"context"
	"github.com/google/wire"
	"github.com/pkg/errors"
	clientv3 "go.etcd.io/etcd/client/v3"
	"os"
	"time"
)

type Client struct {
	cli *clientv3.Client
}

func New() (*Client, error) {
	etcdAddress := os.Getenv("ETCD_ADDRESS")
	etcdUsername := os.Getenv("ETCD_USERNAME")
	etcdPassword := os.Getenv("ETCD_PASSWORD")
	if etcdAddress == "" {
		return nil, errors.New("empty etcd address")
	}
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{etcdAddress},
		DialTimeout: 5 * time.Second,
		Username:    etcdUsername,
		Password:    etcdPassword,
	})
	if err != nil {
		return nil, err
	}
	return &Client{cli: cli}, nil
}

func (c *Client) Put(ctx context.Context, key, val string, opts ...clientv3.OpOption) (*clientv3.PutResponse, error) {
	return c.cli.Put(ctx, key, val, opts...)
}

func (c *Client) Get(ctx context.Context, key string, opts ...clientv3.OpOption) (*clientv3.GetResponse, error) {
	return c.cli.Get(ctx, key, opts...)
}

var ProviderSet = wire.NewSet(New)
