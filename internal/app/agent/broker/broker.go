package broker

import (
	"errors"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Broker struct {
	Org    string
	Bucket string
	logger *zap.Logger
	influx influxdb2.Client
}

func NewBroker(v *viper.Viper, logger *zap.Logger, influx influxdb2.Client) (*Broker, error) {
	var err error
	o := new(Broker)
	o.logger = logger
	o.influx = influx

	if err = v.UnmarshalKey("influx", o); err != nil {
		return nil, errors.New("unmarshal app option error")
	}

	return o, err
}

type Runner interface {
	Run()
}
