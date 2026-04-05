package middleware

import (
	"crypto/rsa"
	"errors"
	"net/http"
	"strings"
	"voidspaceGateway/internal/api/responses"
	shared_constants "github.com/vhysxl/voidspace/shared/utils/constants"
	"voidspaceGateway/internal/models"
	"voidspaceGateway/utils"

	"github.com/labstack/echo/v4"
)

func extractAuthUser(c echo.Context, publicKey *rsa.PublicKey) (*models.AuthUser, error) {
	// Extract the Authorization header from the request
	auth := c.Request().Header.Get("Authorization")
	if auth == "" {
		return nil, errors.New("no token")
	}

	// Parse the Authorization header to ensure it follows "Bearer <token>" format
	parts := strings.Split(auth, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return nil, errors.New("invalid format")
	}

	// Verify the JWT token using the provided RSA public key
	claims, err := utils.VerifyToken(parts[1], publicKey)
	if err != nil {
		return nil, err
	}

	user := &models.AuthUser{}
	if id, ok := claims["ID"].(string); ok {
		user.ID = id
	}
	if username, ok := claims["Username"].(string); ok {
		user.Username = username
	}

	return user, nil
}

func ApiMiddleware(apiKey string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			receivedKey := c.Request().Header.Get("x-api-key")
			if receivedKey != apiKey {
				return responses.ErrorResponseMessage(c, http.StatusUnauthorized, shared_constants.Unauthorized)
			}
			return next(c)
		}
	}
}

// AuthMiddleware creates an Echo middleware function that validates JWT tokens using RSA public key verification.
// It extracts user information from valid tokens and makes it available to subsequent handlers.
func AuthMiddleware(publicKey *rsa.PublicKey) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user, err := extractAuthUser(c, publicKey)
			if err != nil {
				return responses.ErrorResponseMessage(c, http.StatusUnauthorized, shared_constants.Unauthorized)
			}
			c.Set("authUser", user)

			// Continue to the next handler
			return next(c)
		}
	}
}

func OptionalAuthMiddleware(publicKey *rsa.PublicKey) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user, _ := extractAuthUser(c, publicKey) // error ignored, set nil if failed
			c.Set("authUser", user)
			return next(c)
		}
	}
}
