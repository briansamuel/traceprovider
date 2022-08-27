package traceprovider

import (
	"context"
	trace "github.com/briansamuel/traceprovider/otel"
	"testing"
)

func TestTracer(t *testing.T) {
	var traceProvider Provider
	traceProvider = trace.NewTraceProvider("test", "http://localhost:14268")
	// Set Provider
	traceProvider.SetProvider()
	// Get Tracer
	traceProvider.GetTracer()
	//	span with tracer
	_, span := traceProvider.GetTracer().Start(context.Background(), "test span")
	defer span.End()

}
