# KafkaPubsub
> move by briansamuel/traceprovider helper use concurrent to managerment job test.

[![Go Reference](https://pkg.go.dev/badge/github.com/princjef/gomarkdoc.svg)](https://pkg.go.dev/github.com/briansamuel/traceprovider)

### Install

``` bash
 go get github.com/briansamuel/traceprovider
```

### Usage

Library simple for jaeger trace.
For example to set up:

* Create New traceProvider 
* serviceName mean trace name
* jaegerService example "http://localhost:14268"
``` go
var traceProvider traceprovider.Provider
traceProvider = trace.NewTraceProvider(serviceName, jaegerService)
```


* Set TraceProvider

``` go
trace.SetProvider()
```


* Get Tracer
``` go
tr := traceProvider.GetTracer()
```

* Get ParentContext
``` go

parentContext := traceProvider.ParentContext(ctx, traceID, spanID)
```