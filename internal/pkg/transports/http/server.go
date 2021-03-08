package http

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/spf13/viper"
	"github.com/tsundata/assistant/internal/pkg/influx"
	"github.com/tsundata/assistant/internal/pkg/utils"
	"log"
	"net/http"
)

type Options struct {
	Name string
	Host string
	Port int
	Mode string

	Org    string
	Bucket string
}

func NewOptions(v *viper.Viper) (*Options, error) {
	var (
		err error
		o   = new(Options)
	)

	if err = v.UnmarshalKey("http", o); err != nil {
		return nil, err
	}

	if err = v.UnmarshalKey("influx", o); err != nil {
		return nil, errors.New("unmarshal influx option error")
	}

	return o, err
}

type Server struct {
	o          *Options
	router     func (router fiber.Router)
	httpServer *fiber.App
	in         influxdb2.Client
}

func New(o *Options, router func (router fiber.Router), in influxdb2.Client) (*Server, error) {
	var s = &Server{
		o:      o,
		router: router,
		in:     in,
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

	// server
	s.httpServer = fiber.New()

	// init router
	s.router(s.httpServer)

	go func() {
		if err := s.httpServer.Listen(addr); err != nil && err != http.ErrServerClosed {
			log.Fatal("start http server err", err)
			return
		}
	}()

	if err := s.register(); err != nil {
		log.Fatal("register http server error")
	}

	// metrics
	go influx.PushGoServerMetrics(s.in, s.o.Name, s.o.Org, s.o.Bucket)

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
