// Copyright 2025 Deutsche Telekom AG
//
// SPDX-License-Identifier: Apache-2.0

package message

type Event struct {
	Id              string `json:"id"              validate:"required,uuid4"`
	Type            string `json:"type"            validate:"required,eventType"`
	Source          string `json:"source"          validate:"required"`
	SpecVersion     string `json:"specVersion"     validate:"required"`
	DataContentType string `json:"dataContentType"`
	DataRef         string `json:"dataRef"`
	Time            string `json:"time"            validate:"omitempty,isoTime"`
	Data            any    `json:"data"`
}
