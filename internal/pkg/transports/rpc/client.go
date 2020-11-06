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
	d *xclient.RegistryDiscovery
}

func NewClient(o *ClientOptions) (*Client, error) {
	d := xclient.NewRegistryDiscovery(o.Registry, o.Wait)
	return &Client{
		o: o,
		d: d,
	}, nil
}

func (c *Client) Auth(servicePath, token string)  {
	xc, ok := c.xc[servicePath]
	if !ok {
		xc = xclient.NewXClient(servicePath, c.d, xclient.RoundRobinSelect, nil)
	}

	xc.Auth(token)
}

func (c *Client) Call(ctx context.Context, servicePath, serviceMethod string, args, reply interface{}) error {
	xc, ok := c.xc[servicePath]
	if !ok {
		xc = xclient.NewXClient(servicePath, c.d, xclient.RoundRobinSelect, nil)
	}

	return xc.Call(ctx, serviceMethod, args, reply)
}

func (c *Client) Broadcast(ctx context.Context, servicePath, serviceMethod string, args, reply interface{}) error {
	xc, ok := c.xc[servicePath]
	if !ok {
		xc = xclient.NewXClient(servicePath, c.d, xclient.RoundRobinSelect, nil)
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
