// Copyright 2024 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package enum

import (
	"fmt"
	"strconv"
	"strings"
)

type CircuitBreakerStatus string

const (
	CircuitBreakerStatusOpen         = "OPEN"
	CircuitBreakerStatusRepublishing = "REPUBLISHING"
	CircuitBreakerStatusCooldown     = "COOLDOWN"
	CircuitBreakerStatusChecking     = "CHECKING"
)

func ParseCircuitBreakerStatus(s string) (CircuitBreakerStatus, error) {
	switch CircuitBreakerStatus(s) {

	case CircuitBreakerStatusOpen, CircuitBreakerStatusRepublishing, CircuitBreakerStatusCooldown, CircuitBreakerStatusChecking:
		return CircuitBreakerStatus(s), nil

	default:
		return "", fmt.Errorf("could not parse '%s' as circuitBreakerStatus", s)

	}
}

func (cbStatus *CircuitBreakerStatus) UnmarshalJSON(bytes []byte) error {
	var data = string(bytes)

	if data == "null" {
		return nil
	}

	if strings.HasPrefix(data, `"`) && strings.HasSuffix(data, `"`) {
		data, _ = strconv.Unquote(data)
	}

	circuitBreakerStatus, err := ParseCircuitBreakerStatus(data)
	if err != nil {
		return err
	}

	*cbStatus = circuitBreakerStatus
	return nil
}

func (cbStatus *CircuitBreakerStatus) MarshalJSON() ([]byte, error) {
	var s = fmt.Sprintf(`"%s"`, cbStatus.String())
	return []byte(s), nil
}

func (cbStatus *CircuitBreakerStatus) String() string {
	return string(*cbStatus)
}
