package enum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseMessageStatus(t *testing.T) {
	var inputs = []struct {
		Value         string
		Expected      MessageStatus
		ExpectedError bool
	}{
		{"PROCESSED", StatusProcessed, false},
		{"DELIVERING", StatusDelivering, false},
		{"WAITING", StatusWaiting, false},
		{"DELIVERED", StatusDelivered, false},
		{"FAILED", StatusFailed, false},
		{"DROPPED", StatusDropped, false},
		{"DUPLICATE", StatusDuplicate, false},
		{"INVALID", "", true},
	}

	for _, input := range inputs {
		t.Run(input.Value, func(t *testing.T) {
			var assertions = assert.New(t)
			msgStatus, err := ParseMessageStatus(input.Value)
			assertions.Equal(input.Expected, msgStatus)
			assertions.Equal(input.ExpectedError, err != nil)
		})
	}
}

func TestMessageStatus_UnmarshalJSON(t *testing.T) {
	var inputs = []struct {
		Value       string
		Expected    MessageStatus
		ExpectError bool
	}{
		{"PROCESSED", StatusProcessed, false},
		{"DELIVERING", StatusDelivering, false},
		{"WAITING", StatusWaiting, false},
		{"DELIVERED", StatusDelivered, false},
		{"FAILED", StatusFailed, false},
		{"DROPPED", StatusDropped, false},
		{"DUPLICATE", StatusDuplicate, false},
		{"INVALID", "", true},
	}

	for _, input := range inputs {
		t.Run(input.Value, func(t *testing.T) {
			var assertions = assert.New(t)
			var messageStatus = new(MessageStatus)

			var err = messageStatus.UnmarshalJSON([]byte(input.Value))
			if input.ExpectError {
				assertions.NotNil(err)
			} else {
				assertions.Equal(input.Expected, *messageStatus)
			}
		})
	}
}