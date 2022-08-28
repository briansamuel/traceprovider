package main

import (
	"context"
	"github.com/briansamuel/traceprovider"
	trace "github.com/briansamuel/traceprovider/otel"
	"log"
	"time"
)

func main() {
	var traceProvider traceprovider.Provider
	traceProvider = trace.NewTraceProvider("test-hallo", "http://localhost:14268")
	// Set Provider
	traceProvider.SetProvider()
	// Get Tracer
	traceProvider.GetTracer()
	tracer := trace.Tracer()
	//	span with tracer
	_, span := tracer.Start(context.Background(), "test span OL")

	time.Sleep(3 * time.Second)
	span.End()
	time.Sleep(3 * time.Second)
	log.Println("End")

}
