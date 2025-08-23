package router

import (
	"voidspaceGateway/bootstrap"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(app *bootstrap.Application, e *echo.Echo) {

	api := e.Group("/api/v1")

	// Auth group
	auth := api.Group("/auth")
	auth.POST("/register", func(c echo.Context) error {
		return c.String(200, "Register endpoint not implemented yet")
	})
	auth.POST("/login", func(c echo.Context) error {
		return c.String(200, "Login endpoint not implemented yet")
	})
	//protected route
	auth.POST("/refresh", func(c echo.Context) error {
		return c.String(200, "Refresh token endpoint not implemented yet")
	})

	//user group
	// Public routes
	usersPublic := api.Group("/users")
	usersPublic.GET("/username/:username", func(c echo.Context) error {
		return c.String(200, "Get user by username endpoint not implemented yet")
	})
	// Protected routes
	usersPrivate := api.Group("/users") // to do : add middleware here for protecting routes
	usersPrivate.GET("/profile", func(c echo.Context) error {
		return c.String(200, "Get user by ID endpoint not implemented yet")
	})
	usersPrivate.PATCH("/profile", func(c echo.Context) error {
		return c.String(200, "Update user endpoint not implemented yet")
	})
	usersPrivate.DELETE("/profile0", func(c echo.Context) error {
		return c.String(200, "Delete user endpoint not implemented yet")
	})

	//follow group
	follow := api.Group("/follow") // to do : add middleware here for protecting routes
	follow.POST("/:id", func(c echo.Context) error {
		return c.String(200, "Follow user endpoint not implemented yet")
	})
	follow.DELETE("/:id", func(c echo.Context) error {
		return c.String(200, "Unfollow user endpoint not implemented yet")
	})

	// Posts group
	// Public posts group
	postsPublic := api.Group("/posts")
	postsPublic.GET("/:id", func(c echo.Context) error {
		return c.String(200, "Get post by ID endpoint not implemented yet")
	})
	postsPublic.GET("/user/:userId", func(c echo.Context) error {
		return c.String(200, "Get posts by user ID endpoint not implemented yet")
	})
	// Protected posts group
	postsPrivate := api.Group("/posts") // to do : add middleware here for protecting routes
	postsPrivate.POST("/", func(c echo.Context) error {
		return c.String(200, "Create post endpoint not implemented yet")
	})
	postsPrivate.PATCH("/:id", func(c echo.Context) error {
		return c.String(200, "Update post endpoint not implemented yet")
	})
	postsPrivate.DELETE("/:id", func(c echo.Context) error {
		return c.String(200, "Delete post endpoint not implemented yet")
	})

	// Feed group
	feed := api.Group("/feed")
	feed.GET("/", func(c echo.Context) error {
		return c.String(200, "Global feed")
	})
	feed.GET("/followFeed", func(c echo.Context) error {
		return c.String(200, "follow feed")
	})

	// LIkes group
	likes := api.Group("/likes") // to do : add middleware here for protecting routes
	likes.POST("/:postId", func(c echo.Context) error {
		return c.String(200, "Like post endpoint not implemented yet")
	})
	likes.DELETE("/:postId", func(c echo.Context) error {
		return c.String(200, "Unlike post endpoint not implemented yet")
	})
}
