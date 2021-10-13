package rpc

import (
	"fmt"
	"github.com/google/wire"
	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcrecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/pkg/errors"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/util"
	"github.com/tsundata/assistant/internal/pkg/vendors/rollbar"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

type Server struct {
	conf   *config.AppConfig
	server *grpc.Server
}

type InitServer func(s *grpc.Server)

func NewServer(opt *config.AppConfig, init InitServer) (*Server, error) {
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
			),
		),
		grpc.UnaryInterceptor(
			grpcmiddleware.ChainUnaryServer(
				grpcrecovery.UnaryServerInterceptor(recoveryOpts...),
				rollbar.UnaryServerInterceptor(),
			),
		),
	)
	init(gs)

	return &Server{
		conf:   opt,
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

	addr := fmt.Sprintf("%s:%d", s.conf.Rpc.Host, s.conf.Rpc.Port)
	log.Println("rpc server starting ... ", zap.String("addr", addr), zap.String("uuid", s.conf.ID))

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Println(err)
		return err
	}

	go func() {
		err = s.server.Serve(lis)
		if err != nil {
			log.Fatal(err)
		}
	}()

	return nil
}

func (s *Server) Register(f func(gs *grpc.Server) error) error {
	return f(s.server)
}

func (s *Server) Stop() error {
	log.Println("grpc server stopping ...")
	s.server.GracefulStop()
	return nil
}

var ProviderSet = wire.NewSet(NewServer, NewClient, NewClientOptions)
