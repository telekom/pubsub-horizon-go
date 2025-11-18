// Copyright 2025 Deutsche Telekom AG
//
// SPDX-License-Identifier: Apache-2.0

package enum

import (
	"fmt"
)

type CircuitBreakerStatus string

const (
	CircuitBreakerStatusOpen   CircuitBreakerStatus = "OPEN"
	CircuitBreakerStatusClosed CircuitBreakerStatus = "CLOSED"
)

func ParseCircuitBreakerStatus(s string) (CircuitBreakerStatus, error) {
	switch CircuitBreakerStatus(s) {
	case CircuitBreakerStatusOpen, CircuitBreakerStatusClosed:
		return CircuitBreakerStatus(s), nil

	default:
		return "", fmt.Errorf("could not parse '%s' as circuitBreakerStatus", s)
	}
}

func (cbStatus *CircuitBreakerStatus) UnmarshalJSON(bytes []byte) error {
	return UnmarshalEnum(bytes, cbStatus, ParseCircuitBreakerStatus)
}

func (cbStatus *CircuitBreakerStatus) MarshalJSON() ([]byte, error) {
	s := fmt.Sprintf(`"%s"`, cbStatus.String())
	return []byte(s), nil
}

func (cbStatus *CircuitBreakerStatus) String() string {
	return string(*cbStatus)
}
