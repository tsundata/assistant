package http

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"github.com/tsundata/assistant/internal/pkg/utils"
	"github.com/valyala/fasthttp"
	"log"
	"net/http"
)

type Options struct {
	Name string
	Host string
	Port int
	Mode string
}

type Server struct {
	o          *Options
	router     *fasthttp.RequestHandler
	httpServer *fasthttp.Server
}

func NewOptions(v *viper.Viper) (*Options, error) {
	var (
		err error
		o   = new(Options)
	)

	if err = v.UnmarshalKey("http", o); err != nil {
		return nil, err
	}

	return o, err
}

func New(o *Options, router *fasthttp.RequestHandler) (*Server, error) {
	var s = &Server{
		router: router,
		o:      o,
	}
	return s, nil
}

func (s *Server) Application(name string) {
	s.o.Name = name
}

func (s *Server) Start() error {
	if s.o.Port == 0 {
		s.o.Port = utils.GetAvailablePort()
	}

	if s.o.Host == "" {
		s.o.Host = utils.GetLocalIP4()
	}
	if s.o.Host == "" {
		return errors.New("get local ipv4 error")
	}

	addr := fmt.Sprintf("%s:%d", s.o.Host, s.o.Port)

	log.Println("start http server ", addr)

	s.httpServer = &fasthttp.Server{Handler: *s.router}

	go func() {
		if err := s.httpServer.ListenAndServe(addr); err != nil && err != http.ErrServerClosed {
			log.Fatal("start http server err", err)
			return
		}
	}()

	if err := s.register(); err != nil {
		log.Fatal("register http server error")
	}
	return nil
}

func (s *Server) register() error {
	return nil
}

func (s *Server) Stop() error {
	if err := s.httpServer.Shutdown(); err != nil {
		return errors.New("shutdown http server error")
	}

	return nil
}
