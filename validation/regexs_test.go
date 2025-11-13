// Copyright 2024 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEventTypeRegEx(t *testing.T) {
	inputs := []struct {
		Value    string
		Expected bool
	}{
		{"my.valid.event.type.v1", true},
		{"my.invalid.event.type.v1@", false},
	}

	for _, input := range inputs {
		assertions := assert.New(t)
		t.Run(input.Value, func(t *testing.T) {
			assertions.Equal(input.Expected, EventTypeRegEx.MatchString(input.Value))
		})
	}
}

func TestIso8601RegEx(t *testing.T) {
	inputs := []struct {
		Value    string
		Expected bool
	}{
		{"2024-01-01T00:00:00.000Z", true},
		{"2024-01-01T00:00Z", false},
	}

	for _, input := range inputs {
		assertions := assert.New(t)
		t.Run(input.Value, func(t *testing.T) {
			assertions.Equal(input.Expected, Iso8601RegEx.MatchString(input.Value))
		})
	}
}
