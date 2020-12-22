package rpc

import (
	"context"
	"errors"
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/go-grpc-middleware/ratelimit"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/opentracing/opentracing-go"
	"github.com/spf13/viper"
	"github.com/tsundata/assistant/internal/pkg/utils"
	"go.etcd.io/etcd/clientv3"
	etcdnaming "go.etcd.io/etcd/clientv3/naming"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/naming"
	"google.golang.org/grpc/status"
	"net"
)

type ServerOptions struct {
	Host string
	Port int
	Etcd string
}

func NewServerOptions(v *viper.Viper) (*ServerOptions, error) {
	var (
		err error
		o   = new(ServerOptions)
	)

	if err = v.UnmarshalKey("rpc", o); err != nil {
		return nil, err
	}

	return o, err
}

type Server struct {
	o      *ServerOptions
	logger *zap.Logger
	app    string
	host   string
	port   int
	etcd   string
	r      *etcdnaming.GRPCResolver
	server *grpc.Server
}

type InitServers func(s *grpc.Server)

func NewServer(o *ServerOptions, logger *zap.Logger, tracer opentracing.Tracer, init InitServers) (*Server, error) {
	// recovery
	recoveryOpts := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandler(func(p interface{}) (err error) {
			return status.Errorf(codes.Unknown, "panic triggered: %v", p)
		}),
	}

	// TODO limiter
	limiter := &alwaysPassLimiter{}

	// register discovery
	cli, err := clientv3.NewFromURL(o.Etcd)
	if err != nil {
		panic(err)
	}
	r := &etcdnaming.GRPCResolver{Client: cli}

	gs := grpc.NewServer(
		grpc.StreamInterceptor(
			grpc_middleware.ChainStreamServer(
				grpc_zap.StreamServerInterceptor(logger),
				grpc_recovery.StreamServerInterceptor(recoveryOpts...),
				ratelimit.StreamServerInterceptor(limiter),
				otgrpc.OpenTracingStreamServerInterceptor(tracer),
				grpc_prometheus.StreamServerInterceptor,
			),
		),
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpc_zap.UnaryServerInterceptor(logger),
				grpc_recovery.UnaryServerInterceptor(recoveryOpts...),
				ratelimit.UnaryServerInterceptor(limiter),
				otgrpc.OpenTracingServerInterceptor(tracer),
				grpc_prometheus.UnaryServerInterceptor,
			),
		),
	)

	return &Server{
		r:      r,
		o:      o,
		logger: logger,
		server: gs,
	}, nil
}

func (s *Server) Application(name string) {
	s.app = name
}

func (s *Server) Start() error {
	s.etcd = s.o.Etcd
	if s.etcd == "" {
		return errors.New("etcd error")
	}

	s.port = s.o.Port
	if s.port == 0 {
		s.port = utils.GetAvailablePort()
	}

	// FIXME
	// s.host = utils.GetLocalIP4()
	s.host = "127.0.0.1"
	if s.host == "" {
		return errors.New("get local ipv4 error")
	}

	addr := fmt.Sprintf("%s:%d", s.host, s.port)

	s.logger.Info("rpc server starting ... " + addr)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}

	err = s.r.Update(context.TODO(), s.app, naming.Update{Op: naming.Add, Addr: addr}) // nolint
	if err != nil {
		panic(err)
	}

	err = s.server.Serve(lis)
	if err != nil {
		s.logger.Error(err.Error())
	}

	return nil
}

func (s *Server) Register(f func(gs *grpc.Server) error) error {
	return f(s.server)
}

func (s *Server) Stop() error {
	addr := fmt.Sprintf("%s:%d", s.host, s.port)
	err := s.r.Update(context.TODO(), s.app, naming.Update{Op: naming.Delete, Addr: addr}) // nolint
	if err != nil {
		return err
	}
	s.server.Stop()
	return nil
}

type alwaysPassLimiter struct{}

func (*alwaysPassLimiter) Limit() bool {
	return false
}
