package tracing

import (
	"context"
	"io"
	"log"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"

	appconfig "order-service/pkg/config"
)

func Initialize(cfg appconfig.JaegerConfig) (tracer opentracing.Tracer, closer io.Closer, err error) {
	configuration := config.Configuration{
		ServiceName: cfg.ServiceName,
		Disabled:    !cfg.Enabled,
		Sampler: &config.SamplerConfig{
			Type:  cfg.SamplerType,
			Param: cfg.SamplerParam,
		},
	}

	log.Printf("Initializing Config log %s", cfg.Endpoint)

	tracer, closer, err = configuration.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		log.Fatal("CANT INIT TRACING")
		return
	}
	opentracing.SetGlobalTracer(tracer)
	return
}

func StartSpanFromCtx(ctx context.Context,
	operationName string, opts ...opentracing.StartSpanOption) (opentracing.Span, context.Context) {
	return opentracing.StartSpanFromContext(ctx, operationName, opts...)
}
