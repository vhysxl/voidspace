package utils

import (
	"net/http"
	"voidspaceGateway/internal/api/responses"
	"voidspaceGateway/internal/constants"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// this code is buggy, fix later
func BindAndValidate(c echo.Context, validator *validator.Validate, req any) error {
	if err := c.Bind(req); err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, constants.ErrInvalidRequest)
	}

	if err := validator.Struct(req); err != nil {
		return responses.ErrorResponseMessage(c, http.StatusBadRequest, FormatValidationError(err))
	}

	return nil
}
