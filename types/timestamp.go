package types

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Timestamp time.Time

func NewTimestamp(time time.Time) *Timestamp {
	var timestamp = Timestamp(time)
	return &timestamp
}

func (c *Timestamp) MarshalJSON() ([]byte, error) {
	var current = time.Time(*c)

	var s = fmt.Sprintf(`"%s"`, current.Format(time.RFC3339))
	return []byte(s), nil
}

func (c *Timestamp) UnmarshalJSON(bytes []byte) error {
	var data = string(bytes)
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
