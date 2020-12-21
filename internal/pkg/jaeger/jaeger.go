package jaeger

import (
	"github.com/opentracing/opentracing-go"
	"github.com/spf13/viper"
	"github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-lib/metrics/prometheus"
	"go.uber.org/zap"
)

func NewConfiguration(v *viper.Viper, logger *zap.Logger) (*config.Configuration, error) {
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
	metricsFactory := prometheus.New()
	tracer, _, err := c.NewTracer(config.Metrics(metricsFactory))

	if err != nil {
		return nil, err
	}

	return tracer, nil
}
