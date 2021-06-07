package influx

import (
	"github.com/google/wire"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/tsundata/assistant/internal/pkg/config"
)

func New(c *config.AppConfig) (influxdb2.Client, error) {
	client := influxdb2.NewClient(c.Influx.Url, c.Influx.Token)
	return client, nil
}

var ProviderSet = wire.NewSet(New)
