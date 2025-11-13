// Copyright 2024 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package tracing

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTraceContext_StartSpan(t *testing.T) {
	assertions := assert.New(t)
	traceCtx := NewTraceContext(context.Background(), "myservice", false)
	defer traceExporter.Reset()

	traceCtx.StartSpan("myspan")
	traceCtx.SetAttribute("foo", "bar")
	traceCtx.EndCurrentSpan()

	snapshots := traceExporter.GetSpans().Snapshots()
	assertions.Equal(1, len(snapshots))

	firstSpan := snapshots[0]
	assertions.Equal("myspan", firstSpan.Name())
	assertions.Equal(0, firstSpan.ChildSpanCount())
	assertions.LessOrEqual(firstSpan.EndTime(), time.Now())

	assertions.Len(firstSpan.Attributes(), 1)
	assertions.Equal("bar", firstSpan.Attributes()[0].Value.AsString())

	assertions.Nil(traceCtx.CurrentSpan())
	assertions.NotNil(traceCtx.LastSpan())
}

func TestTraceContext_StartDetailedSpan(t *testing.T) {
	t.Run("detailed enabled", func(t *testing.T) {
		assertions := assert.New(t)
		detailedTraceCtx := NewTraceContext(context.Background(), "myservice", true)
		defer traceExporter.Reset()

		detailedTraceCtx.StartDetailedSpan("mydetailedspan")
		detailedTraceCtx.EndCurrentDetailedSpan()

		snapshots := traceExporter.GetSpans().Snapshots()
		assertions.Equal(1, len(snapshots))
		assertions.Equal("mydetailedspan", snapshots[0].Name())
	})

	t.Run("detailed disabled", func(t *testing.T) {
		assertions := assert.New(t)
		regularTraceCtx := NewTraceContext(context.Background(), "myservice", false)
		defer traceExporter.Reset()

		regularTraceCtx.StartDetailedSpan("mydetailedspan")
		regularTraceCtx.EndCurrentDetailedSpan()

		snapshots := traceExporter.GetSpans().Snapshots()
		assertions.Equal(0, len(snapshots))
	})
}
