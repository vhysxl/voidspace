package responses

import (
	"github.com/labstack/echo/v4"
)

type errorResponse struct {
	Success bool   `json:"success"`
	Detail  string `json:"detail"`
}

func ErrorResponseMessage(c echo.Context, statuscode int, detail string) error {
	if statuscode == 503 {
		return c.JSON(statuscode, "Service is temporarily unavailable. Please try again later.")
	}
	res := errorResponse{
		Success: false,
		Detail:  detail,
	}
	return c.JSON(statuscode, res)
}
