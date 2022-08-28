package trace

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	oteltrace "go.opentelemetry.io/otel/trace"
	"log"
	"os"
)

var tracer oteltrace.Tracer

func Tracer() oteltrace.Tracer {
	return tracer
}

type traceProvider struct {
	service       string
	jaegerService string
	tracer        oteltrace.Tracer
}

func NewTraceProvider(service, jaegerService string) *traceProvider {
	return &traceProvider{service: service, jaegerService: jaegerService}

}

func (trace *traceProvider) Generate() (*tracesdk.TracerProvider, error) {

	epJaeger := fmt.Sprintf("%s/api/traces", trace.jaegerService)
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(epJaeger)))
	if err != nil {
		return nil, err
	}

	tp := tracesdk.NewTracerProvider(

		tracesdk.WithSampler(tracesdk.AlwaysSample()),
		tracesdk.WithSampler(tracesdk.TraceIDRatioBased(1)),
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in a Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(trace.service),
			attribute.String("environment", os.Getenv("GIN_MODE")),

		)),
	)
	return tp, nil
}

func (trace *traceProvider) SetProvider() {
	tp, err := trace.Generate()
	if err != nil {
		log.Fatal(err)
	}
	otel.SetTracerProvider(tp)
	tr := otel.Tracer(trace.service)
	trace.tracer = tr
	tracer = tr
}

func (trace *traceProvider) GetTracerWithService(service string) *oteltrace.Tracer {
	tr := otel.Tracer(service)
	return &tr
}

func (trace *traceProvider) GetTracer() oteltrace.Tracer {

	return trace.tracer
}

func (trace *traceProvider) ParentContext(ctx context.Context, TraceID, SpanID string) context.Context {
	traceID, _ := oteltrace.TraceIDFromHex(TraceID)
	spanID, _ := oteltrace.SpanIDFromHex(SpanID)
	var spanContextConfig oteltrace.SpanContextConfig
	spanContextConfig.TraceID = traceID
	spanContextConfig.SpanID = spanID
	spanContext := oteltrace.NewSpanContext(spanContextConfig)
	parentContext := oteltrace.ContextWithSpanContext(ctx, spanContext)
	return parentContext
}
