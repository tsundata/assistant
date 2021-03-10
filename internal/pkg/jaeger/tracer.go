package jaeger

import (
	"github.com/opentracing/opentracing-go"
	"github.com/spf13/viper"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/uber/jaeger-client-go/config"
)

func NewConfiguration(v *viper.Viper, logger *logger.Logger) (*config.Configuration, error) {
	var (
		err error
		c   = new(config.Configuration)
	)

	if err = v.UnmarshalKey("jaeger", c); err != nil {
		return nil, err
	}

	logger.Info("load jaeger configuration success")

	return c, nil
}

func New(c *config.Configuration) (opentracing.Tracer, error) {
	tracer, _, err := c.NewTracer()

	if err != nil {
		return nil, err
	}

	return tracer, nil
}
