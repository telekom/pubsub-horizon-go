// Copyright 2024 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package tracing

import (
	"context"
	"github.com/IBM/sarama"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"strings"
)

func WithB3FromMessage(ctx context.Context, msg *sarama.ConsumerMessage) context.Context {
	var carrier = propagation.HeaderCarrier{}
	var propagator = otel.GetTextMapPropagator()

	for _, header := range msg.Headers {
		var key = string(header.Key)
		if !strings.HasPrefix(strings.ToLower(key), "x-b3") {
			continue
		}

		var value = string(header.Value)
		carrier.Set(key, value)
	}

	return propagator.Extract(ctx, carrier)
}
