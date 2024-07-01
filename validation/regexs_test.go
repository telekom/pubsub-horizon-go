package validation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEventTypeRegEx(t *testing.T) {
	var inputs = []struct {
		Value    string
		Expected bool
	}{
		{"my.valid.event.type.v1", true},
		{"my.invalid.event.type.v1@", false},
	}

	for _, input := range inputs {
		var assertions = assert.New(t)
		t.Run(input.Value, func(t *testing.T) {
			assertions.Equal(input.Expected, EventTypeRegEx.MatchString(input.Value))
		})
	}
}

func TestIso8601RegEx(t *testing.T) {
	var inputs = []struct {
		Value    string
		Expected bool
	}{
		{"2024-01-01T00:00:00.000Z", true},
		{"2024-01-01T00:00Z", false},
	}

	for _, input := range inputs {
		var assertions = assert.New(t)
		t.Run(input.Value, func(t *testing.T) {
			assertions.Equal(input.Expected, Iso8601RegEx.MatchString(input.Value))
		})
	}
}
