package router

import (
	post_handler "voidspaceGateway/internal/api/handlers/post"

	"github.com/labstack/echo/v4"
)

func PostRoutes(
	api *echo.Group,
	postHandler *post_handler.PostHandler,
	optionalAuthMiddleware echo.MiddlewareFunc,
	authMiddleware echo.MiddlewareFunc,
) {
	// Public post routes
	postsPublic := api.Group("/posts")
	postsPublic.Use(optionalAuthMiddleware)
	postsPublic.GET("/:id", postHandler.GetPost)
	postsPublic.GET("/user/:username", postHandler.GetUserPosts)
	postsPublic.GET("/liked/:username", postHandler.GetLikedPosts)

	// Protected post routes
	postsPrivate := api.Group("/posts")
	postsPrivate.Use(authMiddleware)
	postsPrivate.POST("", postHandler.Create)
	postsPrivate.PUT("/:id", postHandler.Update)
	postsPrivate.DELETE("/:id", postHandler.Delete)
	postsPrivate.POST("/:id/like", postHandler.LikePost)
	postsPrivate.DELETE("/:id/like", postHandler.UnlikePost)

	// Feed routes
	feed := api.Group("/feed")
	feed.Use(optionalAuthMiddleware)
	feed.GET("", postHandler.GetGlobalFeed)

	feedFollowing := api.Group("/feed/following")
	feedFollowing.Use(authMiddleware)
	feedFollowing.GET("", postHandler.GetFollowingFeed)
}
