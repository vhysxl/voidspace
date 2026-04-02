package auth

import (
	"net/http"
	"voidspaceGateway/internal/api/responses"
	"voidspaceGateway/internal/constants"

	"github.com/labstack/echo/v4"
)

func (h *AuthHandler) Logout(c echo.Context) error {
	cookie := &http.Cookie{
		Name:     "refresh_token",
		Value:    "",
		HttpOnly: true,
		Secure:   false, // set to true in production
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
		MaxAge:   -1, // Delete cookie
	}
	c.SetCookie(cookie)

	return responses.SuccessResponseMessage(
		c, http.StatusOK,
		constants.LogoutSuccess,
		nil,
	)
}
