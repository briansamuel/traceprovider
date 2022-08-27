package traceprovider

import (
	"context"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

type Provider interface {
	Generate() (*tracesdk.TracerProvider, error)
	SetProvider()
	GetTracerWithService(service string) *trace.Tracer
	GetTracer() trace.Tracer
	ParentContext(ctx context.Context, TraceID, SpanID string) context.Context
}
