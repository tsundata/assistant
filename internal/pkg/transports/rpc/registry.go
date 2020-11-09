package rpc

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"github.com/tsundata/assistant/internal/pkg/rpc/registry"
	"github.com/tsundata/assistant/internal/pkg/utils"
	"log"
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
	app    string
	host   string
	port   int
	server *http.Server
}

type InitRegistry func(s *http.Server)

func NewRegistry(o *RegistryOptions, init InitRegistry) (*Registry, error) {
	return &Registry{
		o: o,
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

	log.Println("rpc registry starting ...", addr)

	go func() {
		l, err := net.Listen("tcp", addr)
		if err != nil {
			log.Println(err)
			return
		}
		registry.HandleHTTP()
		err = http.Serve(l, nil)
		log.Println(err)
	}()
	return nil
}

func (r *Registry) Stop() error {
	return nil
}
