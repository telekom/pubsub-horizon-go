// Copyright 2024 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package enum

import (
	"fmt"
	"strconv"
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
	var data = string(bytes)

	if data == "null" {
		return nil
	}

	if strings.HasPrefix(data, `"`) && strings.HasSuffix(data, `"`) {
		data, _ = strconv.Unquote(data)
	}

	rfm, err := ParseResponseFilterMode(data)
	if err != nil {
		return err
	}

	*m = rfm
	return nil
}

func (m *ResponseFilterMode) MarshalJSON() ([]byte, error) {
	var s = fmt.Sprintf(`"%s"`, m.String())
	return []byte(s), nil
}

func (m *ResponseFilterMode) String() string {
	return string(*m)
}
