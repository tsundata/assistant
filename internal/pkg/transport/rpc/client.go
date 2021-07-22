package rpc

import (
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/newrelic/go-agent/v3/integrations/nrgrpc"
	"github.com/opentracing/opentracing-go"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc/discovery"
	"go.uber.org/zap"
	"google.golang.org/grpc"
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
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(
			grpcMiddleware.ChainUnaryClient(
				otgrpc.OpenTracingClientInterceptor(tracer),
				nrgrpc.UnaryClientInterceptor,
			),
		),
		grpc.WithStreamInterceptor(
			grpcMiddleware.ChainStreamClient(
				otgrpc.OpenTracingStreamClientInterceptor(tracer),
				nrgrpc.StreamClientInterceptor,
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
	conf   *config.AppConfig
	logger log.Logger
}

func NewClient(o *ClientOptions, conf *config.AppConfig, logger log.Logger) (*Client, error) {
	return &Client{
		o:      o,
		conf:   conf,
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

	target := discovery.SvcAddr(c.conf, service)
	conn, err := grpc.Dial(target, o.GrpcDialOptions...)
	if err != nil {
		c.logger.Warn(err.Error(), zap.String("service", service))
		return conn, nil
	}

	return conn, nil
}
