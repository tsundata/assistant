package rollbar

import (
	"errors"
	"github.com/google/wire"
	"github.com/rollbar/rollbar-go"
	"github.com/spf13/viper"
	"github.com/tsundata/assistant/internal/pkg/version"
)

type Options struct {
	Token       string `yaml:"token"`
	Environment string `yaml:"environment"`
}

func NewOptions(v *viper.Viper) (*Options, error) {
	var err error
	o := new(Options)
	if err = v.UnmarshalKey("rollbar", o); err != nil {
		return nil, errors.New("unmarshal rollbar option error")
	}

	return o, err
}

type Rollbar struct {
	o *Options
}

func New(o *Options) *Rollbar {
	return &Rollbar{o: o}
}

func (r *Rollbar) Config() {
	rollbar.SetToken(r.o.Token)
	rollbar.SetEnvironment(r.o.Environment)
	rollbar.SetCodeVersion(version.Version)
	rollbar.SetServerRoot("github.com/tsundata/assistant")
}

var ProviderSet = wire.NewSet(New, NewOptions)
