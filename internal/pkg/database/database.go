package database

import (
	"errors"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Options struct {
	URL   string `yaml:"url"`
	Debug bool
}

func NewOptions(v *viper.Viper) (*Options, error) {
	var err error
	o := new(Options)
	if err = v.UnmarshalKey("db", o); err != nil {
		return nil, errors.New("unmarshal db option error")
	}

	log.Println("load database options success")

	return o, err
}

func New(o *Options) (*gorm.DB, error) {
	var err error
	db, err := gorm.Open(mysql.Open(o.URL), &gorm.Config{})
	if err != nil {
		return nil, errors.New("gorm open database connection error")
	}

	if o.Debug {
		db = db.Debug()
	}

	return db, nil
}
