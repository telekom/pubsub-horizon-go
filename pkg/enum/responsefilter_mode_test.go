package enum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseResponseFilterMode(t *testing.T) {
	var inputs = []struct {
		Value       string
		Expected    ResponseFilterMode
		ExpectError bool
	}{
		{"INCLUDE", ResponseFilterModeInclude, false},
		{"EXCLUDE", ResponseFilterModeExclude, false},
		{"INVALID", "", true},
	}

	for _, input := range inputs {
		t.Run(input.Value, func(t *testing.T) {
			var assertions = assert.New(t)

			rpm, err := ParseResponseFilterMode(input.Value)
			assertions.Equal(input.Expected, rpm)
			assertions.Equal(input.ExpectError, err != nil)
		})
	}
}

func TestResponseFilterMode_UnmarshalJSON(t *testing.T) {
	var inputs = []struct {
		Value       string
		Expected    ResponseFilterMode
		ExpectError bool
	}{
		{"INCLUDE", ResponseFilterModeInclude, false},
		{"EXCLUDE", ResponseFilterModeExclude, false},
		{"INVALID", "", true},
	}

	for _, input := range inputs {
		t.Run(input.Value, func(t *testing.T) {
			var assertions = assert.New(t)
			var rfm = new(ResponseFilterMode)

			var err = rfm.UnmarshalJSON([]byte(input.Value))
			if input.ExpectError {
				assertions.Error(err)
			} else {
				assertions.Equal(input.Expected, *rfm)
			}
		})
	}
}
