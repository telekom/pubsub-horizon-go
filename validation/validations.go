// Copyright 2024 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package validation

import "github.com/go-playground/validator/v10"

func ValidateEventType(fl validator.FieldLevel) bool {
	var value = fl.Field().String()
	return EventTypeRegEx.MatchString(value)
}

func ValidateIsoTime(fl validator.FieldLevel) bool {
	var value = fl.Field().String()
	return Iso8601RegEx.MatchString(value)
}
