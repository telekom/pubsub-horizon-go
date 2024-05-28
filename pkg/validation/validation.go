package validation

import "github.com/go-playground/validator/v10"

func NewValidator() *validator.Validate {
	var validate = validator.New()

	if err := validate.RegisterValidation("eventType", ValidateEventType); err != nil {
		panic(err)
	}

	if err := validate.RegisterValidation("isoTime", ValidateIsoTime); err != nil {
		panic(err)
	}

	return validate
}
