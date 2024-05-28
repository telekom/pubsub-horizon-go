package enum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseDeliveryType(t *testing.T) {
	var inputs = []struct {
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
			var assertions = assert.New(t)

			deliveryType, err := ParseDeliveryType(input.Value)
			assertions.Equal(input.Expected, deliveryType)
			assertions.Equal(input.ExpectError, err != nil)
		})
	}
}

func TestDeliveryType_UnmarshalJSON(t *testing.T) {
	var inputs = []struct {
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
			var assertions = assert.New(t)
			var deliveryType = new(DeliveryType)

			var err = deliveryType.UnmarshalJSON([]byte(input.Value))
			if input.ExpectError {
				assertions.Error(err)
			} else {
				assertions.Equal(input.Expected, *deliveryType)
			}
		})
	}
}
