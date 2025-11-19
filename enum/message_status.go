// Copyright 2025 Deutsche Telekom AG
//
// SPDX-License-Identifier: Apache-2.0

package enum

import (
	"errors"
	"fmt"
)

type MessageStatus string

const (
	StatusProcessed  MessageStatus = "PROCESSED"
	StatusDelivering MessageStatus = "DELIVERING"
	StatusWaiting    MessageStatus = "WAITING"
	StatusDelivered  MessageStatus = "DELIVERED"
	StatusFailed     MessageStatus = "FAILED"
	StatusDropped    MessageStatus = "DROPPED"
	StatusDuplicate  MessageStatus = "DUPLICATE"
)

func ParseMessageStatus(status string) (MessageStatus, error) {
	switch MessageStatus(status) {
	case StatusProcessed, StatusDelivering, StatusWaiting, StatusDelivered, StatusFailed, StatusDropped, StatusDuplicate:
		return MessageStatus(status), nil

	default:
		return "", errors.New("invalid message status")
	}
}

func (ms *MessageStatus) UnmarshalJSON(bytes []byte) error {
	return UnmarshalEnum(bytes, ms, ParseMessageStatus)
}

func (ms *MessageStatus) MarshalJSON() ([]byte, error) {
	s := fmt.Sprintf(`"%s"`, ms.String())
	return []byte(s), nil
}

func (ms *MessageStatus) String() string {
	return string(*ms)
}
