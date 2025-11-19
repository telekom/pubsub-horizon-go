// Copyright 2025 Deutsche Telekom AG
//
// SPDX-License-Identifier: Apache-2.0

package enum

import (
	"encoding/json"
	"fmt"
)

type EventRetentionTime string

const (
	Ttl7Days   EventRetentionTime = "TTL_7_DAYS"
	Ttl5Days   EventRetentionTime = "TTL_5_DAYS"
	Ttl3Days   EventRetentionTime = "TTL_3_DAYS"
	Ttl1Day    EventRetentionTime = "TTL_1_DAY"
	Ttl1Hour   EventRetentionTime = "TTL_1_HOUR"
	TtlDefault EventRetentionTime = "DEFAULT"
)

var EventRetentionTimes = map[EventRetentionTime]struct {
	Topic         string
	RetentionInMs int64
}{
	Ttl7Days:   {"subscribed", 604800000},
	Ttl5Days:   {"subscribed_5d", 432000000},
	Ttl3Days:   {"subscribed_3d", 259200000},
	Ttl1Day:    {"subscribed_1d", 86400000},
	Ttl1Hour:   {"subscribed_1h", 3600000},
	TtlDefault: {"subscribed", 604800000},
}

func ParseEventRetentionTime(s string) (EventRetentionTime, error) {
	for key := range EventRetentionTimes {
		if string(key) == s {
			return key, nil
		}
	}
	return TtlDefault, fmt.Errorf("could not parse '%s' as EventRetentionTime", s)
}

func (ttl *EventRetentionTime) UnmarshalJSON(bytes []byte) error {
	return UnmarshalEnum(bytes, ttl, ParseEventRetentionTime)
}

func (ttl *EventRetentionTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(*ttl)
}

func (ttl *EventRetentionTime) ToRoverConfigString() string {
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
	case TtlDefault:
		return "7d"
	default:
		return "7d"
	}
}
