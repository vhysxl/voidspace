package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func SetRefreshTokenCookie(c echo.Context, refreshToken string) error {
	cookie := &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		HttpOnly: true,
		Secure:   false, // set to true in production
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
		MaxAge:   7 * 24 * 60 * 60,
	}
	c.SetCookie(cookie)
	return nil
}
