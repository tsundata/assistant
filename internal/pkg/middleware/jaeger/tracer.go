package jaeger

import (
	"github.com/google/wire"
	"github.com/opentracing/opentracing-go"
	appConfig "github.com/tsundata/assistant/internal/pkg/config"
	"github.com/tsundata/assistant/internal/pkg/logger"
	"github.com/uber/jaeger-client-go/config"
)

func NewConfiguration(c *appConfig.AppConfig, logger *logger.Logger) (*config.Configuration, error) {
	var cc = new(config.Configuration)

	cc.ServiceName = c.Name
	cc.Reporter = &config.ReporterConfig{
		LocalAgentHostPort: c.Jaeger.Reporter.LocalAgentHostPort,
	}
	cc.Sampler = &config.SamplerConfig{
		Type:  c.Jaeger.Sampler.Type,
		Param: c.Jaeger.Sampler.Param,
	}

	logger.Info("load jaeger configuration success")

	return cc, nil
}

func New(c *config.Configuration) (opentracing.Tracer, error) {
	tracer, _, err := c.NewTracer()

	if err != nil {
		return nil, err
	}

	return tracer, nil
}

var ProviderSet = wire.NewSet(New, NewConfiguration)
