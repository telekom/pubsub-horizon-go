package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"testing"
)

var validate *validator.Validate

func TestNewValidator(t *testing.T) {
	var assertions = assert.New(t)
	assertions.NotPanics(func() {
		validate = NewValidator()
	})
}

func TestValidateEventType(t *testing.T) {
	var inputs = []struct {
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
			var assertions = assert.New(t)
			var err = validate.Struct(input.Value)

			if input.ExpectError {
				assertions.Error(err)
			} else {
				assertions.NoError(err)
			}
		})
	}
}

func TestValidateIsoTime(t *testing.T) {
	var inputs = []struct {
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
			var assertions = assert.New(t)

			if input.ExpectError {
				assertions.Error(validate.Struct(input.Value))
			} else {
				assertions.NoError(validate.Struct(input.Value))
			}
		})
	}
}
