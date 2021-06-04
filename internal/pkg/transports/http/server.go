package http

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/influx"
	"github.com/tsundata/assistant/internal/pkg/utils"
	"log"
	"net/http"
)

type Server struct {
	c          *config.AppConfig
	router     func(router fiber.Router)
	httpServer *fiber.App
	in         influxdb2.Client
}

func New(c *config.AppConfig, router func(router fiber.Router), in influxdb2.Client) (*Server, error) {
	var s = &Server{
		c:      c,
		router: router,
		in:     in,
	}
	return s, nil
}

func (s *Server) Application(name string) {
	// s.c.Name = name fixme
}

func (s *Server) Start() error {
	if s.c.Http.Port == 0 {
		s.c.Http.Port = utils.GetAvailablePort()
	}

	if s.c.Http.Host == "" {
		s.c.Http.Host = utils.GetLocalIP4()
	}
	if s.c.Http.Host == "" {
		return errors.New("get local ipv4 error")
	}

	addr := fmt.Sprintf("%s:%d", s.c.Http.Host, s.c.Http.Port)

	log.Println("start http server ", addr)

	// server
	s.httpServer = fiber.New()

	// init router
	s.router(s.httpServer)

	go func() {
		if err := s.httpServer.Listen(addr); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("start http server err", err)
			return
		}
	}()

	if err := s.register(); err != nil {
		log.Fatal("register http server error")
	}

	// metrics
	go influx.PushGoServerMetrics(s.in, s.c.Name, s.c.Influx.Org, s.c.Influx.Bucket)

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

var ProviderSet = wire.NewSet(New, NewClient)
