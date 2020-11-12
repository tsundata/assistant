package http

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/tsundata/assistant/internal/pkg/utils"
	"log"
	"net/http"
	"time"
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
	router     *gin.Engine
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

type InitControllers func(r *gin.Engine)

func NewRouter(o *Options, init InitControllers) *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	init(r)

	return r
}

func New(o *Options, router *gin.Engine) (*Server, error) {
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

	s.httpServer = http.Server{Addr: addr, Handler: s.router}

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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		return errors.New("shutdown http server error")
	}

	return nil
}
