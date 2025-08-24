package utils

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func FormatValidationError(err error) string {
	var fieldLabels = map[string]string{
		"usernameoremail": "Username or Email",
		"password":        "Password",
		"username":        "Username",
		"email":           "Email",
	}

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		var messages []string

		for _, v := range validationErrors {
			field := strings.ToLower(v.Field())

			label, exists := fieldLabels[field]
			if !exists {
				label = field
			}

			switch v.Tag() {
			case "required":
				messages = append(messages, fmt.Sprintf("%s is required", label))
			case "email":
				messages = append(messages, fmt.Sprintf("%s must be a valid email", label))
			case "min":
				messages = append(messages, fmt.Sprintf("%s must be at least %s characters", label, v.Param()))
			case "max":
				messages = append(messages, fmt.Sprintf("%s must be at most %s characters", label, v.Param()))
			default:
				messages = append(messages, fmt.Sprintf("%s is invalid", label))
			}
		}

		return strings.Join(messages, ", ")
	}

	return "Invalid request format"
}
