package main

import (
	"errors"
	"os"

	"github.com/go-playground/validator/v10"
)

// Retrieves the error message according to the validation applied
func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "phone":
		return "The number must be 10 digits"
	case "email":
		return "The mail does not comply with the correct format"
	case "alphanum":
		return "The username does not comply with the correct format"
	case "required_if":
		return "Email or username is required"
	case "password":
		return "The password does not comply with the format"
	}
	return "Unknown error"
}

// Retrieves the group of validation errors
func getError(ve validator.ValidationErrors, err error) []ErrorFormatMsg {
	if errors.As(err, &ve) {
		out := make([]ErrorFormatMsg, len(ve))
		for i, fe := range ve {
			out[i] = ErrorFormatMsg{fe.Field(), getErrorMsg(fe)}
		}
		return out
	}

	return nil
}

// Retrieves an environment variable, if it does not exist, it returns the default value
func getEnv(key, defaultValue string) string {
	if len(os.Getenv(key)) == 0 {
		return defaultValue
	} else {
		return os.Getenv(key)
	}
}
