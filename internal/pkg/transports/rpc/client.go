package rpc

import (
	"context"
	"time"

	"github.com/tsundata/rpc/xclient"
)

type ClientOptions struct {
	Wait     time.Duration
	Tag      string
	Registry string
}

func NewClientOptions() (*ClientOptions, error) {
	var (
		err error
		o   = new(ClientOptions)
	)
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
	xc *xclient.XClient
}

func NewClient(o *ClientOptions) (*Client, error) {
	d := xclient.NewRegistryDiscovery(o.Registry, 0)
	xc := xclient.NewXClient(d, xclient.RandomSelect, nil)

	return &Client{
		o:  o,
		xc: xc,
	}, nil
}

func (c *Client) Call(ctx context.Context, serviceMethod string, args, reply interface{}) error {
	return c.xc.Call(ctx, serviceMethod, args, reply)
}

func (c *Client) Close() error {
	return c.xc.Close()
}
