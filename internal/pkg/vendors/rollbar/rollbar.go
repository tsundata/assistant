package rollbar

import (
	"errors"
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

func Config(o *Options) {
	rollbar.SetToken(o.Token)
	rollbar.SetEnvironment(o.Environment)
	rollbar.SetCodeVersion(version.Version)
	rollbar.SetServerRoot("github.com/tsundata/assistant")
}
