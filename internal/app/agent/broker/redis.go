package broker

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"regexp"
	"strings"
	"time"
)

type Redis struct {
	broker *Broker
	r      *redis.Client
}

func NewRedis(broker *Broker, r *redis.Client) (*Redis, error) {
	s := &Redis{}
	s.broker = broker
	s.r = r

	return s, nil
}

func (b *Redis) Run() {
	t := time.NewTicker(time.Second)

	writeAPI := b.broker.influx.WriteAPI(b.broker.Org, b.broker.Bucket)
	for range t.C {
		// server
		err := writeSection(b, writeAPI, "server")
		if err != nil {
			continue
		}

		// clients
		err = writeSection(b, writeAPI, "clients")
		if err != nil {
			continue
		}

		// memory
		err = writeSection(b, writeAPI, "memory")
		if err != nil {
			continue
		}

		// persistence
		err = writeSection(b, writeAPI, "persistence")
		if err != nil {
			continue
		}

		// stats
		err = writeSection(b, writeAPI, "stats")
		if err != nil {
			continue
		}

		// replication
		err = writeSection(b, writeAPI, "replication")
		if err != nil {
			continue
		}

		// cpu
		err = writeSection(b, writeAPI, "cpu")
		if err != nil {
			continue
		}

		// modules
		err = writeSection(b, writeAPI, "modules")
		if err != nil {
			continue
		}

		// cluster
		err = writeSection(b, writeAPI, "cluster")
		if err != nil {
			continue
		}

		// keyspace
		err = writeSection(b, writeAPI, "keyspace")
		if err != nil {
			continue
		}
	}
}

func writeSection(b *Redis, writeAPI api.WriteAPI, section string) error {
	re, err := regexp.Compile(`^\-?\d+(?:\.\d+)?$`)
	if err != nil {
		b.broker.logger.Error(err.Error())
		return err
	}
	var line strings.Builder
	str, err := b.r.Info(context.Background(), section).Result()
	if err != nil {
		b.broker.logger.Error(err.Error())
		return err
	}
	rows := strings.Split(str, "\r\n")

	line.Grow(300)
	line.WriteString("redis_")
	line.WriteString(section)
	line.WriteString(" ")
	for _, row := range rows {
		kv := strings.Split(row, ":")
		if len(kv) == 2 {
			if strings.Contains(kv[1], ",") {
				values := strings.Split(kv[1], ",")
				for _, ki := range values {
					vKV := strings.Split(ki, "=")
					if len(vKV) == 2 {
						isNumber := re.MatchString(vKV[1])

						line.WriteString(kv[0])
						line.WriteString("_")
						line.WriteString(vKV[0])
						line.WriteString("=")
						if !isNumber {
							line.WriteString(`"`)
						}
						line.WriteString(vKV[1])
						if !isNumber {
							line.WriteString(`"`)
						}
						line.WriteString(",")
					}
				}
			} else {
				isNumber := re.MatchString(kv[1])

				line.WriteString(kv[0])
				line.WriteString("=")
				if !isNumber {
					line.WriteString(`"`)
				}
				line.WriteString(kv[1])
				if !isNumber {
					line.WriteString(`"`)
				}
				line.WriteString(",")
			}
		}
	}
	line.WriteString("ok=1")
	writeAPI.WriteRecord(line.String())
	writeAPI.Flush()
	return nil
}
