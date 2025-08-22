package main

import (
	"log"
	"net/http"
	"voidspaceGateway/bootstrap"

	"github.com/labstack/echo/v4"
)

func main() {
	app, err := bootstrap.App()
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(app.Config.Port))

}
