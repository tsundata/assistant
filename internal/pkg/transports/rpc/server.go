package rpc

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpczap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/go-grpc-middleware/ratelimit"
	grpcrecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/opentracing/opentracing-go"
	"github.com/spf13/viper"
	"github.com/tsundata/assistant/internal/pkg/influx"
	"github.com/tsundata/assistant/internal/pkg/logger"
	redisPkg "github.com/tsundata/assistant/internal/pkg/redis"
	"github.com/tsundata/assistant/internal/pkg/utils"
	"github.com/tsundata/assistant/internal/pkg/vendors/rollbar"
	"go.etcd.io/etcd/clientv3"
	etcdnaming "go.etcd.io/etcd/clientv3/naming"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/naming"
	"google.golang.org/grpc/status"
	"net"
)

type alwaysPassLimiter struct{}

func (*alwaysPassLimiter) Limit() bool {
	return false
}

type Options struct {
	Name string
	Host string
	Port int

	Org    string
	Bucket string
}

func NewOptions(v *viper.Viper) (*Options, error) {
	var (
		err error
		o   = new(Options)
	)

	if err = v.UnmarshalKey("rpc", o); err != nil {
		return nil, err
	}

	if err = v.UnmarshalKey("influx", o); err != nil {
		return nil, errors.New("unmarshal influx option error")
	}

	return o, err
}

type Server struct {
	o        *Options
	logger   *logger.Logger
	resolver *etcdnaming.GRPCResolver
	server   *grpc.Server
	in       influxdb2.Client
}

type InitServers func(s *grpc.Server)

func NewServer(opt *Options, logger *logger.Logger, tracer opentracing.Tracer, etcd *clientv3.Client, in influxdb2.Client, rdb *redis.Client) (*Server, error) {
	// recovery
	recoveryOpts := []grpcrecovery.Option{
		grpcrecovery.WithRecoveryHandler(func(p interface{}) (err error) {
			return status.Errorf(codes.Unknown, "panic triggered: %v", p)
		}),
	}

	// TODO limiter
	limiter := &alwaysPassLimiter{}

	// register discovery
	resolver := &etcdnaming.GRPCResolver{Client: etcd}

	gs := grpc.NewServer(
		grpc.StreamInterceptor(
			grpcmiddleware.ChainStreamServer(
				rollbar.StreamServerInterceptor(),
				influx.StreamServerInterceptor(in, opt.Org, opt.Bucket),
				grpczap.StreamServerInterceptor(logger.Zap),
				grpcrecovery.StreamServerInterceptor(recoveryOpts...),
				ratelimit.StreamServerInterceptor(limiter),
				otgrpc.OpenTracingStreamServerInterceptor(tracer),
				redisPkg.StatsStreamServerInterceptor(rdb),
			),
		),
		grpc.UnaryInterceptor(
			grpcmiddleware.ChainUnaryServer(
				rollbar.UnaryServerInterceptor(),
				influx.UnaryServerInterceptor(in, opt.Org, opt.Bucket),
				grpczap.UnaryServerInterceptor(logger.Zap),
				grpcrecovery.UnaryServerInterceptor(recoveryOpts...),
				ratelimit.UnaryServerInterceptor(limiter),
				otgrpc.OpenTracingServerInterceptor(tracer),
				redisPkg.StatsUnaryServerInterceptor(rdb),
			),
		),
	)

	return &Server{
		o:        opt,
		logger:   logger,
		resolver: resolver,
		server:   gs,
		in:       in,
	}, nil
}

func (s *Server) Application(name string) {
	s.o.Name = name
}

func (s *Server) Start() error {
	if s.o.Port == 0 {
		s.o.Port = utils.GetAvailablePort()
	}

	if s.o.Host == "" {
		s.o.Host = utils.GetLocalIP4()
	}
	if s.o.Host == "" {
		return errors.New("get local ipv4 error")
	}

	addr := fmt.Sprintf("%s:%d", s.o.Host, s.o.Port)

	s.logger.Info("rpc server starting ... " + addr)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		s.logger.Error(err)
		return err
	}

	rpcAddr := fmt.Sprintf("%s:%d", s.o.Host, s.o.Port)
	s.logger.Info("register rpc service ... " + rpcAddr)
	err = s.resolver.Update(context.TODO(), s.o.Name, naming.Update{Op: naming.Add, Addr: rpcAddr}) // nolint
	if err != nil {
		panic(err)
	}

	go func() {
		err = s.server.Serve(lis)
		if err != nil {
			s.logger.Error(err)
		}
	}()

	// metrics
	go influx.PushGoServerMetrics(s.in, s.o.Name, s.o.Org, s.o.Bucket)

	return nil
}

func (s *Server) Register(f func(gs *grpc.Server) error) error {
	return f(s.server)
}

func (s *Server) Stop() error {
	addr := fmt.Sprintf("%s:%d", s.o.Host, s.o.Port)
	err := s.resolver.Update(context.TODO(), s.o.Name, naming.Update{Op: naming.Delete, Addr: addr}) // nolint
	if err != nil {
		s.logger.Error(err)
	}
	s.server.Stop()
	return err
}

var ProviderSet = wire.NewSet(NewServer, NewOptions, NewClient, NewClientOptions)
