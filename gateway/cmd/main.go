package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"voidspaceGateway/bootstrap"
	"voidspaceGateway/internal/api/router"
	// "voidspaceGateway/temporal"

	cstmMiddleware "voidspaceGateway/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

func main() {
	app, err := bootstrap.App()
	if err != nil {
		log.Fatal(err)
	}

	// temporal.RegisterTemporal(app)

	go func() {
		if err := app.TemporalService.TemporalStart(); err != nil {
			app.Logger.Fatal("Failed to start Temporal worker", zap.Error(err))
		}
	}()

	e := echo.New()
	e.Use(middleware.RateLimiterWithConfig(cstmMiddleware.RateLimitConfig()))
	e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
		XSSProtection:      "1; mode=block",
		ContentTypeNosniff: "nosniff",
		XFrameOptions:      "DENY",
		HSTSMaxAge:         3600,
	}))

	e.Use(middleware.BodyLimit("2M"))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "x-api-key"},
		AllowCredentials: true,
	}))

	router.SetupRoutes(app, e)

	e.File("/openapi", "/app/api_docs.yaml")

	e.GET("/docs", func(c echo.Context) error {
		html := `
	<!DOCTYPE html>
	<html>
	  <head>
	    <title>Swagger UI</title>
	    <link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist/swagger-ui.css" />
	  </head>
	  <body>
	    <div id="swagger-ui"></div>
	    <script src="https://unpkg.com/swagger-ui-dist/swagger-ui-bundle.js"></script>
	    <script>
	      const ui = SwaggerUIBundle({
	        url: '/openapi',
	        dom_id: '#swagger-ui',
	      })
	    </script>
	  </body>
	</html>
	`
		return c.HTML(200, html)
	})

	// ==== SIGNAL HANDLER ====
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		if err := e.Start(":" + app.Config.Port); err != nil && err != http.ErrServerClosed {
			app.Logger.Fatal("Echo failed", zap.Error(err))
		}
	}()

	<-ctx.Done()

	app.Logger.Info("Shutting down...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_ = e.Shutdown(shutdownCtx)
	app.TemporalService.Stop()

	app.Logger.Info("Exit cleanly")
}
