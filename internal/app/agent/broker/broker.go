package broker

import (
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/logger"
)

type Broker struct {
	c *config.AppConfig
	logger *logger.Logger
	influx influxdb2.Client
}

func NewBroker(c *config.AppConfig, logger *logger.Logger, influx influxdb2.Client) (*Broker, error) {
	var err error
	o := new(Broker)
	o.c = c
	o.logger = logger
	o.influx = influx

	return o, err
}

type Runner interface {
	Run()
}
