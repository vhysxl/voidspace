package router

import (
	follow_handler "voidspaceGateway/internal/api/handlers/follow"
	user_handler "voidspaceGateway/internal/api/handlers/user"

	"github.com/labstack/echo/v4"
)

func UserRoutes(
	api *echo.Group,
	userHandler *user_handler.UserHandler,
	followHandler *follow_handler.FollowHandler,
	optionalAuthMiddleware echo.MiddlewareFunc,
	authMiddleware echo.MiddlewareFunc,
) {
	user := api.Group("/user")
	user.Use(optionalAuthMiddleware)

	user.GET("/:username", userHandler.GetUser)
	user.GET("/:username/followers", followHandler.ListFollowers)
	user.GET("/:username/following", followHandler.ListFollowing)
	user.GET("/me", userHandler.GetCurrentUser, authMiddleware)
	user.PUT("/me", userHandler.UpdateProfile, authMiddleware)
	user.DELETE("/me", userHandler.DeleteUser, authMiddleware)
}
