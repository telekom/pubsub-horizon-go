// Copyright 2024 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package enum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseCircuitBreakerStatus(t *testing.T) {
	var inputs = []struct {
		Value       string
		Expected    CircuitBreakerStatus
		ExpectError bool
	}{
		{"OPEN", CircuitBreakerStatusOpen, false},
		{"REPUBLISHING", CircuitBreakerStatusRepublishing, false},
		{"COOLDOWN", CircuitBreakerStatusCooldown, false},
		{"CHECKING", CircuitBreakerStatusChecking, false},
		{"INVALID", "", true},
	}

	for _, input := range inputs {
		t.Run(input.Value, func(t *testing.T) {
			var assertions = assert.New(t)

			circuitBreakerStatus, err := ParseCircuitBreakerStatus(input.Value)
			assertions.Equal(input.Expected, circuitBreakerStatus)
			assertions.Equal(input.ExpectError, err != nil)
		})
	}
}

func TestCircuitBreakerStatus_UnmarshalJSON(t *testing.T) {
	var inputs = []struct {
		Value       string
		Expected    CircuitBreakerStatus
		ExpectError bool
	}{
		{"OPEN", CircuitBreakerStatusOpen, false},
		{"REPUBLISHING", CircuitBreakerStatusRepublishing, false},
		{"COOLDOWN", CircuitBreakerStatusCooldown, false},
		{"CHECKING", CircuitBreakerStatusChecking, false},
		{"INVALID", "", true},
	}

	for _, input := range inputs {
		t.Run(input.Value, func(t *testing.T) {
			var assertions = assert.New(t)
			var cbStatus = new(CircuitBreakerStatus)

			var err = cbStatus.UnmarshalJSON([]byte(input.Value))
			if input.ExpectError {
				assertions.Error(err)
			} else {
				assertions.Equal(input.Expected, *cbStatus)
			}
		})
	}
}

func TestCircuitBreakerStatus_MarshalJSON(t *testing.T) {
	var inputs = []struct {
		Value    CircuitBreakerStatus
		Expected string
	}{
		{CircuitBreakerStatusOpen, `"OPEN"`},
		{CircuitBreakerStatusRepublishing, `"REPUBLISHING"`},
		{CircuitBreakerStatusCooldown, `"COOLDOWN"`},
		{CircuitBreakerStatusChecking, `"CHECKING"`},
	}

	for _, input := range inputs {
		t.Run(input.Value.String(), func(t *testing.T) {
			var assertions = assert.New(t)
			marshalled, err := input.Value.MarshalJSON()
			assertions.NoError(err)
			assertions.Equal(input.Expected, string(marshalled))
		})
	}
}
