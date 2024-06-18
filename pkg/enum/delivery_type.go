package enum

import (
	"fmt"
	"strconv"
	"strings"
)

type DeliveryType string

const (
	DeliveryTypeSse      = "server_sent_event"
	DeliveryTypeCallback = "callback"
)

func ParseDeliveryType(s string) (DeliveryType, error) {
	switch strings.ToLower(s) {

	case "sse":
		return DeliveryTypeSse, nil

	case "server_sent_event", "callback":
		return DeliveryType(s), nil

	default:
		return "", fmt.Errorf("could not parse '%s' as delivery type", s)

	}
}

func (t *DeliveryType) UnmarshalJSON(bytes []byte) error {
	var data = string(bytes)

	if data == "null" {
		return nil
	}

	if strings.HasPrefix(data, `"`) && strings.HasSuffix(data, `"`) {
		data, _ = strconv.Unquote(data)
	}

	deliveryType, err := ParseDeliveryType(data)
	if err != nil {
		return err
	}

	*t = deliveryType
	return nil
}

func (t *DeliveryType) MarshalJSON() ([]byte, error) {
	var s = fmt.Sprintf(`"%s"`, t.String())
	return []byte(s), nil
}

func (t *DeliveryType) String() string {
	return string(*t)
}
