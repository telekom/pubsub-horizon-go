package enum

import (
	"fmt"
	"strings"
)

type ResponseFilterMode string

const (
	ResponseFilterModeInclude = "INCLUDE"
	ResponseFilterModeExclude = "EXCLUDE"
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

	rfm, err := ParseResponseFilterMode(data)
	if err != nil {
		return err
	}

	*m = rfm
	return nil
}

func (m *ResponseFilterMode) String() string {
	return string(*m)
}
