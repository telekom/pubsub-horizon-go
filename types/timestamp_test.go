// Copyright 2025 Deutsche Telekom AG
//
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCircuitBreakerTime_MarshalJSON(t *testing.T) {
	assertions := assert.New(t)
	dummyTime := time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)

	bytes, err := dummyTime.MarshalJSON()
	assertions.NoError(err)
	assertions.Equal(`"0001-01-01T00:00:00Z"`, string(bytes))
}

func TestCircuitBreakerTime_UnmarshalJSON(t *testing.T) {
	assertions := assert.New(t)
	expectation := time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)
	bytes := []byte(`{"time": "0001-01-01T00:00:00Z"}`)

	dummy := struct {
		Timestamp Timestamp `json:"time"`
	}{}

	assertions.NoError(json.Unmarshal(bytes, &dummy))

	unmarshalledTime := dummy.Timestamp.ToTime()
	assertions.True(unmarshalledTime.Equal(expectation))
}
