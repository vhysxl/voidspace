package router

import (
	upload_handler "voidspaceGateway/internal/api/handlers/upload"

	"github.com/labstack/echo/v4"
)

func UploadRoutes(api *echo.Group, uploadHandler *upload_handler.UploadHandler, authMiddleware echo.MiddlewareFunc) {
	upload := api.Group("/upload")
	upload.Use(authMiddleware)

	upload.POST("/signed-url", uploadHandler.GenerateSignedURL)
}
