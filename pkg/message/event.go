package message

import (
	"time"
)

type Event struct {
	Id              string    `json:"id"`
	Type            string    `json:"type"`
	Source          string    `json:"source"`
	SpecVersion     string    `json:"specVersion"`
	DataContentType string    `json:"dataContentType"`
	DataRef         string    `json:"dataRef"`
	Time            time.Time `json:"time"`
	Data            any       `json:"data"`
}
