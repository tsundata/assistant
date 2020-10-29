package rpc

import (
	"errors"
	"log"
	"net"

	"github.com/tsundata/rpc"
	"github.com/tsundata/rpc/registry"
)

type ServerOptions struct {
	RegistryAddr string
}

func NewServerOptions() (*ServerOptions, error) {
	var (
		err error
		o   = new(ServerOptions)
	)

	return o, err
}

type Server struct {
	o            *ServerOptions
	app          string
	registryAddr string
	server       *rpc.Server
}

type InitServers func(s *rpc.Server)

func NewServer(o *ServerOptions, init InitServers) (*Server, error) {
	return &Server{
		o:      o,
		server: rpc.NewServer(),
	}, nil
}

func (s *Server) Application(name string) {
	s.app = name
}

func (s *Server) Start() error {
	s.registryAddr = s.o.RegistryAddr

	go func() {
		l, err := net.Listen("tcp", ":0")
		if err != nil {
			log.Println(err)
		}

		log.Println("rpc server starting ...", "tcp@"+l.Addr().String())

		// s.server.Register()
		registry.Heartbeat(s.registryAddr, "tcp@"+l.Addr().String(), 0)
		s.server.Accept(l)
	}()

	if err := s.register(); err != nil {
		return errors.New("register grpc server error")
	}

	return nil
}

func (s *Server) register() error {
	return nil
}

// TODO
func (s *Server) Stop() error {
	return nil
}
