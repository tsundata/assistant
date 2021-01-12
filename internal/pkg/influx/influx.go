package influx

import (
	"errors"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/spf13/viper"
)

type Options struct {
	Token  string
	Url    string
}

func NewOptions(v *viper.Viper) (*Options, error) {
	var err error
	o := new(Options)

	if err = v.UnmarshalKey("influx", o); err != nil {
		return nil, errors.New("unmarshal influx option error")
	}

	return o, err
}

func New(o *Options) (influxdb2.Client, error) {
	client := influxdb2.NewClient(o.Url, o.Token)
	return client, nil
}
