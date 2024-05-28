package message

import (
	"time"
)

type Event struct {
	Id              string    `json:"id" validate:"required,uuid4"`
	Type            string    `json:"type" validate:"required,eventType"`
	Source          string    `json:"source" validate:"required"`
	SpecVersion     string    `json:"specVersion" validate:"required"`
	DataContentType string    `json:"dataContentType"`
	DataRef         string    `json:"dataRef"`
	Time            time.Time `json:"time" validate:"required,isoTime"`
	Data            any       `json:"data"`
}
