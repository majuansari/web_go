package telemetry

import (
	"errors"
	"github.com/labstack/gommon/log"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/trace/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/semconv"
	"web/config"
)

// initTracer creates a new trace provider instance and registers it as global trace provider.
func InitTracer(cfg config.TracerConfig) func() {
	// Create and install Jaeger export pipeline.
	var (
		flush func()
		err   error
	)
	switch cfg.Provider {
	case "":
		err = errors.New("no provider given")
	case "jaeger":
		flush, err = initJaeger(cfg)
	default:
		flush, err = initJaeger(cfg)
	}
	if err != nil {
		log.Fatal(err)
	}
	return flush
}

func initJaeger(cfg config.TracerConfig) (func(), error) {
	flush, err := jaeger.InstallNewPipeline(
		jaeger.WithCollectorEndpoint(cfg.ReporterUri),
		jaeger.WithSDKOptions(
			sdktrace.WithSampler(sdktrace.TraceIDRatioBased(1)),
			sdktrace.WithResource(resource.NewWithAttributes(
				semconv.ServiceNameKey.String(cfg.ServiceName),
				attribute.String("exporter", "jaeger"),
				attribute.Float64("float", 312.23),
			)),
		),
	)
	//It’s important to note that if you do not set a propagator, the default is to use the NoOp option,
	//which means that context will not be shared between multiple services.
	//this includes trace identifiers, which ensure that all spans for a single request are part of the same trace,
	//as well as baggage
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	return flush, err
}
