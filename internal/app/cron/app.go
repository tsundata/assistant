package cron

import (
	"errors"
	"github.com/spf13/viper"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/rulebot"
	"go.uber.org/zap"
)

type Options struct {
	Name string
}

func NewOptions(v *viper.Viper) (*Options, error) {
	var err error
	o := new(Options)

	if err = v.UnmarshalKey("app", o); err != nil {
		return nil, errors.New("unmarshal app option error")
	}

	return o, err
}

func NewApp(o *Options, logger *zap.Logger, b *rulebot.RuleBot) (*app.Application, error) {
	logger.Info("start cron bot " + b.Name())

	a, err := app.New(o.Name, logger)
	if err != nil {
		return nil, err
	}

	return a, nil
}
