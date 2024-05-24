package message

import "eni.telekom.de/horizon2go/pkg/enum"

type PublishedMessage struct {
	Uuid             string              `json:"uuid"`
	Environment      string              `json:"environment"`
	AdditionalFields map[string]any      `json:"additionalFields"`
	Event            Event               `json:"event"`
	Status           enum.MessageStatus  `json:"status"`
	HttpHeaders      map[string][]string `json:"httpHeaders"`
}
