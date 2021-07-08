package rpc

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpczap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpcrecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/hashicorp/consul/api"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/tsundata/assistant/internal/pkg/middleware/influx"
	redisMiddle "github.com/tsundata/assistant/internal/pkg/middleware/redis"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc/discovery"
	"github.com/tsundata/assistant/internal/pkg/util"
	"github.com/tsundata/assistant/internal/pkg/vendors/rollbar"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"
)

type Server struct {
	conf   *config.AppConfig
	logger *logger.Logger
	server *grpc.Server
	in     influxdb2.Client
	consul *api.Client
}

type InitServer func(s *grpc.Server)

func NewServer(opt *config.AppConfig, logger *logger.Logger, init InitServer, tracer opentracing.Tracer, in influxdb2.Client, rdb *redis.Client, consul *api.Client) (*Server, error) {
	// recovery
	recoveryOpts := []grpcrecovery.Option{
		grpcrecovery.WithRecoveryHandler(func(p interface{}) (err error) {
			return status.Errorf(codes.Unknown, "panic triggered: %v", p)
		}),
	}

	gs := grpc.NewServer(
		grpc.StreamInterceptor(
			grpcmiddleware.ChainStreamServer(
				rollbar.StreamServerInterceptor(),
				influx.StreamServerInterceptor(in, opt.Influx.Org, opt.Influx.Bucket),
				grpczap.StreamServerInterceptor(logger.Zap),
				grpcrecovery.StreamServerInterceptor(recoveryOpts...),
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
				otgrpc.OpenTracingServerInterceptor(tracer),
				redisMiddle.StatsUnaryServerInterceptor(rdb),
			),
		),
	)
	init(gs)

	return &Server{
		conf:   opt,
		logger: logger,
		server: gs,
		in:     in,
		consul: consul,
	}, nil
}

func (s *Server) Start() error {
	if s.conf.Rpc.Port == 0 {
		s.conf.Rpc.Port = util.GetAvailablePort()
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
			s.logger.Error(err)
		}
	}()

	if err := s.register(); err != nil {
		return errors.Wrap(err, "register grpc server error")
	}

	// metrics
	go influx.PushGoServerMetrics(s.in, s.conf.Name, s.conf.Influx.Org, s.conf.Influx.Bucket)

	return nil
}

func (s *Server) register() error {
	rpcAddr := fmt.Sprintf("%s:%d", s.conf.Rpc.Host, s.conf.Rpc.Port)
	s.logger.Info("register rpc service ... ", zap.String("addr", rpcAddr), zap.String("uuid", s.conf.ID))

	// discovery
	discovery.RegisterService(rpcAddr, &discovery.ConsulService{
		ID:   s.conf.ID,
		IP:   s.conf.Rpc.Host,
		Port: s.conf.Rpc.Port,
		Tag:  []string{"grpc"},
		Name: s.conf.Name,
	})

	// Health Check
	// grpc_health_v1.RegisterHealthServer(s.server, &discovery.HealthImpl{})

	return nil
}

func (s *Server) deregister() error {
	for range s.server.GetServiceInfo() {
		id := fmt.Sprintf("%v/%v:%v", s.conf.Name, s.conf.Rpc.Host, s.conf.Rpc.Port)

		err := s.consul.Agent().ServiceDeregister(id)
		if err != nil {
			return errors.Wrapf(err, "deregister service error[id=%s]", id)
		}
		s.logger.Info("deregister service success", zap.String("id", id))
	}
	return nil
}

func (s *Server) Register(f func(gs *grpc.Server) error) error {
	return f(s.server)
}

func (s *Server) Stop() error {
	s.logger.Info("grpc server stopping ...")
	if err := s.deregister(); err != nil {
		return errors.Wrap(err, "deregister grpc server error")
	}
	s.server.GracefulStop()
	return nil
}

var ProviderSet = wire.NewSet(NewServer, NewClient, NewClientOptions)
