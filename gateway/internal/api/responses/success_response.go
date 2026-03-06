package responses

import "github.com/labstack/echo/v4"

type successResponse struct {
	Success bool   `json:"success"`
	Detail  string `json:"detail"`
	Data    any    `json:"data,omitempty"`
}

func SuccessResponseMessage(c echo.Context, statuscode int, detail string, data any) error {
	res := successResponse{
		Success: true,
		Detail:  detail,
		Data:    data,
	}
	return c.JSON(statuscode, res)
}
