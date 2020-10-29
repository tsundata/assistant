package rpc

import (
	"log"
	"net"

	"github.com/spf13/viper"
	"github.com/tsundata/rpc"
	"github.com/tsundata/rpc/registry"
)

type ServerOptions struct {
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
	app      string
	registry string
	server   *rpc.Server
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
	s.registry = s.o.Registry

	go func() {
		l, err := net.Listen("tcp", ":0")
		if err != nil {
			log.Println(err)
		}

		log.Println("rpc server starting ...", "tcp@"+l.Addr().String())

		// s.server.Register()
		registry.Heartbeat(s.registry, "tcp@"+l.Addr().String(), 0)
		s.server.Accept(l)
	}()

	return nil
}

func (s *Server) Register(rcvr interface{}) {
	s.server.Register(rcvr)
}

// TODO
func (s *Server) Stop() error {
	return nil
}
