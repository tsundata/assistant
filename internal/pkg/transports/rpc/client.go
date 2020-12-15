package rpc

import (
	"github.com/spf13/viper"
	"go.etcd.io/etcd/clientv3"
	etcdnaming "go.etcd.io/etcd/clientv3/naming"
	"google.golang.org/grpc"
	"time"
)

type ClientOptions struct {
	Wait time.Duration
	Tag  string
	Etcd string
}

func NewClientOptions(v *viper.Viper, options ...ClientOptional) (*ClientOptions, error) {
	var (
		err error
		o   = new(ClientOptions)
	)

	if err = v.UnmarshalKey("rpc", o); err != nil {
		return nil, err
	}

	for _, option := range options {
		option(o)
	}

	return o, err
}

type ClientOptional func(o *ClientOptions)

func WithTimeout(d time.Duration) ClientOptional {
	return func(o *ClientOptions) {
		o.Wait = d
	}
}

type Client struct {
	o  *ClientOptions
	CC *grpc.ClientConn
}

func NewClient(o *ClientOptions, service string) (*Client, error) {
	etcdCli, err := clientv3.NewFromURL(o.Etcd)
	if err != nil {
		return nil, err
	}
	re := &etcdnaming.GRPCResolver{Client: etcdCli}
	rr := grpc.RoundRobin(re)

	conn, err := grpc.Dial(service, grpc.WithBalancer(rr), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &Client{
		CC: conn,
	}, nil
}

func (c *Client) Close() error {
	return c.CC.Close()
}
