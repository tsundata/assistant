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
	Host string
	Port int
	Mode string
}

type Server struct {
	o          *Options
	app        string
	host       string
	port       int
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

func NewRouter(o *Options, init fasthttp.RequestHandler) *fasthttp.RequestHandler {
	return &init
}

func New(o *Options, router *fasthttp.RequestHandler) (*Server, error) {
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
	if s.port == 0 {
		s.port = utils.GetAvailablePort()
	}

	s.host = s.o.Host
	if s.host == "" {
		s.host = utils.GetLocalIP4()
	}
	if s.host == "" {
		return errors.New("get local ipv4 error")
	}

	addr := fmt.Sprintf("%s:%d", s.host, s.port)

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
