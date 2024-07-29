// Copyright 2024 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package tracing

import (
	"context"
	"github.com/IBM/sarama"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWithB3FromMessage(t *testing.T) {
	var assertions = assert.New(t)

	var (
		dummyTraceId      = "314880ec28a27f39a0de087bf4c1d6f6"
		dummyParentSpanId = "d34acdaa84a7183f"
		dummySpanId       = "3659192a527bba23"
	)

	var dummyMessage = sarama.ConsumerMessage{}
	dummyMessage.Headers = []*sarama.RecordHeader{
		{
			Key:   []byte("X-B3-TraceId"),
			Value: []byte(dummyTraceId),
		},
		{
			Key:   []byte("X-B3-ParentSpanId"),
			Value: []byte(dummyParentSpanId),
		},
		{
			Key:   []byte("X-B3-SpanId"),
			Value: []byte(dummySpanId),
		},
		{
			Key:   []byte("X-B3-Sampled"),
			Value: []byte("1"),
		},
		{
			Key:   []byte("foo"),
			Value: []byte("bar"),
		},
	}

	var ctx = WithB3FromMessage(context.Background(), &dummyMessage)
	var traceCtx = NewTraceContext(ctx, "myservice", false)
	defer traceExporter.Reset()

	traceCtx.StartSpan("myspan")
	traceCtx.EndCurrentSpan()

	var lastSpan = traceCtx.LastSpan()
	assertions.Equal(dummyTraceId, lastSpan.SpanContext().TraceID().String())

	var snapshots = traceExporter.GetSpans().Snapshots()
	assertions.Len(snapshots, 1)

	var parentTraceId = snapshots[0].Parent().TraceID()
	assertions.Equal(dummyTraceId, parentTraceId.String())
}

func TestWithB3FromMap(t *testing.T) {
	var assertions = assert.New(t)

	var (
		dummyTraceId      = "314880ec28a27f39a0de087bf4c1d6f6"
		dummyParentSpanId = "d34acdaa84a7183f"
		dummySpanId       = "3659192a527bba23"
	)

	var dummyMap = map[string]string{
		"X-B3-TraceId":      dummyTraceId,
		"X-B3-ParentSpanId": dummyParentSpanId,
		"X-B3-SpanId":       dummySpanId,
		"X-B3-Sampled":      "1",
		"foo":               "bar",
	}

	var ctx = WithB3FromMap(context.Background(), dummyMap)
	var traceCtx = NewTraceContext(ctx, "myservice", false)
	defer traceExporter.Reset()

	traceCtx.StartSpan("myspan")
	traceCtx.EndCurrentSpan()

	var lastSpan = traceCtx.LastSpan()
	assertions.Equal(dummyTraceId, lastSpan.SpanContext().TraceID().String())

	var snapshots = traceExporter.GetSpans().Snapshots()
	assertions.Len(snapshots, 1)

	var parentTraceId = snapshots[0].Parent().TraceID()
	assertions.Equal(dummyTraceId, parentTraceId.String())
}

func TestDumpToB3Map(t *testing.T) {
	var assertions = assert.New(t)

	var traceCtx = NewTraceContext(context.Background(), "myservice", false)
	defer traceExporter.Reset()

	traceCtx.StartSpan("myspan")
	traceCtx.EndCurrentSpan()

	var lastSpan = traceCtx.LastSpan()
	assertions.NotNil(lastSpan)

	var traceId, spanId = lastSpan.SpanContext().TraceID().String(), lastSpan.SpanContext().SpanID().String()
	var dump = DumpToB3Map(traceCtx)
	assertions.Equal(traceId, dump["X-B3-Traceid"])
	assertions.Equal(spanId, dump["X-B3-Spanid"])
}
