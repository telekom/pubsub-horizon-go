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
		{"TTL_1_HOUR", Ttl1Hour, false},
		{"TTL_1_DAY", Ttl1Day, false},
		{"TTL_3_DAYS", Ttl3Days, false},
		{"TTL_5_DAYS", Ttl5Days, false},
		{"TTL_7_DAYS", Ttl7Days, false},
		{"DEFAULT", TtlDefault, false},
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
		{"TTL_1_HOUR", Ttl1Hour, false},
		{"TTL_1_DAY", Ttl1Day, false},
		{"TTL_3_DAYS", Ttl3Days, false},
		{"TTL_5_DAYS", Ttl5Days, false},
		{"TTL_7_DAYS", Ttl7Days, false},
		{"DEFAULT", TtlDefault, false},
		{"8d", TTL{}, true},
	}

	for _, input := range inputs {
		t.Run(input.Value, func(t *testing.T) {
			assertions := assert.New(t)
			eventRetentionTime := new(TTL)

			err := eventRetentionTime.UnmarshalJSON([]byte(`"` + input.Value + `"`))
			if input.ExpectError {
				assertions.Error(err, "Expected error for input: %s", input.Value)
			} else {
				assertions.NoError(err, "Unexpected error for input: %s", input.Value)
				assertions.Equal(input.Expected, *eventRetentionTime, "Unexpected TTL value for input: %s", input.Value)
			}
		})
	}
}

func TestEventRetentionTime_MarshalJSON(t *testing.T) {
	var inputs = []struct {
		Value    TTL
		Expected string
	}{
		{Ttl1Hour, `"TTL_1_HOUR"`},
		{Ttl1Day, `"TTL_1_DAY"`},
		{Ttl3Days, `"TTL_3_DAYS"`},
		{Ttl5Days, `"TTL_5_DAYS"`},
		{Ttl7Days, `"TTL_7_DAYS"`},
		{TtlDefault, `"DEFAULT"`},
	}

	for _, input := range inputs {
		t.Run(input.Value.Topic, func(t *testing.T) {
			var assertions = assert.New(t)
			marshalled, err := input.Value.MarshalJSON()
			assertions.NoError(err)
			assertions.Equal(input.Expected, string(marshalled))
		})
	}
}
