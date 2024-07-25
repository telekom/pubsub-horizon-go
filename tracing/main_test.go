// Copyright 2024 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package tracing

import (
	"go.opentelemetry.io/contrib/propagators/b3"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"os"
	"testing"
)

var traceExporter *tracetest.InMemoryExporter

func TestMain(m *testing.M) {
	traceExporter = configureTestTraceExporter()
	os.Exit(m.Run())
}

func configureTestTraceExporter() *tracetest.InMemoryExporter {
	var exporter = tracetest.NewInMemoryExporter()
	var provider = tracesdk.NewTracerProvider(
		tracesdk.WithSyncer(exporter),
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("horizon-go"),
		)),
	)
	otel.SetTracerProvider(provider)

	var b3Propagator = b3.New(b3.WithInjectEncoding(b3.B3MultipleHeader))
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(b3Propagator, propagation.TraceContext{}, propagation.Baggage{}))

	return exporter
}
