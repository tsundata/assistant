package agent

import (
	"errors"
	"github.com/spf13/viper"
	"github.com/tsundata/assistant/internal/app/agent/broker"
	"github.com/tsundata/assistant/internal/pkg/app"
	"go.uber.org/zap"
)

type Options struct {
	Name   string
}

func NewOptions(v *viper.Viper) (*Options, error) {
	var err error
	o := new(Options)

	if err = v.UnmarshalKey("app", o); err != nil {
		return nil, errors.New("unmarshal app option error")
	}

	return o, err
}

func NewApp(o *Options, logger *zap.Logger, b broker.Runner) (*app.Application, error) {
	logger.Info("start agent " + o.Name)

	a, err := app.New(o.Name, logger)
	if err != nil {
		return nil, err
	}

	b.Run()

	return a, nil
}
