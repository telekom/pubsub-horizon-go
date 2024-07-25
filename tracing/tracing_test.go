// Copyright 2024 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package tracing

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/contrib/propagators/b3"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"os"
	"testing"
	"time"
)

var (
	traceExporter *tracetest.InMemoryExporter
)

func TestMain(m *testing.M) {
	traceExporter = configureTestProvider()
	os.Exit(m.Run())
}

func configureTestProvider() *tracetest.InMemoryExporter {
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
	otel.SetTextMapPropagator(b3Propagator)

	return exporter
}

func TestTraceContext_StartSpan(t *testing.T) {
	var assertions = assert.New(t)
	var traceCtx = NewTraceContext(context.Background(), "myservice", false)
	defer traceExporter.Reset()

	traceCtx.StartSpan("myspan")
	traceCtx.EndCurrentSpan()

	var snapshots = traceExporter.GetSpans().Snapshots()
	assertions.Equal(1, len(snapshots))

	var firstSpan = snapshots[0]
	assertions.Equal("myspan", firstSpan.Name())
	assertions.Equal(0, firstSpan.ChildSpanCount())
	assertions.LessOrEqual(firstSpan.EndTime(), time.Now())

	assertions.Nil(traceCtx.CurrentSpan())
	assertions.NotNil(traceCtx.LastSpan())
}

func TestTraceContext_StartDetailedSpan(t *testing.T) {
	t.Run("detailed enabled", func(t *testing.T) {
		var assertions = assert.New(t)
		var detailedTraceCtx = NewTraceContext(context.Background(), "myservice", true)
		defer traceExporter.Reset()

		detailedTraceCtx.StartDetailedSpan("mydetailedspan")
		detailedTraceCtx.EndCurrentDetailedSpan()

		var snapshots = traceExporter.GetSpans().Snapshots()
		assertions.Equal(1, len(snapshots))
		assertions.Equal("mydetailedspan", snapshots[0].Name())
	})

	t.Run("detailed disabled", func(t *testing.T) {
		var assertions = assert.New(t)
		var regularTraceCtx = NewTraceContext(context.Background(), "myservice", false)
		defer traceExporter.Reset()

		regularTraceCtx.StartDetailedSpan("mydetailedspan")
		regularTraceCtx.EndCurrentDetailedSpan()

		var snapshots = traceExporter.GetSpans().Snapshots()
		assertions.Equal(0, len(snapshots))
	})
}
