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
