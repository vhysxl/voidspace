package router

import (
	comment_handler "voidspaceGateway/internal/api/handlers/comment"

	"github.com/labstack/echo/v4"
)

func CommentRoutes(
	api *echo.Group,
	commentHandler *comment_handler.CommentHandler,
	authMiddleware echo.MiddlewareFunc,
) {
	// Protected comment routes
	comment := api.Group("/comments")
	comment.Use(authMiddleware)
	comment.POST("", commentHandler.Create)
	comment.DELETE("/:id", commentHandler.Delete)

	// Public comment routes
	commentPublic := api.Group("/comments")
	commentPublic.GET("/post/:id", commentHandler.GetAllByPostID)
	commentPublic.GET("/user/:username", commentHandler.GetAllByUser)
}
