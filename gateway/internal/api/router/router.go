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
	feedHandler := handlers.NewFeedHandler(app.ContextTimeout, app.Logger, app.Validator, app.FeedService)
	commentHandler := handlers.NewCommentHandler(app.ContextTimeout, app.Logger, app.Validator, app.CommentService)
	uploadHandler := handlers.NewUploadHandler(app.ContextTimeout, app.Logger, app.Validator, app.UploadService)
	authMiddleware := middleware.AuthMiddleware((app.Config.PublicKey))
	optionalAuthMiddleware := middleware.OptionalAuthMiddleware(app.Config.PublicKey)

	api := e.Group("/api/v1")

	// Auth group
	auth := api.Group("/auth")
	auth.Use(middleware.ApiMiddleware("palalobautai"))
	auth.POST("/register", authHandler.Register)
	auth.POST("/login", authHandler.Login)
	auth.POST("/logout", authHandler.Logout)
	//protected route
	auth.POST("/refresh", authHandler.RefreshToken)

	//user group
	// Public routes
	usersPublic := api.Group("/users")
	usersPublic.Use(optionalAuthMiddleware)
	usersPublic.GET("/:username", userHandler.GetUser)
	// Protected routes
	usersPrivate := api.Group("/users")
	usersPrivate.Use(authMiddleware)
	usersPrivate.GET("/me", userHandler.GetCurrentUser)
	usersPrivate.PUT("/me", userHandler.UpdateProfile)
	usersPrivate.DELETE("/me", userHandler.DeleteUser)

	//follow group
	follow := api.Group("/follow")
	follow.Use(authMiddleware)
	follow.POST("/:username", userHandler.Follow)
	follow.DELETE("/:username", userHandler.Unfollow)

	// Posts group
	// Public posts group
	postsPublic := api.Group("/posts")
	postsPublic.Use(optionalAuthMiddleware)
	postsPublic.GET("/:id", postHandler.GetPost)
	postsPublic.GET("/user/:username", postHandler.GetUserPosts)
	// Protected posts group
	postsPrivate := api.Group("/posts")
	postsPrivate.Use(authMiddleware)
	postsPrivate.POST("/", postHandler.Create)
	postsPrivate.PUT("/:id", postHandler.Update)
	postsPrivate.DELETE("/:id", postHandler.Delete)

	// Global feed (bisa tanpa login)
	feed := api.Group("/feed")
	feed.Use(optionalAuthMiddleware)
	feed.GET("/", feedHandler.GetGlobalFeed)
	// Follow feed (harus login)
	followFeed := api.Group("/feed/following")
	followFeed.Use(authMiddleware)
	followFeed.GET("/", feedHandler.GetFollowFeed)

	// Likes group
	likes := api.Group("/likes")
	likes.Use(authMiddleware)
	likes.POST("/:postId", likeHandler.Like)
	likes.DELETE("/:postId", likeHandler.Unlike)

	// Upload Group
	upload := api.Group("/upload/signed-url")
	upload.Use(authMiddleware)
	upload.POST("", uploadHandler.GenerateSignedURL)

	//Comment Group
	comment := api.Group("/comment")
	comment.Use(authMiddleware)
	comment.POST("", commentHandler.Create)
	commentPublic := api.Group("/comment")
	commentPublic.GET("/post/:id", commentHandler.GetAllByPostID)
}
