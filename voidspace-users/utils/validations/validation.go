package validations

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func FormatValidationError(err error) string {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		var messages []string

		for _, v := range validationErrors {
			field := strings.ToLower(v.Field())

			switch v.Tag() {
			case "required":
				messages = append(messages, fmt.Sprintf("%s is required", field))
			case "email":
				messages = append(messages, fmt.Sprintf("%s must be a valid email", field))
			case "min":
				messages = append(messages, fmt.Sprintf("%s must be at least %s characters", field, v.Param()))
			case "max":
				messages = append(messages, fmt.Sprintf("%s must be at most %s characters", field, v.Param()))
			default:
				messages = append(messages, fmt.Sprintf("%s is invalid", field))
			}
		}

		return strings.Join(messages, ", ")
	}

	return "Invalid request format"
}
