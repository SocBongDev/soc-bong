package otel

import (
	"context"
	"os"

	"github.com/SocBongDev/soc-bong/internal/config"
	"github.com/SocBongDev/soc-bong/internal/logger"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/log/global"
	"go.opentelemetry.io/otel/propagation"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/resource"
	oteltrace "go.opentelemetry.io/otel/sdk/trace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"go.opentelemetry.io/otel/trace"
)

type (
	OtelConfig interface {
		Shutdown(context.Context) error
	}
	otelConfig struct {
		logExp  *otlploghttp.Exporter
		lp      *sdklog.LoggerProvider
		spanExp oteltrace.SpanExporter
		tp      *sdktrace.TracerProvider
	}
)

var (
	tracer       trace.Tracer
	otlpEndpoint string
	_            OtelConfig = (*otelConfig)(nil)
)

func (o *otelConfig) Shutdown(ctx context.Context) error {
	if err := o.logExp.Shutdown(ctx); err != nil {
		return err
	}
	if err := o.lp.Shutdown(ctx); err != nil {
		return err
	}
	if err := o.spanExp.Shutdown(ctx); err != nil {
		return err
	}
	if err := o.tp.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}

func init() {
	otlpEndpoint = config.GetConfig().OtelEndpoint
	if otlpEndpoint == "" {
		logger.Error("You MUST set OTLP_ENDPOINT env variable!")
		panic("You MUST set OTLP_ENDPOINT env variable!")
	}
}

func newTraceExporter(ctx context.Context) (oteltrace.SpanExporter, error) {
	// Change default HTTPS -> HTTP
	insecureOpt := otlptracehttp.WithInsecure()

	// Update default OTLP reciver endpoint
	endpointOpt := otlptracehttp.WithEndpoint(otlpEndpoint)

	return otlptracehttp.New(ctx, insecureOpt, endpointOpt)
}

func newStdoutExporter(ctx context.Context) (oteltrace.SpanExporter, error) {
	return stdouttrace.New(stdouttrace.WithPrettyPrint(), stdouttrace.WithWriter(os.Stdout))
}

func newTraceProvider(exp sdktrace.SpanExporter) *sdktrace.TracerProvider {
	// Ensure default SDK resources and the required service name are set.
	r, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("soc-bong"),
		),
	)
	if err != nil {
		panic(err)
	}

	return sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(r),
	)
}

func newLogExporter(ctx context.Context) *otlploghttp.Exporter {
	logExporter, err := otlploghttp.New(
		ctx,
		otlploghttp.WithEndpoint(otlpEndpoint),
		otlploghttp.WithInsecure(),
	)
	if err != nil {
		panic("failed to initialize exporter")
	}

	return logExporter
}

func newLogProvider(exp sdklog.Exporter) *sdklog.LoggerProvider {
	// Ensure default SDK resources and the required service name are set.
	r, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("soc-bong"),
		),
	)
	if err != nil {
		panic(err)
	}

	return sdklog.NewLoggerProvider(
		sdklog.WithProcessor(sdklog.NewBatchProcessor(exp)),
		sdklog.WithResource(r),
	)
}

func New(ctx context.Context) (*otelConfig, error) {
	spanExp, err := newTraceExporter(ctx)
	if err != nil {
		return nil, err
	}

	logExp := newLogExporter(ctx)
	tp, lp := newTraceProvider(spanExp), newLogProvider(logExp)

	otel.SetTracerProvider(tp)
	global.SetLoggerProvider(lp)

	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{},
			propagation.Baggage{},
		),
	)

	tracer = tp.Tracer("soc-bong")
	return &otelConfig{logExp, lp, spanExp, tp}, nil
}

func TracerStart(ctx context.Context, spanName string, otps ...trace.SpanStartOption) (context.Context, trace.Span) {
	return tracer.Start(ctx, spanName, otps...)
}

func SetTracerProvider(tp trace.TracerProvider) {
	otel.SetTracerProvider(tp)
}
