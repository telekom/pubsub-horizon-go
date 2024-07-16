// Copyright 2024 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCircuitBreakerTime_MarshalJSON(t *testing.T) {
	var assertions = assert.New(t)
	var dummyTime = time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)

	bytes, err := dummyTime.MarshalJSON()
	assertions.NoError(err)
	assertions.Equal(`"0001-01-01T00:00:00Z"`, string(bytes))
}

func TestCircuitBreakerTime_UnmarshalJSON(t *testing.T) {
	var assertions = assert.New(t)
	var expectation = time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)
	var bytes = []byte(`{"time": "0001-01-01T00:00:00Z"}`)

	var dummy = struct {
		Timestamp Timestamp `json:"time"`
	}{}

	assertions.NoError(json.Unmarshal(bytes, &dummy))

	var unmarshalledTime = dummy.Timestamp.ToTime()
	assertions.True(unmarshalledTime.Equal(expectation))
}
