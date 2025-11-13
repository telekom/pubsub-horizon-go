// Copyright 2025 Deutsche Telekom AG
//
// SPDX-License-Identifier: Apache-2.0

package message

import (
	"time"

	"github.com/telekom/pubsub-horizon-go/enum"
)

type StatusMessage struct {
	StateError
	Uuid                     string                  `json:"uuid"                     bson:"_id"`
	Coordinates              *Coordinates            `json:"coordinates"              bson:"coordinates"`
	Status                   enum.MessageStatus      `json:"status"                   bson:"status"`
	Environment              string                  `json:"environment"              bson:"environment"`
	DeliveryType             enum.DeliveryType       `json:"deliveryType"             bson:"deliveryType"`
	SubscriptionId           string                  `json:"subscriptionId"           bson:"subscriptionId"`
	Event                    EventDetails            `json:"event"                    bson:"event"`
	Properties               map[string]any          `json:"properties"               bson:"properties"`
	MultiplexedFrom          string                  `json:"multiplexedFrom"          bson:"multiplexedFrom"`
	EventRetentionTime       enum.EventRetentionTime `json:"eventRetentionTime"       bson:"eventRetentionTime"`
	Topic                    string                  `json:"topic"                    bson:"topic"`
	Timestamp                time.Time               `json:"timestamp"                bson:"timestamp"`
	Modified                 time.Time               `json:"modified"                 bson:"modified"`
	AppliedScopes            []string                `json:"appliedScopes"            bson:"appliedScopes"`
	ScopeEvaluationResult    EvaluationResult        `json:"scopeEvaluationResult"    bson:"scopeEvaluationResult"`
	ConsumerEvaluationResult EvaluationResult        `json:"consumerEvaluationResult" bson:"consumerEvaluationResult"`
}

type Coordinates struct {
	Partition *int32 `json:"partition" bson:"partition"`
	Offset    *int64 `json:"offset"    bson:"offset"`
}

type EventDetails struct {
	Id   string    `json:"id"   bson:"id"`
	Type string    `json:"type" bson:"type"`
	Time time.Time `json:"time" bson:"time"`
}

type StateError struct {
	ErrorMessage string `json:"errorMessage,omitempty" bson:"errorMessage,omitempty"`
	ErrorType    string `json:"errorType,omitempty"    bson:"errorType,omitempty"`
}

type EvaluationResult struct {
	OperatorName     string             `json:"operatorName"     bson:"operatorName"`
	Match            bool               `json:"match"            bson:"match"`
	CauseDescription string             `json:"causeDescription" bson:"causeDescription"`
	ChildOperators   []EvaluationResult `json:"childOperators"   bson:"childOperators"`
}
