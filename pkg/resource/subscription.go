package resource

import "eni.telekom.de/horizon2go/pkg/constant"

type SubscriptionResource struct {
	Spec struct {
		Subscription Subscription `json:"subscription"`
		Environment  string       `json:"environment"`
	} `json:"spec"`
}

type Subscription struct {
	SubscriptionId         string                `json:"subscriptionId"`
	SubscriberId           string                `json:"subscriberId"`
	PublisherId            string                `json:"publisherId"`
	AdditionalPublisherIds []string              `json:"additionalPublisherIds"`
	CreatedAt              string                `json:"createdAt"`
	Trigger                SubscriptionTrigger   `json:"trigger"`
	PublisherTrigger       SubscriptionTrigger   `json:"publisherTrigger"`
	AppliedScopes          []string              `json:"appliedScopes"`
	Type                   string                `json:"type"`
	Callback               string                `json:"callback"`
	PayloadType            string                `json:"payloadType"`
	DeliveryType           constant.DeliveryType `json:"deliveryType"`
	EnforceGetHealthCheck  bool                  `json:"enforceGetHttpRequestMethodForHealthCheck"`
	CircuitBreakerOptOut   bool                  `json:"circuitBreakerOptOut"`
	RetentionTime          string                `json:"eventRetentionTime"`
}

type SubscriptionTrigger struct {
	ResponseFilterMode      constant.ResponseFilterMode `json:"responseFilterMode"`
	ResponseFilter          []string                    `json:"responseFilter"`
	SelectionFilter         map[string]string           `json:"selectionFilter"`
	AdvancedSelectionFilter map[string]any              `json:"advancedSelectionFilter"`
}
