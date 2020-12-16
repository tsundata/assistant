package app

import (
	"errors"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"

	"github.com/tsundata/assistant/internal/pkg/transports/http"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc"
)

type Application struct {
	name       string
	logger     *zap.Logger
	httpServer *http.Server
	rpcServer  *rpc.Server
}

type Option func(app *Application) error

func HTTPServerOption(svr *http.Server) Option {
	return func(app *Application) error {
		svr.Application(app.name)
		app.httpServer = svr

		return nil
	}
}

func RPCServerOption(svr *rpc.Server) Option {
	return func(app *Application) error {
		svr.Application(app.name)
		app.rpcServer = svr

		return nil
	}
}

func New(name string, logger *zap.Logger, options ...Option) (*Application, error) {
	app := &Application{
		name:   name,
		logger: logger,
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
			return errors.New("http server start error")
		}
	}

	if a.rpcServer != nil {
		if err := a.rpcServer.Start(); err != nil {
			return errors.New("rpc server start error")
		}
	}

	return nil
}

func (a *Application) AwaitSignal() {
	c := make(chan os.Signal, 1)
	signal.Reset(syscall.SIGTERM, syscall.SIGINT)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	s := <-c
	a.logger.Info("receive a signal " + s.String())

	if a.httpServer != nil {
		if err := a.httpServer.Stop(); err != nil {
			a.logger.Error("stop http server error " + err.Error())
		}
	}

	if a.rpcServer != nil {
		if err := a.rpcServer.Stop(); err != nil {
			a.logger.Error("stop rpc server error " + err.Error())
		}
	}

	os.Exit(0)
}
