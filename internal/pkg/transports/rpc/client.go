package rpc

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/opentracing/opentracing-go"
	"github.com/spf13/viper"
	"go.etcd.io/etcd/clientv3"
	etcdnaming "go.etcd.io/etcd/clientv3/naming"
	"google.golang.org/grpc"
	"time"
)

type ClientOptions struct {
	Wait            time.Duration
	Tag             string
	Etcd            string
	GrpcDialOptions []grpc.DialOption
}

func NewClientOptions(v *viper.Viper, tracer opentracing.Tracer) (*ClientOptions, error) {
	var (
		err error
		o   = new(ClientOptions)
	)

	if err = v.UnmarshalKey("rpc", o); err != nil {
		return nil, err
	}

	o.GrpcDialOptions = append(o.GrpcDialOptions,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(
			grpc_middleware.ChainUnaryClient(
				grpc_prometheus.UnaryClientInterceptor,
				otgrpc.OpenTracingClientInterceptor(tracer),
			),
		),
		grpc.WithStreamInterceptor(
			grpc_middleware.ChainStreamClient(
				grpc_prometheus.StreamClientInterceptor,
				otgrpc.OpenTracingStreamClientInterceptor(tracer),
			),
		),
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
	CC *grpc.ClientConn
}

func NewClient(o *ClientOptions, service string) (*Client, error) {
	etcdCli, err := clientv3.NewFromURL(o.Etcd)
	if err != nil {
		return nil, err
	}
	re := &etcdnaming.GRPCResolver{Client: etcdCli}
	rr := grpc.RoundRobin(re) // nolint
	gdOptions := append(o.GrpcDialOptions, grpc.WithBalancer(rr)) // nolint

	conn, err := grpc.Dial(service, gdOptions...)
	if err != nil {
		return nil, err
	}

	return &Client{
		o:  o,
		CC: conn,
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

	etcdCli, err := clientv3.NewFromURL(o.Etcd)
	if err != nil {
		return nil, err
	}
	re := &etcdnaming.GRPCResolver{Client: etcdCli}
	rr := grpc.RoundRobin(re) // nolint
	gdOptions := append(o.GrpcDialOptions, grpc.WithBalancer(rr)) // nolint

	conn, err := grpc.Dial(service, gdOptions...)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (c *Client) Close() error {
	return c.CC.Close()
}
