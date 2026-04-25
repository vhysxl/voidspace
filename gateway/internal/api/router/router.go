package router

import (
	"voidspaceGateway/bootstrap"
	auth_handler "voidspaceGateway/internal/api/handlers/auth"
	comment_handler "voidspaceGateway/internal/api/handlers/comment"
	follow_handler "voidspaceGateway/internal/api/handlers/follow"
	post_handler "voidspaceGateway/internal/api/handlers/post"
	upload_handler "voidspaceGateway/internal/api/handlers/upload"
	user_handler "voidspaceGateway/internal/api/handlers/user"
	search_handler "voidspaceGateway/internal/api/handlers/search"
	"voidspaceGateway/middleware"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(app *bootstrap.Application, e *echo.Echo) {
	// HANDLERS
	authHandler := auth_handler.NewAuthHandler(
		app.ContextTimeout,
		app.Logger,
		app.Validator,
		app.UserService,
		app.Config.PublicKey,
	)

	followHandler := follow_handler.NewFollowHandler(
		app.ContextTimeout,
		app.Logger,
		app.Validator,
		app.UserService,
		app.Config.PublicKey,
	)

	userHandler := user_handler.NewUserHandler(
		app.ContextTimeout,
		app.Logger,
		app.Validator,
		app.UserService,
		app.Config.PublicKey,
	)

	postHandler := post_handler.NewPostHandler(
		app.ContextTimeout,
		app.Logger,
		app.Validator,
		app.PostService,
	)

	commentHandler := comment_handler.NewCommentHandler(
		app.ContextTimeout,
		app.Logger,
		app.Validator,
		app.CommentService,
	)

	uploadHandler := upload_handler.NewUploadHandler(
		app.ContextTimeout,
		app.Logger,
		app.Validator,
		app.UploadService,
	)

	searchHandler := search_handler.NewSearchHandler(
		app.UserService,
		app.PostService,
		app.CommentService,
		app.Logger,
	)

	// MIDDLEWARE
	authMiddleware := middleware.AuthMiddleware((app.Config.PublicKey))
	optionalAuthMiddleware := middleware.OptionalAuthMiddleware(app.Config.PublicKey)
	apiMiddleware := middleware.ApiMiddleware(app.Config.ApiSecret)

	api := e.Group("/api/v2")
	api.Use(apiMiddleware)

	// Routes
	AuthRoutes(api, authHandler, authMiddleware)
	FollowRoutes(api, followHandler, authMiddleware)
	UserRoutes(api, userHandler, followHandler, optionalAuthMiddleware, authMiddleware)
	PostRoutes(api, postHandler, optionalAuthMiddleware, authMiddleware)
	CommentRoutes(api, commentHandler, authMiddleware)
	UploadRoutes(api, uploadHandler, authMiddleware)
	SearchRoutes(api, searchHandler)
}
