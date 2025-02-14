// Copyright 2024 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package resource

import "github.com/telekom/pubsub-horizon-go/enum"

type SubscriptionResource struct {
	Metadata Metadata `json:"metadata"`
	Spec     struct {
		Subscription Subscription `json:"subscription"`
		Environment  string       `json:"environment"`
	} `json:"spec"`
}

type Metadata struct {
	Annotations map[string]any `json:"annotations"`
}

type Subscription struct {
	SubscriptionId         string              `json:"subscriptionId"`
	SubscriberId           string              `json:"subscriberId"`
	PublisherId            string              `json:"publisherId"`
	AdditionalPublisherIds []string            `json:"additionalPublisherIds"`
	CreatedAt              string              `json:"createdAt"`
	Trigger                SubscriptionTrigger `json:"trigger"`
	PublisherTrigger       SubscriptionTrigger `json:"publisherTrigger"`
	AppliedScopes          []string            `json:"appliedScopes"`
	Type                   string              `json:"type"`
	Callback               string              `json:"callback"`
	PayloadType            string              `json:"payloadType"`
	DeliveryType           enum.DeliveryType   `json:"deliveryType"`
	EnforceGetHealthCheck  bool                `json:"enforceGetHttpRequestMethodForHealthCheck"`
	CircuitBreakerOptOut   bool                `json:"circuitBreakerOptOut"`
	RetentionTime          string              `json:"eventRetentionTime"`
	RedeliveriesPerSecond  int                 `json:"redeliveriesPerSecond"`
}

type SubscriptionTrigger struct {
	ResponseFilterMode      enum.ResponseFilterMode `json:"responseFilterMode"`
	ResponseFilter          []string                `json:"responseFilter"`
	SelectionFilter         map[string]string       `json:"selectionFilter"`
	AdvancedSelectionFilter map[string]any          `json:"advancedSelectionFilter"`
}
