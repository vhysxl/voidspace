package router

import (
	follow_handler "voidspaceGateway/internal/api/handlers/follow"

	"github.com/labstack/echo/v4"
)

func FollowRoutes(
	api *echo.Group,
	followHandler *follow_handler.FollowHandler,
	authMiddleware echo.MiddlewareFunc,
) {
	follow := api.Group("/follow")
	follow.Use(authMiddleware)

	follow.POST("/:username", followHandler.Follow)
	follow.DELETE("/:username", followHandler.Unfollow)
}
