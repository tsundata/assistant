package rpc

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc/registry"
	"github.com/tsundata/assistant/internal/pkg/utils"
	"go.uber.org/zap"
	"net"
	"net/http"
)

type RegistryOptions struct {
	Port int
}

func NewRegistryOptions(v *viper.Viper) (*RegistryOptions, error) {
	var (
		err error
		o   = new(RegistryOptions)
	)

	if err = v.UnmarshalKey("registry", o); err != nil {
		return nil, err
	}

	return o, err
}

type Registry struct {
	o      *RegistryOptions
	logger *zap.Logger
	app    string
	host   string
	port   int
	server *http.Server
}

type InitRegistry func(s *http.Server)

func NewRegistry(o *RegistryOptions, logger *zap.Logger, init InitRegistry) (*Registry, error) {
	return &Registry{
		o:      o,
		logger: logger,
	}, nil
}

func (r *Registry) Application(name string) {
	r.app = name
}

func (r *Registry) Start() error {
	r.port = r.o.Port
	if r.port == 0 {
		r.port = utils.GetAvailablePort()
	}

	// FIXME
	// r.host = utils.GetLocalIP4()
	r.host = "127.0.0.1"
	if r.host == "" {
		return errors.New("get local ipv4 error")
	}

	addr := fmt.Sprintf("%s:%d", r.host, r.port)

	r.logger.Info("rpc registry starting ... " + addr)

	tcpAddr, err := net.ResolveTCPAddr("tcp4", addr)
	if err != nil {
		r.logger.Error(err.Error())
		return err
	}

	l, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		r.logger.Error(err.Error())
		return err
	}
	defer l.Close()

	registryServer := registry.DefaultRegister
	for {
		c, err := l.Accept()
		if err != nil {
			r.logger.Error(err.Error())
			continue
		}

		go registryServer.HandleConnection(c)
	}
}

func (r *Registry) Stop() error {
	return nil
}
