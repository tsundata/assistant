package broker

import (
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"go.uber.org/zap"
)

type Broker interface {
	Init(org string, bucket string, logger *zap.Logger, influx influxdb2.Client)
	Run()
}
