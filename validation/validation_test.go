// Copyright 2025 Deutsche Telekom AG
//
// SPDX-License-Identifier: Apache-2.0

package validation

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

var validate *validator.Validate

func TestNewValidator(t *testing.T) {
	assertions := assert.New(t)
	assertions.NotPanics(func() {
		validate = NewValidator()
	})
}

func TestValidateEventType(t *testing.T) {
	inputs := []struct {
		Name        string
		Value       any
		ExpectError bool
	}{
		// Valid
		{
			Name: "valid struct",
			Value: struct {
				Field string `validate:"eventType"`
			}{Field: "my.valid.event.type.v1"},
			ExpectError: false,
		},

		// Invalid
		{
			Name: "invalid struct",
			Value: struct {
				Field string `validate:"eventType"`
			}{Field: "my.invalid.event.type.@v1@"},
			ExpectError: true,
		},
	}

	for _, input := range inputs {
		t.Run(input.Name, func(t *testing.T) {
			assertions := assert.New(t)
			err := validate.Struct(input.Value)

			if input.ExpectError {
				assertions.Error(err)
			} else {
				assertions.NoError(err)
			}
		})
	}
}

func TestValidateIsoTime(t *testing.T) {
	inputs := []struct {
		Name        string
		Value       any
		ExpectError bool
	}{
		// Valid
		{
			Name: "valid struct",
			Value: struct {
				Field string `validate:"isoTime"`
			}{Field: "2024-05-28T11:21:25.997Z"},
			ExpectError: false,
		},

		// Invalid
		{
			Name: "invalid struct",
			Value: struct {
				Field string `validate:"isoTime"`
			}{Field: "2024-05-28T11:21Z"},
			ExpectError: true,
		},
	}

	for _, input := range inputs {
		t.Run(input.Name, func(t *testing.T) {
			assertions := assert.New(t)

			if input.ExpectError {
				assertions.Error(validate.Struct(input.Value))
			} else {
				assertions.NoError(validate.Struct(input.Value))
			}
		})
	}
}
