package http

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/spf13/viper"
	"github.com/tsundata/framework"
	"github.com/tsundata/framework/middleware"
)

type Options struct {
	Host string
	Port int
	Mode string
}

type Server struct {
	o          *Options
	app        string
	host       string
	port       int
	router     *framework.Engine
	httpServer http.Server
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

type InitControllers func(r *framework.Engine)

func NewRouter(o *Options, init InitControllers) *framework.Engine {
	r := framework.New()

	r.Use(middleware.Logger())
	r.Use(middleware.Recovery())

	init(r)

	return r
}

func New(o *Options, router *framework.Engine) (*Server, error) {
	var s = &Server{
		router: router,
		o:      o,
	}
	return s, nil
}

func (s *Server) Application(name string) {
	s.app = name
}

func (s *Server) Start() error {
	s.port = s.o.Port
	s.host = s.o.Host
	addr := fmt.Sprintf("%s:%d", s.host, s.port)

	s.httpServer = http.Server{Addr: addr, Handler: s.router}

	log.Println("http server starting ...", addr)
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
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
	log.Println("stopping http server")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		return errors.New("shutdown http server error")
	}

	return nil
}
