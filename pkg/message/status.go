package message

import (
	"eni.telekom.de/horizon2go/pkg/enum"
	"time"
)

type StatusMessage struct {
	Uuid                     string             `json:"uuid"`
	Coordinates              *Coordinates       `json:"coordinates"`
	Status                   enum.MessageStatus `json:"status"`
	Environment              string             `json:"environment"`
	DeliveryType             enum.DeliveryType  `json:"deliveryType"`
	SubscriptionId           string             `json:"subscriptionId"`
	Event                    EventDetails       `json:"event"`
	Properties               map[string]string  `json:"properties"`
	MultiplexedFrom          string             `json:"multiplexedFrom"`
	EventRetentionTime       enum.TTL           `json:"eventRetentionTime"`
	Topic                    string             `json:"topic"`
	Timestamp                time.Time          `json:"timestamp"`
	Modified                 time.Time          `json:"modified"`
	StateError               StateError         `json:"stateError"`
	AppliedScopes            []string           `json:"appliedScopes"`
	ScopeEvaluationResult    EvaluationResult   `json:"scopeEvaluationResult"`
	ConsumerEvaluationResult EvaluationResult   `json:"consumerEvaluationResult"`
}

type Coordinates struct {
	Partition *int32 `json:"partition"`
	Offset    *int64 `json:"offset"`
}

type EventDetails struct {
	Id   string    `json:"id"`
	Type string    `json:"type"`
	Time time.Time `json:"time"`
}

type StateError struct {
	Message    string `json:"message"`
	Type       string `json:"type"`
	StackTrace string `json:"stackTrace"`
}

type EvaluationResult struct {
	OperatorName     string             `json:"operatorName"`
	Match            bool               `json:"match"`
	CauseDescription string             `json:"causeDescription"`
	ChildOperators   []EvaluationResult `json:"childOperators"`
}
