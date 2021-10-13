package app

import (
	"github.com/pkg/errors"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/log"
	"github.com/tsundata/assistant/internal/pkg/transport/http"
	"github.com/tsundata/assistant/internal/pkg/transport/rpc"
	"os"
	"os/signal"
	"syscall"
)



type Application struct {
	name       string
	logger     log.Logger
	httpServer *http.Server
	rpcServer  *rpc.Server
}

type Option func(app *Application) error

func HTTPServerOption(svr *http.Server) Option {
	return func(app *Application) error {
		app.httpServer = svr

		return nil
	}
}

func RPCServerOption(svr *rpc.Server) Option {
	return func(app *Application) error {
		app.rpcServer = svr

		return nil
	}
}

func New(c *config.AppConfig, options ...Option) (*Application, error) {
	app := &Application{
		name:   c.Name,
	}

	for _, option := range options {
		if err := option(app); err != nil {
			return nil, err
		}
	}

	return app, nil
}

func (a *Application) Start() error {
	if a.httpServer != nil {
		if err := a.httpServer.Start(); err != nil {
			return errors.Wrap(err, "http server start error")
		}
	}

	if a.rpcServer != nil {
		if err := a.rpcServer.Start(); err != nil {
			return errors.Wrap(err, "rpc server start error")
		}
	}

	return nil
}

func (a *Application) AwaitSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP)

	s := <-c
	a.logger.Info("receive a signal " + s.String())

	if a.httpServer != nil {
		if err := a.httpServer.Stop(); err != nil {
			a.logger.Error(err)
		}
	}

	if a.rpcServer != nil {
		if err := a.rpcServer.Stop(); err != nil {
			a.logger.Error(err)
		}
	}

	a.logger.Info("Complete end")
	os.Exit(0)
}
