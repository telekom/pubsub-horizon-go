package enum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseEventRetentionTime(t *testing.T) {
	var inputs = []struct {
		Value       string
		Expected    TTL
		ExpectError bool
	}{
		{"subscribed_1h", Ttl1Hour, false},
		{"subscribed_1d", Ttl1Day, false},
		{"subscribed_3d", Ttl3Days, false},
		{"subscribed_5d", Ttl5Days, false},
		{"subscribed", Ttl7Days, false},
		{"subscribed", TtlDefault, false},
		{"8d", TTL{}, true},
	}

	for _, input := range inputs {
		t.Run(input.Value, func(t *testing.T) {
			var assertions = assert.New(t)

			eventRetentionTime, err := ParseEventRetentionTime(input.Value)
			assertions.Equal(input.Expected, eventRetentionTime)
			assertions.Equal(input.ExpectError, err != nil)
		})
	}
}

func TestRoverString(t *testing.T) {
	var inputs = []struct {
		Value       TTL
		Expected    string
		ExpectError bool
	}{
		{Ttl1Hour, "1h", false},
		{Ttl1Day, "1d", false},
		{Ttl3Days, "3d", false},
		{Ttl5Days, "5d", false},
		{Ttl7Days, "7d", false},
		{TtlDefault, "7d", false},
		{TTL{}, "7d", false},
	}

	for _, input := range inputs {
		t.Run(input.Value.Topic, func(t *testing.T) {
			var assertions = assert.New(t)

			eventRetentionTime := input.Value.RoverString()
			assertions.Equal(input.Expected, eventRetentionTime)
		})
	}
}

func TestEventRetentionTime_UnmarshalJSON(t *testing.T) {
	var inputs = []struct {
		Value       string
		Expected    TTL
		ExpectError bool
	}{
		{"subscribed_1h", Ttl1Hour, false},
		{"subscribed_1d", Ttl1Day, false},
		{"subscribed_3d", Ttl3Days, false},
		{"subscribed_5d", Ttl5Days, false},
		{"subscribed", Ttl7Days, false},
		{"subscribed", TtlDefault, false},
		{"8d", TTL{}, true},
	}

	for _, input := range inputs {
		t.Run(input.Value, func(t *testing.T) {
			var assertions = assert.New(t)
			var eventRetentionTime = new(TTL)

			var err = eventRetentionTime.UnmarshalJSON([]byte(input.Value))
			if input.ExpectError {
				assertions.Error(err)
			} else {
				assertions.Equal(input.Expected, *eventRetentionTime)
			}
		})
	}
}
