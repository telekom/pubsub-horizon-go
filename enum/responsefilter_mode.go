// Copyright 2025 Deutsche Telekom AG
//
// SPDX-License-Identifier: Apache-2.0

package enum

import (
	"fmt"
	"strings"
)

type ResponseFilterMode string

const (
	ResponseFilterModeInclude ResponseFilterMode = "INCLUDE"
	ResponseFilterModeExclude ResponseFilterMode = "EXCLUDE"
)

func ParseResponseFilterMode(s string) (ResponseFilterMode, error) {
	switch strings.ToLower(s) {
	case "include", "exclude":
		return ResponseFilterMode(s), nil

	default:
		return "", fmt.Errorf("could not parse '%s' as response-filter-mode", s)
	}
}

func (m *ResponseFilterMode) UnmarshalJSON(bytes []byte) error {
	return UnmarshalEnum(bytes, m, ParseResponseFilterMode)
}

func (m *ResponseFilterMode) MarshalJSON() ([]byte, error) {
	s := fmt.Sprintf(`"%s"`, m.String())
	return []byte(s), nil
}

func (m *ResponseFilterMode) String() string {
	return string(*m)
}
