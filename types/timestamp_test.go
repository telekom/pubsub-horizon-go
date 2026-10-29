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

	timestamp := Timestamp(dummyTime)
	bytes, err := timestamp.MarshalJSON()
	assertions.NoError(err)
	assertions.Equal(`"0001-01-01T00:00:00Z"`, string(bytes))
}

func TestTimestampJSONRoundTrip(t *testing.T) {
	for _, timestampText := range []string{
		"0001-01-01T00:00:00Z",
		"2026-10-29T12:34:56Z",
		"2026-10-29T12:34:56+02:00",
		"2026-10-29T12:34:56.123456789Z",
		"2026-10-29T12:34:56.123456789-05:30",
	} {
		t.Run(timestampText, func(t *testing.T) {
			parsed, err := time.Parse(time.RFC3339Nano, timestampText)
			if err != nil {
				t.Fatal(err)
			}
			value := Timestamp(parsed)
			expected := `"` + timestampText + `"`
			for _, testCase := range []struct {
				name     string
				value    any
				expected string
			}{
				{"value", value, expected},
				{"pointer", &value, expected},
				{"struct field", struct {
					Time Timestamp `json:"time"`
				}{value}, `{"time":` + expected + `}`},
				{"map value", map[string]Timestamp{"time": value}, `{"time":` + expected + `}`},
				{"slice value", []Timestamp{value}, `[` + expected + `]`},
			} {
				t.Run(testCase.name, func(t *testing.T) {
					encoded, err := json.Marshal(testCase.value)
					assert.NoError(t, err)
					assert.Equal(t, testCase.expected, string(encoded))
				})
			}
			encoded, err := json.Marshal(&value)
			if err != nil {
				t.Fatal(err)
			}
			var restored Timestamp
			assert.NoError(t, json.Unmarshal(encoded, &restored))
			assert.True(t, parsed.Equal(restored.ToTime()))
		})
	}
}

func TestTimestampJSONNilPointer(t *testing.T) {
	var value *Timestamp
	encoded, err := json.Marshal(value)
	assert.NoError(t, err)
	assert.Equal(t, "null", string(encoded))
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
