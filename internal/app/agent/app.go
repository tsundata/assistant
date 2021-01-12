package agent

import (
	"errors"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/spf13/viper"
	"github.com/tsundata/assistant/internal/app/agent/broker"
	"github.com/tsundata/assistant/internal/pkg/app"
	"go.uber.org/zap"
)

type Options struct {
	Name   string
	Org    string
	Bucket string
}

func NewOptions(v *viper.Viper) (*Options, error) {
	var err error
	o := new(Options)

	if err = v.UnmarshalKey("app", o); err != nil {
		return nil, errors.New("unmarshal app option error")
	}

	if err = v.UnmarshalKey("influx", o); err != nil {
		return nil, errors.New("unmarshal influx option error")
	}

	return o, err
}

func NewApp(o *Options, logger *zap.Logger, influx influxdb2.Client, b broker.Broker) (*app.Application, error) {
	logger.Info("start agent " + o.Name)

	a, err := app.New(o.Name, logger)
	if err != nil {
		return nil, err
	}

	b.Init(o.Org, o.Bucket, logger, influx)
	b.Run()

	return a, nil
}
