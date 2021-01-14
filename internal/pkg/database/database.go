package database

import (
	"errors"
	"github.com/spf13/viper"
	"go.etcd.io/bbolt"
)

type Options struct {
	Path   string `yaml:"path"`
}

func NewOptions(v *viper.Viper) (*Options, error) {
	var err error
	o := new(Options)
	if err = v.UnmarshalKey("db", o); err != nil {
		return nil, errors.New("unmarshal db option error")
	}

	return o, err
}

func New(o *Options) (*bbolt.DB, error) {
	db, err := bbolt.Open(o.Path, 0600, nil)
	if err != nil {
		return nil, err
	}
	return db, nil
}
