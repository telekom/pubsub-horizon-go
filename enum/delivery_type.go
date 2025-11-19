// Copyright 2025 Deutsche Telekom AG
//
// SPDX-License-Identifier: Apache-2.0

package enum

import (
	"fmt"
	"strings"
)

type DeliveryType string

const (
	DeliveryTypeSse      DeliveryType = "server_sent_event"
	DeliveryTypeCallback DeliveryType = "callback"
)

func ParseDeliveryType(s string) (DeliveryType, error) {
	switch strings.ToLower(s) {
	case "sse":
		return DeliveryTypeSse, nil

	case "server_sent_event", "callback":
		return DeliveryType(s), nil

	default:
		return "", fmt.Errorf("could not parse '%s' as delivery type", s)
	}
}

func (t *DeliveryType) UnmarshalJSON(bytes []byte) error {
	return UnmarshalEnum(bytes, t, ParseDeliveryType)
}

func (t *DeliveryType) MarshalJSON() ([]byte, error) {
	s := fmt.Sprintf(`"%s"`, t.String())
	return []byte(s), nil
}

func (t *DeliveryType) String() string {
	return string(*t)
}
