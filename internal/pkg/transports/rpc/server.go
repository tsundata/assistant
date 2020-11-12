package rpc

import (
	"errors"
	"fmt"
	"github.com/smallnest/rpcx/server"
	"github.com/spf13/viper"
	"github.com/tsundata/assistant/internal/pkg/rpc/registry"
	"github.com/tsundata/assistant/internal/pkg/utils"
	"go.uber.org/zap"
)

type ServerOptions struct {
	Host     string
	Port     int
	Registry string
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
	o        *ServerOptions
	logger   *zap.Logger
	app      string
	host     string
	port     int
	registry string
	server   *server.Server
}

type InitServers func(s *server.Server)

func NewServer(o *ServerOptions, logger *zap.Logger, init InitServers) (*Server, error) {
	return &Server{
		o:      o,
		logger: logger,
		server: server.NewServer(),
	}, nil
}

func (s *Server) Application(name string) {
	s.app = name
}

func (s *Server) Start() error {
	s.registry = s.o.Registry
	if s.registry == "" {
		return errors.New("registry error")
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

	go func() {
		registry.Heartbeat(s.registry, s.app, "tcp@"+addr, 0)

		err := s.server.Serve("tcp", addr)
		if err != nil {
			s.logger.Error(err.Error())
		}
	}()

	return nil
}

func (s *Server) Register(rcvr interface{}, metadata string) error {
	return s.server.Register(rcvr, metadata)
}

func (s *Server) Stop() error {
	return nil
}
