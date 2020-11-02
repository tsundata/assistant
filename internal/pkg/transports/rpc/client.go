package rpc

import (
	"context"
	"errors"
	"github.com/spf13/viper"
	"github.com/tsundata/rpc/xclient"
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
	xc map[string]*xclient.XClient
}

func NewClient(o *ClientOptions) (*Client, error) {
	return &Client{
		o: o,
	}, nil
}

func (c *Client) Call(ctx context.Context, servicePath, serviceMethod string, args, reply interface{}) error {
	xc, ok := c.xc[servicePath]
	if !ok {
		d := xclient.NewRegistryDiscovery(c.o.Registry, c.o.Wait)
		xc = xclient.NewXClient(servicePath, d, xclient.RandomSelect, nil)
	}

	return xc.Call(ctx, serviceMethod, args, reply)
}

func (c *Client) Broadcast(ctx context.Context, servicePath, serviceMethod string, args, reply interface{}) error {
	xc, ok := c.xc[servicePath]
	if !ok {
		d := xclient.NewRegistryDiscovery(c.o.Registry, c.o.Wait)
		xc = xclient.NewXClient(servicePath, d, xclient.RandomSelect, nil)
	}

	return xc.Broadcast(ctx, serviceMethod, args, reply)
}

func (c *Client) Close(servicePath string) error {
	xc, ok := c.xc[servicePath]
	if ok {
		delete(c.xc, servicePath)
		return xc.Close()
	}

	return errors.New("error xc")
}
