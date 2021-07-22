package rpc

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpcrecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/hashicorp/consul/api"
	"github.com/newrelic/go-agent/v3/integrations/nrgrpc"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/log"
	redisMiddle "github.com/tsundata/assistant/internal/pkg/middleware/redis"
	"github.com/tsundata/assistant/internal/pkg/util"
	"github.com/tsundata/assistant/internal/pkg/vendors/newrelic"
	"github.com/tsundata/assistant/internal/pkg/vendors/rollbar"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"
)

type Server struct {
	conf   *config.AppConfig
	logger log.Logger
	server *grpc.Server
	consul *api.Client
}

type InitServer func(s *grpc.Server)

func NewServer(opt *config.AppConfig, z *zap.Logger, logger log.Logger, init InitServer, tracer opentracing.Tracer, rdb *redis.Client, consul *api.Client, nc *newrelic.App) (*Server, error) {
	// recovery
	recoveryOpts := []grpcrecovery.Option{
		grpcrecovery.WithRecoveryHandler(func(p interface{}) (err error) {
			return status.Errorf(codes.Unknown, "panic triggered: %v", p)
		}),
	}

	gs := grpc.NewServer(
		grpc.StreamInterceptor(
			grpcmiddleware.ChainStreamServer(
				grpcrecovery.StreamServerInterceptor(recoveryOpts...),
				rollbar.StreamServerInterceptor(),
				grpc_zap.StreamServerInterceptor(z),
				otgrpc.OpenTracingStreamServerInterceptor(tracer),
				redisMiddle.StatsStreamServerInterceptor(rdb),
				nrgrpc.StreamServerInterceptor(nc.Application()),
			),
		),
		grpc.UnaryInterceptor(
			grpcmiddleware.ChainUnaryServer(
				grpcrecovery.UnaryServerInterceptor(recoveryOpts...),
				rollbar.UnaryServerInterceptor(),
				grpc_zap.UnaryServerInterceptor(z),
				otgrpc.OpenTracingServerInterceptor(tracer),
				redisMiddle.StatsUnaryServerInterceptor(rdb),
				nrgrpc.UnaryServerInterceptor(nc.Application()),
			),
		),
	)
	init(gs)

	return &Server{
		conf:   opt,
		logger: logger,
		server: gs,
		consul: consul,
	}, nil
}

func (s *Server) Start() error {
	if s.conf.Rpc.Port == 0 {
		s.conf.Rpc.Port = util.GetAvailablePort()
	}
	if s.conf.Rpc.Port == 0 {
		return errors.New("get available port error")
	}

	if s.conf.Rpc.Host == "" {
		s.conf.Rpc.Host = util.GetLocalIP4()
	}
	if s.conf.Rpc.Host == "" {
		return errors.New("get local ipv4 error")
	}

	addr := fmt.Sprintf("%s:%d", s.conf.Rpc.Host, s.conf.Rpc.Port)
	s.logger.Info("rpc server starting ... ", zap.String("addr", addr), zap.String("uuid", s.conf.ID))

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		s.logger.Error(err)
		return err
	}

	go func() {
		err = s.server.Serve(lis)
		if err != nil {
			s.logger.Fatal(err)
		}
	}()

	return nil
}

func (s *Server) Register(f func(gs *grpc.Server) error) error {
	return f(s.server)
}

func (s *Server) Stop() error {
	s.logger.Info("grpc server stopping ...")
	s.server.GracefulStop()
	return nil
}

var ProviderSet = wire.NewSet(NewServer, NewClient, NewClientOptions)
