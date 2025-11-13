// Copyright 2024 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package tracing

import (
	"context"
	"strings"

	"github.com/IBM/sarama"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

func WithB3FromMessage(ctx context.Context, msg *sarama.ConsumerMessage) context.Context {
	carrier := propagation.HeaderCarrier{}
	propagator := otel.GetTextMapPropagator()

	for _, header := range msg.Headers {
		key := string(header.Key)
		if !strings.HasPrefix(strings.ToLower(key), "x-b3") {
			continue
		}

		value := string(header.Value)
		carrier.Set(key, value)
	}

	return propagator.Extract(ctx, carrier)
}

func WithB3FromMap(ctx context.Context, b3Map map[string]string) context.Context {
	carrier := propagation.HeaderCarrier{}
	propagator := otel.GetTextMapPropagator()

	for key, val := range b3Map {
		if !strings.HasPrefix(strings.ToLower(key), "x-b3") {
			continue
		}
		carrier.Set(key, val)
	}

	return propagator.Extract(ctx, carrier)
}

func DumpToB3Map(traceCtx *TraceContext) map[string]string {
	carrier := propagation.HeaderCarrier{}
	propagator := otel.GetTextMapPropagator()
	propagator.Inject(traceCtx.Context(), carrier)

	b3Map := make(map[string]string)
	for _, key := range carrier.Keys() {
		b3Map[key] = carrier.Get(key)
	}
	return b3Map
}
