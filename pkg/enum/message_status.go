package enum

import (
	"errors"
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
	var data = string(bytes)

	if data == "null" {
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
