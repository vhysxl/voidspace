package middleware

import (
	"crypto/rsa"
	"net/http"
	"strings"
	"voidspaceGateway/internal/api/responses"
	"voidspaceGateway/internal/constants"
	"voidspaceGateway/internal/models"
	"voidspaceGateway/utils"

	"github.com/labstack/echo/v4"
)

// AuthMiddleware creates an Echo middleware function that validates JWT tokens using RSA public key verification.
// It extracts user information from valid tokens and makes it available to subsequent handlers.
func AuthMiddleware(PublicKey *rsa.PublicKey) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Extract the Authorization header from the request
			auth := c.Request().Header.Get("Authorization")
			if auth == "" {
				return responses.ErrorResponseMessage(c, http.StatusUnauthorized, constants.ErrUnauthorized)
			}

			// Parse the Authorization header to ensure it follows "Bearer <token>" format
			parts := strings.Split(auth, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				return responses.ErrorResponseMessage(c, http.StatusUnauthorized, constants.ErrUnauthorized)
			}
			// Verify the JWT token using the provided RSA public key
			claims, err := utils.VerifyRefreshToken(parts[1], PublicKey)
			if err != nil {
				return responses.ErrorResponseMessage(c, http.StatusUnauthorized, constants.ErrUnauthorized)
			}

			user := &models.AuthUser{
				ID:       claims["ID"].(string),
				Username: claims["Username"].(string),
			}
			c.Set("authUser", user)

			// Continue to the next handler
			return next(c)
		}
	}
}
