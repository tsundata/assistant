package rpc

import (
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
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/middleware/influx"
	redisMiddle "github.com/tsundata/assistant/internal/pkg/middleware/redis"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc/discovery"
	"github.com/tsundata/assistant/internal/pkg/utils"
	"github.com/tsundata/assistant/internal/pkg/vendors/rollbar"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
	"net"
)

type alwaysPassLimiter struct{}

func (*alwaysPassLimiter) Limit() bool {
	return false
}

type Server struct {
	conf     *config.AppConfig
	logger   *logger.Logger
	server   *grpc.Server
	in       influxdb2.Client
}

type InitServers func(s *grpc.Server)

func NewServer(opt *config.AppConfig, logger *logger.Logger, tracer opentracing.Tracer, in influxdb2.Client, rdb *redis.Client) (*Server, error) {
	// recovery
	recoveryOpts := []grpcrecovery.Option{
		grpcrecovery.WithRecoveryHandler(func(p interface{}) (err error) {
			return status.Errorf(codes.Unknown, "panic triggered: %v", p)
		}),
	}

	// TODO limiter
	limiter := &alwaysPassLimiter{}

	gs := grpc.NewServer(
		grpc.StreamInterceptor(
			grpcmiddleware.ChainStreamServer(
				rollbar.StreamServerInterceptor(),
				influx.StreamServerInterceptor(in, opt.Influx.Org, opt.Influx.Bucket),
				grpczap.StreamServerInterceptor(logger.Zap),
				grpcrecovery.StreamServerInterceptor(recoveryOpts...),
				ratelimit.StreamServerInterceptor(limiter),
				otgrpc.OpenTracingStreamServerInterceptor(tracer),
				redisMiddle.StatsStreamServerInterceptor(rdb),
			),
		),
		grpc.UnaryInterceptor(
			grpcmiddleware.ChainUnaryServer(
				rollbar.UnaryServerInterceptor(),
				influx.UnaryServerInterceptor(in, opt.Influx.Org, opt.Influx.Bucket),
				grpczap.UnaryServerInterceptor(logger.Zap),
				grpcrecovery.UnaryServerInterceptor(recoveryOpts...),
				ratelimit.UnaryServerInterceptor(limiter),
				otgrpc.OpenTracingServerInterceptor(tracer),
				redisMiddle.StatsUnaryServerInterceptor(rdb),
			),
		),
	)

	return &Server{
		conf:     opt,
		logger:   logger,
		server:   gs,
		in:       in,
	}, nil
}

func (s *Server) Start() error {
	if s.conf.Rpc.Port == 0 {
		s.conf.Rpc.Port = utils.GetAvailablePort()
	}

	if s.conf.Rpc.Host == "" {
		s.conf.Rpc.Host = utils.GetLocalIP4()
	}
	if s.conf.Rpc.Host == "" {
		return errors.New("get local ipv4 error")
	}

	addr := fmt.Sprintf("%s:%d", s.conf.Rpc.Host, s.conf.Rpc.Port)

	s.logger.Info("rpc server starting ... " + addr)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		s.logger.Error(err)
		return err
	}

	rpcAddr := fmt.Sprintf("%s:%d", s.conf.Rpc.Host, s.conf.Rpc.Port)
	s.logger.Info("register rpc service ... " + rpcAddr)

	// discovery
	discovery.RegisterService(rpcAddr, &discovery.ConsulService{
		IP:   s.conf.Rpc.Host,
		Port: s.conf.Rpc.Port,
		Tag:  []string{s.conf.Name},
		Name: s.conf.Name,
	})

	// Health Check
	grpc_health_v1.RegisterHealthServer(s.server, &discovery.HealthImpl{})

	go func() {
		err = s.server.Serve(lis)
		if err != nil {
			s.logger.Error(err)
		}
	}()

	// metrics
	go influx.PushGoServerMetrics(s.in, s.conf.Name, s.conf.Influx.Org, s.conf.Influx.Bucket)

	return nil
}

func (s *Server) Register(f func(gs *grpc.Server) error) error {
	return f(s.server)
}

func (s *Server) Stop() error {
	s.server.Stop()
	return nil
}

var ProviderSet = wire.NewSet(NewServer, NewClient, NewClientOptions)
