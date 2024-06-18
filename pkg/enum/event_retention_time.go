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
	Ttl7Days   = TTL{"subscribed", 604800000}
	Ttl5Days   = TTL{"subscribed_5d", 432000000}
	Ttl3Days   = TTL{"subscribed_3d", 259200000}
	Ttl1Day    = TTL{"subscribed_1d", 86400000}
	Ttl1Hour   = TTL{"subscribed_1h", 3600000}
	TtlDefault = Ttl7Days
)

func ParseEventRetentionTime(s string) (TTL, error) {
	switch strings.ToLower(s) {

	case "subscribed":
		return Ttl7Days, nil
	case "subscribed_5d":
		return Ttl5Days, nil
	case "subscribed_3d":
		return Ttl3Days, nil
	case "subscribed_1d":
		return Ttl1Day, nil
	case "subscribed_1h":
		return Ttl1Hour, nil
	case "default":
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
