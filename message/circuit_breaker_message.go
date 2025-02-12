// Copyright 2024 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package message

import (
	"github.com/telekom/pubsub-horizon-go/enum"
	"github.com/telekom/pubsub-horizon-go/types"
)

type CircuitBreakerMessage struct {
	SubscriptionId  string                    `json:"subscriptionId"`
	Environment     string                    `json:"environment"`
	EventType       string                    `json:"eventType"`
	LastModified    *types.Timestamp          `json:"lastModified"`
	LastOpened      *types.Timestamp          `json:"lastOpened,omitempty"`
	LoopCounter     int                       `json:"loopCounter"`
	Notified        bool                      `json:"notified"`
	OriginMessageId string                    `json:"originMessageId"`
	Status          enum.CircuitBreakerStatus `json:"status"`
}
