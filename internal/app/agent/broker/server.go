package broker

import (
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"go.uber.org/zap"
	"math/rand"
	"time"
)

type Server struct {
	org    string
	bucket string

	logger *zap.Logger
	influx influxdb2.Client
}

func (b *Server) Init(org string, bucket string, logger *zap.Logger, influx influxdb2.Client) {
	b.org = org
	b.bucket = bucket
	b.logger = logger
	b.influx = influx
}

func (b *Server) Run() {
	t := time.NewTicker(time.Second)
	writeAPI := b.influx.WriteAPI(b.org, b.bucket)
	for {
		select {
		case <-t.C:
			rand.Seed(time.Now().Unix())
			row := fmt.Sprintf("stat,unit=temperature avg=%d,max=%d", rand.Intn(100), rand.Intn(100)+10)
			b.logger.Info(row)
			writeAPI.WriteRecord(row)
			writeAPI.Flush()
		}
	}
}
