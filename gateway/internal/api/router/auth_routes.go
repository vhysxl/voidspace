package router

import (
	auth_handler "voidspaceGateway/internal/api/handlers/auth"

	"github.com/labstack/echo/v4"
)

func AuthRoutes(
	api *echo.Group,
	authHandler *auth_handler.AuthHandler,
	authMiddleware echo.MiddlewareFunc,
) {
	auth := api.Group("/auth")

	auth.POST("/register", authHandler.Register)
	auth.POST("/login", authHandler.Login)
	auth.POST("/logout", authHandler.Logout)
	auth.POST("/refresh", authHandler.RefreshToken)
}
