package database

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

type Options struct {
	URL string `yaml:"url"`
}

func NewOptions(v *viper.Viper) (*Options, error) {
	var err error
	o := new(Options)
	if err = v.UnmarshalKey("db", o); err != nil {
		return nil, errors.New("unmarshal db option error")
	}

	return o, err
}

func New(o *Options) (*sqlx.DB, error) {
	db, err := sqlx.Connect("mysql", o.URL)
	if err != nil {
		return nil, err
	}
	return db, nil
}

var ProviderSet = wire.NewSet(New, NewOptions)
