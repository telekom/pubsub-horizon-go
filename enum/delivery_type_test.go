// Copyright 2024 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package enum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseDeliveryType(t *testing.T) {
	inputs := []struct {
		Value       string
		Expected    DeliveryType
		ExpectError bool
	}{
		{"sse", DeliveryTypeSse, false},
		{"server_sent_event", DeliveryTypeSse, false},
		{"callback", DeliveryTypeCallback, false},
		{"invalid", "", true},
	}

	for _, input := range inputs {
		t.Run(input.Value, func(t *testing.T) {
			assertions := assert.New(t)

			deliveryType, err := ParseDeliveryType(input.Value)
			assertions.Equal(input.Expected, deliveryType)
			assertions.Equal(input.ExpectError, err != nil)
		})
	}
}

func TestDeliveryType_UnmarshalJSON(t *testing.T) {
	inputs := []struct {
		Value       string
		Expected    DeliveryType
		ExpectError bool
	}{
		{"sse", DeliveryTypeSse, false},
		{"server_sent_event", DeliveryTypeSse, false},
		{"callback", DeliveryTypeCallback, false},
		{"invalid", "", true},
	}

	for _, input := range inputs {
		t.Run(input.Value, func(t *testing.T) {
			assertions := assert.New(t)
			deliveryType := new(DeliveryType)

			err := deliveryType.UnmarshalJSON([]byte(input.Value))
			if input.ExpectError {
				assertions.Error(err)
			} else {
				assertions.Equal(input.Expected, *deliveryType)
			}
		})
	}
}

func TestDeliveryType_MarshalJSON(t *testing.T) {
	inputs := []struct {
		Value    DeliveryType
		Expected string
	}{
		{DeliveryTypeCallback, `"callback"`},
		{DeliveryTypeSse, `"server_sent_event"`},
	}

	for _, input := range inputs {
		t.Run(input.Value.String(), func(t *testing.T) {
			assertions := assert.New(t)
			marshalled, err := input.Value.MarshalJSON()
			assertions.NoError(err)
			assertions.Equal(input.Expected, string(marshalled))
		})
	}
}
