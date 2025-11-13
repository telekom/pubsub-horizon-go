// Copyright 2025 Deutsche Telekom AG
//
// SPDX-License-Identifier: Apache-2.0

package enum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCircuitBreakerStatus(t *testing.T) {
	inputs := []struct {
		Value       string
		Expected    CircuitBreakerStatus
		ExpectError bool
	}{
		{"OPEN", CircuitBreakerStatusOpen, false},
		{"CLOSED", CircuitBreakerStatusClosed, false},
		{"INVALID", "", true},
	}

	for _, input := range inputs {
		t.Run(input.Value, func(t *testing.T) {
			assertions := assert.New(t)

			circuitBreakerStatus, err := ParseCircuitBreakerStatus(input.Value)
			assertions.Equal(input.Expected, circuitBreakerStatus)
			assertions.Equal(input.ExpectError, err != nil)
		})
	}
}

func TestCircuitBreakerStatus_UnmarshalJSON(t *testing.T) {
	inputs := []struct {
		Value       string
		Expected    CircuitBreakerStatus
		ExpectError bool
	}{
		{"OPEN", CircuitBreakerStatusOpen, false},
		{"CLOSED", CircuitBreakerStatusClosed, false},
		{"INVALID", "", true},
	}

	for _, input := range inputs {
		t.Run(input.Value, func(t *testing.T) {
			assertions := assert.New(t)
			cbStatus := new(CircuitBreakerStatus)

			err := cbStatus.UnmarshalJSON([]byte(input.Value))
			if input.ExpectError {
				assertions.Error(err)
			} else {
				assertions.Equal(input.Expected, *cbStatus)
			}
		})
	}
}

func TestCircuitBreakerStatus_MarshalJSON(t *testing.T) {
	inputs := []struct {
		Value    CircuitBreakerStatus
		Expected string
	}{
		{CircuitBreakerStatusOpen, `"OPEN"`},
		{CircuitBreakerStatusClosed, `"CLOSED"`},
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
