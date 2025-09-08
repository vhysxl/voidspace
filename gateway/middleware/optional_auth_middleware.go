package middleware

import (
	"crypto/rsa"
	"strings"
	"voidspaceGateway/utils"

	"github.com/labstack/echo/v4"
)

func OptionalAuthMiddleware(publicKey *rsa.PublicKey) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			auth := c.Request().Header.Get("Authorization")

			if auth != "" {
				parts := strings.Split(auth, " ")
				if len(parts) == 2 && parts[0] == "Bearer" {
					claims, err := utils.VerifyRefreshToken(parts[1], publicKey)

					if err == nil {
						// Inject user info ke context
						if id, ok := claims["ID"].(string); ok { // JWT biasanya float64
							c.Set("ID", id)
						}
						if username, ok := claims["Username"].(string); ok {
							c.Set("username", username)
						}
					}
				}
			}

			return next(c)
		}
	}
}
