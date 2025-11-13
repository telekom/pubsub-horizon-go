// Copyright 2025 Deutsche Telekom AG
//
// SPDX-License-Identifier: Apache-2.0

package enum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
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
	data := string(bytes)

	if data == jsonNull {
		return nil
	}

	if strings.HasPrefix(data, `"`) && strings.HasSuffix(data, `"`) {
		data, _ = strconv.Unquote(data)
	}

	messageStatus, err := ParseMessageStatus(data)
	if err != nil {
		return err
	}

	*ms = messageStatus
	return nil
}

func (ms *MessageStatus) MarshalJSON() ([]byte, error) {
	s := fmt.Sprintf(`"%s"`, ms.String())
	return []byte(s), nil
}

func (ms *MessageStatus) String() string {
	return string(*ms)
}
