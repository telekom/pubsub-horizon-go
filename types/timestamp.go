// Copyright 2025 Deutsche Telekom AG
//
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"strconv"
	"strings"
	"time"
)

type Timestamp time.Time

func (c Timestamp) MarshalJSON() ([]byte, error) {
	return time.Time(c).MarshalJSON()
}

func (c *Timestamp) UnmarshalJSON(bytes []byte) error {
	data := string(bytes)
	var parsedTime time.Time
	var err error

	if data == "null" {
		parsedTime = time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)
		*c = Timestamp(parsedTime)
		return nil
	}

	if strings.HasPrefix(data, `"`) && strings.HasSuffix(data, `"`) {
		data, _ = strconv.Unquote(data)
	}

	parsedTime, err = time.Parse(time.RFC3339, data)
	if err != nil {
		return err
	}

	*c = Timestamp(parsedTime)
	return nil
}

func (c *Timestamp) ToTime() time.Time {
	return time.Time(*c)
}
