package router

import (
	"voidspaceGateway/bootstrap"
	"voidspaceGateway/internal/api/handlers"
	"voidspaceGateway/middleware"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(app *bootstrap.Application, e *echo.Echo) {
	// HANDLERS
	authHandler := handlers.NewAuthHandler(app.ContextTimeout, app.Logger, app.Validator, app.AuthService, app.Config.PublicKey)
	userHandler := handlers.NewUserHandler(app.ContextTimeout, app.Logger, app.Validator, app.UserService)
	postHandler := handlers.NewPostHandler(app.ContextTimeout, app.Logger, app.Validator, app.PostService)
	likeHandler := handlers.NewLikeHandler(app.ContextTimeout, app.Logger, app.Validator, app.LikeService)

	api := e.Group("/api/v1")

	// Auth group
	auth := api.Group("/auth")
	auth.POST("/register", authHandler.Register)
	auth.POST("/login", authHandler.Login)
	auth.POST("/logout", authHandler.Logout)
	//protected route
	auth.POST("/refresh", authHandler.RefreshToken)

	//user group
	// Public routes
	usersPublic := api.Group("/users")
	usersPublic.GET("/:username", userHandler.GetUser)
	// Protected routes
	usersPrivate := api.Group("/users")
	usersPrivate.Use(middleware.AuthMiddleware(app.Config.PublicKey))
	usersPrivate.GET("/me", userHandler.GetCurrentUser)
	usersPrivate.PUT("/me", userHandler.UpdateProfile)
	usersPrivate.DELETE("/me", userHandler.DeleteUser)

	//follow group
	follow := api.Group("/follow")
	follow.Use(middleware.AuthMiddleware(app.Config.PublicKey))
	follow.POST("/:username", userHandler.Follow)
	follow.DELETE("/:username", userHandler.Unfollow)

	// Posts group
	// Public posts group
	postsPublic := api.Group("/posts")
	postsPublic.GET("/:id", postHandler.GetPost)
	postsPublic.GET("/user/:username", postHandler.GetUserPosts)
	// Protected posts group
	postsPrivate := api.Group("/posts")
	postsPrivate.Use(middleware.AuthMiddleware(app.Config.PublicKey))
	postsPrivate.POST("/", postHandler.Create)
	postsPrivate.PUT("/:id", postHandler.Update)
	postsPrivate.DELETE("/:id", postHandler.Delete)

	// Feed group
	feed := api.Group("/feed")
	feed.GET("/", func(c echo.Context) error {
		return c.String(200, "Global feed")
	})
	followFeed := api.Group("/feed")
	followFeed.Use(middleware.AuthMiddleware(app.Config.PublicKey))
	feed.GET("/followFeed", func(c echo.Context) error {
		return c.String(200, "follow feed")
	})

	// LIkes group
	likes := api.Group("/likes")
	likes.Use(middleware.AuthMiddleware(app.Config.PublicKey))
	likes.POST("/:postId", likeHandler.Like)
	likes.DELETE("/:postId", likeHandler.Unlike)
}
