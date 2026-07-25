package utils

import (
	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// getValidationMessage returns a user-friendly error message based on the validation tag.
func getValidationMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email format"
	case "min":
		return "Value is too short"
	case "max":
		return "Value is too long"
	case "oneof":
		return "Invalid value"
	default:
		return "Invalid value"
	}
}

// ParseValidationErrors takes a validation error and converts it into a slice of ValidationError structs.
func ParseValidationErrors(err error) []ValidationError {
	var validationErrors []ValidationError

	fieldErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return validationErrors
	}

	for _, fieldError := range fieldErrors {
		message := getValidationMessage(fieldError)

		validationErrors = append(validationErrors, ValidationError{
			Field:   ToSnakeCase(fieldError.Field()),
			Message: message,
		})
	}

	return validationErrors
}
