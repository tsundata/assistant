package rpc

import (
	"context"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/protocol"
	"github.com/spf13/viper"
	"github.com/tsundata/assistant/internal/pkg/rpc/discovery"
	"time"
)

type ClientOptions struct {
	Wait     time.Duration
	Tag      string
	Registry string
}

func NewClientOptions(v *viper.Viper) (*ClientOptions, error) {
	var (
		err error
		o   = new(ClientOptions)
	)

	if err = v.UnmarshalKey("rpc", o); err != nil {
		return nil, err
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
	xc *client.XClient
}

func NewClient(o *ClientOptions, service, servicePath string) (*Client, error) {
	co := client.DefaultOption
	co.SerializeType = protocol.ProtoBuffer
	d := discovery.NewMultiServiceDiscovery(service, o.Registry)
	xc := client.NewXClient(servicePath, client.Failtry, client.RandomSelect, d, co)
	return &Client{
		xc: &xc,
	}, nil
}

func (c *Client) Call(ctx context.Context, serviceMethod string, args, reply interface{}) error {
	return (*c.xc).Call(ctx, serviceMethod, args, reply)
}

func (c *Client) Broadcast(ctx context.Context, serviceMethod string, args, reply interface{}) error {
	return (*c.xc).Broadcast(ctx, serviceMethod, args, reply)
}

func (c *Client) Close() error {
	return (*c.xc).Close()
}

// FIXME
func (c *Client) Reconnection() {

}
