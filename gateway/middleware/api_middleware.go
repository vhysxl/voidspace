package middleware

import (
	"fmt"
	"net/http"
	"voidspaceGateway/internal/api/responses"
	"voidspaceGateway/internal/constants"

	"github.com/labstack/echo/v4"
)

func ApiMiddleware(apiKey string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			receivedKey := c.Request().Header.Get("x-api-key")
			fmt.Println(receivedKey == apiKey)
			if receivedKey != apiKey {
				return responses.ErrorResponseMessage(c, http.StatusUnauthorized, constants.ErrUnauthorized)
			}
			return next(c)
		}
	}
}
