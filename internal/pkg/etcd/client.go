package etcd

import (
	"errors"
	"github.com/spf13/viper"
	"go.etcd.io/etcd/clientv3"
)

type Options struct {
	URL      string `yaml:"url"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func NewOptions(v *viper.Viper) (*Options, error) {
	var err error
	o := new(Options)
	if err = v.UnmarshalKey("etcd", o); err != nil {
		return nil, errors.New("unmarshal etcd option error")
	}

	return o, err
}

func New(o *Options) (*clientv3.Client, error) {
	return clientv3.New(clientv3.Config{
		Endpoints:   []string{o.URL},
		Username:    o.Username,
		Password:    o.Password,
		DialTimeout: 0,
	})
}
