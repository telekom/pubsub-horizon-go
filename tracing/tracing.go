// Copyright 2024 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package tracing

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// TraceContext provides wrapper functions for tracing flows using the opentelemetry standard.
type TraceContext struct {
	tracer   trace.Tracer
	traceCtx context.Context
	detailed bool
	spans    []trace.Span
}

// NewTraceContext creates a new trace for the given service.
// In additional detailed tracing can be enabled.
func NewTraceContext(ctx context.Context, service string, detailed bool) *TraceContext {
	var provider = otel.GetTracerProvider()
	return &TraceContext{
		tracer:   provider.Tracer(service),
		traceCtx: ctx,
		detailed: detailed,
		spans:    make([]trace.Span, 0),
	}
}

// StartSpan starts a new span.
func (c *TraceContext) StartSpan(name string) {
	var span trace.Span
	c.traceCtx, span = c.tracer.Start(c.traceCtx, name)
	c.spans = append(c.spans, span)
}

// StartDetailedSpan starts a span that will only be started if detailed tracing is enabled.
func (c *TraceContext) StartDetailedSpan(name string) {
	if c.detailed {
		c.StartSpan(name)
	}
}

// EndCurrentSpan ends the most recent span that is still recording (hasn't ended).
func (c *TraceContext) EndCurrentSpan() {
	if currentSpan := c.CurrentSpan(); currentSpan != nil {
		currentSpan.End()
	}
}

// EndCurrentDetailedSpan ends the most recent span ONLY if detailed tracing is enabled that is still recording (hasn't ended).
func (c *TraceContext) EndCurrentDetailedSpan() {
	if c.detailed {
		if currentSpan := c.CurrentSpan(); currentSpan != nil {
			currentSpan.End()
		}
	}
}

// SetAttribute sets the value of the given key.
func (c *TraceContext) SetAttribute(key string, value string) {
	if currentSpan := c.CurrentSpan(); currentSpan != nil {
		currentSpan.SetAttributes(attribute.String(key, value))
	}
}

// RootSpan returns the root span of the current TraceContext.
func (c *TraceContext) RootSpan() trace.Span {
	if len(c.spans) > 0 {
		return c.spans[0]
	}
	return nil
}

// CurrentSpan returns the most recent span that is still recording (hasn't ended).
func (c *TraceContext) CurrentSpan() trace.Span {
	for i := len(c.spans) - 1; i >= 0; i-- {
		var span = c.spans[i]
		if span.IsRecording() {
			return span
		}
	}
	return nil
}

// GetSpanN retrieves the span with the index of n.
func (c *TraceContext) GetSpanN(n int) trace.Span {
	var exists = (len(c.spans)-1) >= n && n >= 0
	if exists {
		return c.spans[n]
	}
	return nil
}

// LastSpan returns the last span regardless of its recording status or nil if there is none.
func (c *TraceContext) LastSpan() trace.Span {
	return c.GetSpanN(len(c.spans) - 1)
}

// Context returns the context used by the tracer.
func (c *TraceContext) Context() context.Context {
	return c.traceCtx
}
