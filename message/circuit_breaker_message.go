// Copyright 2024 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package message

import (
	"github.com/telekom/pubsub-horizon-go/enum"
	"time"
)

type CircuitBreakerMessage struct {
	SubscriptionId    string                    `json:"subscriptionId"`
	Environment       string                    `json:"environment"`
	EventType         string                    `json:"eventType"`
	LastModified      time.Time                 `json:"lastModified"`
	OriginMessageId   string                    `json:"originMessageId"`
	Status            enum.CircuitBreakerStatus `json:"status"`
	LastRepublished   time.Time                 `json:"lastRepublished"`
	RepublishingCount int                       `json:"republishingCount"`
}
