// Copyright 2025 Deutsche Telekom AG
//
// SPDX-License-Identifier: Apache-2.0

package enum

import "strconv"

// UnmarshalEnum is a generic helper function for unmarshaling JSON into enum types.
// It handles null values, string unquoting, and delegates to a type-specific parse function.
//
// Parameters:
//   - bytes: The JSON bytes to unmarshal
//   - target: Pointer to the enum value to populate
//   - parseFunc: Function that parses a string into the enum type
//
// Returns an error if parsing fails.
func UnmarshalEnum[T any](bytes []byte, target *T, parseFunc func(string) (T, error)) error {
	data := string(bytes)

	// Handle JSON null values
	if data == "null" {
		return nil
	}

	// Remove JSON quotes if present
	if len(data) >= 2 && data[0] == '"' && data[len(data)-1] == '"' {
		var err error
		data, err = strconv.Unquote(data)
		if err != nil {
			return err
		}
	}

	// Parse the string using the provided parse function
	parsed, err := parseFunc(data)
	if err != nil {
		return err
	}

	*target = parsed
	return nil
}
