package main

import (
	"log"
	"net/http"
	"voidspace/users/bootstrap"
	"voidspace/users/internal/api/router"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func main() {
	app, err := bootstrap.App()

	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()

	router.Router(r, app)

	app.Logger.Info("Listening", zap.String("port", app.Config.Port))

	err = http.ListenAndServe(app.Config.Port, r)
	if err != nil {
		app.Logger.Fatal("Server error", zap.Error(err))
	}

}
