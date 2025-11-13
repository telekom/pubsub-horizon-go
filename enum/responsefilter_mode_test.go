// Copyright 2024 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package enum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseResponseFilterMode(t *testing.T) {
	inputs := []struct {
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
			assertions := assert.New(t)

			rpm, err := ParseResponseFilterMode(input.Value)
			assertions.Equal(input.Expected, rpm)
			assertions.Equal(input.ExpectError, err != nil)
		})
	}
}

func TestResponseFilterMode_UnmarshalJSON(t *testing.T) {
	inputs := []struct {
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
			assertions := assert.New(t)
			rfm := new(ResponseFilterMode)

			err := rfm.UnmarshalJSON([]byte(input.Value))
			if input.ExpectError {
				assertions.Error(err)
			} else {
				assertions.Equal(input.Expected, *rfm)
			}
		})
	}
}

func TestResponseFilterMode_MarshalJSON(t *testing.T) {
	inputs := []struct {
		Value    ResponseFilterMode
		Expected string
	}{
		{ResponseFilterModeInclude, `"INCLUDE"`},
		{ResponseFilterModeExclude, `"EXCLUDE"`},
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
