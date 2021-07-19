package http

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/util"
	"net/http"
)

type Server struct {
	conf       *config.AppConfig
	router     func(router fiber.Router)
	httpServer *fiber.App
	logger     log.Logger
}

func New(conf *config.AppConfig, router func(router fiber.Router), logger log.Logger) (*Server, error) {
	var s = &Server{
		conf:   conf,
		router: router,
		logger: logger,
	}
	return s, nil
}

func (s *Server) Start() error {
	if s.conf.Http.Port == 0 {
		s.conf.Http.Port = util.GetAvailablePort()
	}

	if s.conf.Http.Host == "" {
		s.conf.Http.Host = util.GetLocalIP4()
	}
	if s.conf.Http.Host == "" {
		return errors.New("get local ipv4 error")
	}

	addr := fmt.Sprintf("%s:%d", s.conf.Http.Host, s.conf.Http.Port)

	s.logger.Info("start http server " + addr)

	// server
	s.httpServer = fiber.New()

	// init router
	s.router(s.httpServer)

	go func() {
		if err := s.httpServer.Listen(addr); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.logger.Fatal(fmt.Errorf("start http server err, %v", err))
			return
		}
	}()

	if err := s.register(); err != nil {
		s.logger.Error(errors.New("register http server error"))
	}

	return nil
}

func (s *Server) register() error {
	return nil
}

func (s *Server) Stop() error {
	if err := s.httpServer.Shutdown(); err != nil {
		return errors.Wrap(err, "shutdown http server error")
	}

	return nil
}

var ProviderSet = wire.NewSet(New, NewClient)
