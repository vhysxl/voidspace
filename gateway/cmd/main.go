package main

import (
	"log"
	"voidspaceGateway/bootstrap"
	"voidspaceGateway/internal/api/router"

	"github.com/labstack/echo/v4"
)

func main() {
	app, err := bootstrap.App()
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	router.SetupRoutes(app, e)

	e.Logger.Fatal(e.Start(app.Config.Port))
}
