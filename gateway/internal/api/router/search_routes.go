package router

import (
	"voidspaceGateway/internal/api/handlers/search"

	"github.com/labstack/echo/v4"
)

func SearchRoutes(api *echo.Group, h *search.SearchHandler) {
	api.GET("/search", h.Search)
}
