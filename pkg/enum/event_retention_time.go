package enum

import (
	"fmt"
	"strconv"
	"strings"
)

type TTL struct {
	Topic             string
	RetentionTimeInMs int
}

var (
	Ttl7Days   = TTL{"TTL_7_DAYS", 604800000}
	Ttl5Days   = TTL{"TTL_5_DAYS", 432000000}
	Ttl3Days   = TTL{"TTL_3_DAYS", 259200000}
	Ttl1Day    = TTL{"TTL_1_DAY", 86400000}
	Ttl1Hour   = TTL{"TTL_1_HOUR", 3600000}
	TtlDefault = TTL{"DEFAULT", 604800000}
)

func ParseEventRetentionTime(s string) (TTL, error) {
	switch s {

	case "TTL_7_DAYS":
		return Ttl7Days, nil
	case "TTL_5_DAYS":
		return Ttl5Days, nil
	case "TTL_3_DAYS":
		return Ttl3Days, nil
	case "TTL_1_DAY":
		return Ttl1Day, nil
	case "TTL_1_HOUR":
		return Ttl1Hour, nil
	case "DEFAULT":
		return TtlDefault, nil
	default:
		return TTL{}, fmt.Errorf("could not parse '%s' as eventRetentionTime", s)
	}
}

func (ttl *TTL) RoverString() string {
	switch *ttl {
	case Ttl7Days:
		return "7d"
	case Ttl5Days:
		return "5d"
	case Ttl3Days:
		return "3d"
	case Ttl1Day:
		return "1d"
	case Ttl1Hour:
		return "1h"
	default:
		return "7d"
	}
}

func (ttl *TTL) UnmarshalJSON(bytes []byte) error {
	var data = string(bytes)

	if data == "null" {
		return nil
	}

	if strings.HasPrefix(data, `"`) && strings.HasSuffix(data, `"`) {
		data, _ = strconv.Unquote(data)
	}

	eventRetentionTime, err := ParseEventRetentionTime(data)
	if err != nil {
		return err
	}

	*ttl = eventRetentionTime
	return nil
}

func (ttl *TTL) MarshalJSON() ([]byte, error) {
	var s = fmt.Sprintf(`"%s"`, ttl.Topic)
	return []byte(s), nil
}
