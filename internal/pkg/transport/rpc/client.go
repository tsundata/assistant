package rpc

import (
	"context"
	"fmt"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/hashicorp/consul/api"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"github.com/opentracing/opentracing-go"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"google.golang.org/grpc"
	"os"
	"time"
)

type ClientOptions struct {
	Wait            time.Duration
	Tag             string
	GrpcDialOptions []grpc.DialOption
}

func NewClientOptions(tracer opentracing.Tracer) (*ClientOptions, error) {
	var (
		err error
		o   = new(ClientOptions)
	)

	o.GrpcDialOptions = append(o.GrpcDialOptions,
		grpc.WithBlock(),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(
			grpcMiddleware.ChainUnaryClient(
				otgrpc.OpenTracingClientInterceptor(tracer),
			),
		),
		grpc.WithStreamInterceptor(
			grpcMiddleware.ChainStreamClient(
				otgrpc.OpenTracingStreamClientInterceptor(tracer),
			),
		),
	)

	return o, err
}

type ClientOptional func(o *ClientOptions)

//WithTimeout timeout
func WithTimeout(d time.Duration) ClientOptional {
	return func(o *ClientOptions) {
		o.Wait = d
	}
}

//WithTag tag info
func WithTag(tag string) ClientOptional {
	return func(o *ClientOptions) {
		o.Tag = tag
	}
}

type Client struct {
	o      *ClientOptions
	consul *api.Client
	logger *logger.Logger
}

func NewClient(o *ClientOptions, consul *api.Client, logger *logger.Logger) (*Client, error) {
	return &Client{
		o:      o,
		consul: consul,
		logger: logger,
	}, nil
}

func (c *Client) Dial(service string, options ...ClientOptional) (*grpc.ClientConn, error) {
	o := &ClientOptions{
		Wait:            c.o.Wait,
		Tag:             c.o.Tag,
		GrpcDialOptions: c.o.GrpcDialOptions,
	}
	for _, option := range options {
		option(o)
	}

	ctx, cancel := context.WithTimeout(context.Background(), o.Wait)
	defer cancel()

	consulAddress := os.Getenv("CONSUL_ADDRESS")
	target := fmt.Sprintf("consul://%s/%s?wait=%s&tag=%s", consulAddress, service, o.Wait, o.Tag)

	conn, err := grpc.DialContext(ctx, target, o.GrpcDialOptions...)
	if err != nil {
		// return nil, errors.Wrap(err, "grpc dial error") // fixme timeout
	}

	return conn, nil
}
