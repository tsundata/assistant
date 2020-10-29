package app

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/tsundata/assistant/internal/pkg/transports/http"
	"github.com/tsundata/assistant/internal/pkg/transports/rpc"
)

type Application struct {
	name           string
	httpServer     *http.Server
	rpcServer      *rpc.Server
	registryServer *rpc.Registry
}

type Option func(app *Application) error

func HttpServerOption(svr *http.Server) Option {
	return func(app *Application) error {
		svr.Application(app.name)
		app.httpServer = svr

		return nil
	}
}

func RpcServerOption(svr *rpc.Server) Option {
	return func(app *Application) error {
		svr.Application(app.name)
		app.rpcServer = svr

		return nil
	}
}

func RegistryServerOption(svr *rpc.Registry) Option {
	return func(app *Application) error {
		svr.Application(app.name)
		app.registryServer = svr

		return nil
	}
}

func New(name string, options ...Option) (*Application, error) {
	app := &Application{
		name: name,
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

	if a.registryServer != nil {
		if err := a.registryServer.Start(); err != nil {
			return errors.New("registry server start error")
		}
	}

	return nil
}

func (a *Application) AwaitSignal() {
	c := make(chan os.Signal, 1)
	signal.Reset(syscall.SIGTERM, syscall.SIGINT)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	select {
	case s := <-c:
		log.Println("receive a signal", s.String())
		if a.httpServer != nil {
			if err := a.httpServer.Stop(); err != nil {
				log.Println("stop http server error", err)
			}
		}

		if a.rpcServer != nil {
			if err := a.rpcServer.Stop(); err != nil {
				log.Println("stop rpc server error", err)
			}
		}

		if a.registryServer != nil {
			if err := a.registryServer.Stop(); err != nil {
				log.Println("stop registry server error", err)
			}
		}

		os.Exit(0)
	}
}
