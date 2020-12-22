package cron

import (
	"errors"
	"github.com/spf13/viper"
	"github.com/tsundata/assistant/internal/pkg/app"
	"github.com/tsundata/assistant/internal/pkg/bot"
	"go.uber.org/zap"
)

type Options struct {
	Name   string
	logger *zap.Logger
}

func NewOptions(v *viper.Viper, logger *zap.Logger) (*Options, error) {
	var err error
	o := new(Options)
	o.logger = logger

	if err = v.UnmarshalKey("app", o); err != nil {
		return nil, errors.New("unmarshal app option error")
	}

	return o, err
}

func NewApp(o *Options, b *bot.Bot) (*app.Application, error) {
	o.logger.Info("start cron bot " + b.Name())

	a, err := app.New(o.Name, o.logger)
	if err != nil {
		return nil, err
	}

	return a, nil
}
