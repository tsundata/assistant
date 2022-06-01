package rpc

import (
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcZap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpcRecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/newrelic/go-agent/v3/integrations/nrgrpc"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/tsundata/assistant/internal/pkg/auth"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/log"
	redisMiddle "github.com/tsundata/assistant/internal/pkg/middleware/redis"
	"github.com/tsundata/assistant/internal/pkg/util"
	"github.com/tsundata/assistant/internal/pkg/vendors/newrelic"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
	"net"
	"strconv"
)

type Server struct {
	conf   *config.AppConfig
	logger log.Logger
	server *grpc.Server
}

type InitServer func(s *grpc.Server)

func NewServer(opt *config.AppConfig, z *zap.Logger, logger log.Logger, init InitServer, tracer opentracing.Tracer, rdb *redis.Client, nc *newrelic.App) (*Server, error) {
	// recovery
	recoveryOpts := []grpcRecovery.Option{
		grpcRecovery.WithRecoveryHandler(func(p interface{}) (err error) {
			return status.Errorf(codes.Unknown, "panic triggered: %v", p)
		}),
	}

	gs := grpc.NewServer(
		grpc.StreamInterceptor(
			grpcMiddleware.ChainStreamServer(
				grpcRecovery.StreamServerInterceptor(recoveryOpts...),
				//rollbar.StreamServerInterceptor(),
				grpcZap.StreamServerInterceptor(z),
				otgrpc.OpenTracingStreamServerInterceptor(tracer),
				redisMiddle.StatsStreamServerInterceptor(rdb),
				nrgrpc.StreamServerInterceptor(nc.Application()),
				auth.StreamServerInterceptor(),
			),
		),
		grpc.UnaryInterceptor(
			grpcMiddleware.ChainUnaryServer(
				grpcRecovery.UnaryServerInterceptor(recoveryOpts...),
				//rollbar.UnaryServerInterceptor(),
				grpcZap.UnaryServerInterceptor(z, grpcZap.WithDecider(func(fullMethodName string, err error) bool {
					return fullMethodName != "/grpc.health.v1.Health/Check"
				})),
				otgrpc.OpenTracingServerInterceptor(tracer),
				redisMiddle.StatsUnaryServerInterceptor(rdb),
				nrgrpc.UnaryServerInterceptor(nc.Application()),
				auth.UnaryServerInterceptor(),
			),
		),
	)
	init(gs)

	return &Server{
		conf:   opt,
		logger: logger,
		server: gs,
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

	addr := net.JoinHostPort(s.conf.Rpc.Host, strconv.Itoa(s.conf.Rpc.Port))
	s.logger.Info("rpc server starting ... ", zap.String("addr", addr), zap.String("uuid", s.conf.ID))

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		s.logger.Error(err)
		return err
	}

	// health check
	healthCheck := health.NewServer()
	grpc_health_v1.RegisterHealthServer(s.server, healthCheck)
	healthCheck.SetServingStatus(s.conf.Name, grpc_health_v1.HealthCheckResponse_SERVING)

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
