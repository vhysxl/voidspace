package main

import (
	"log"
	"net/http"
	"voidspace/users/bootstrap"
	"voidspace/users/internal/api/router"

	"github.com/go-chi/chi/v5"
)

func main() {
	app, err := bootstrap.App()

	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()

	router.Router(r, app.DB, app.Validator, app.ContextTimeout)

	log.Println("Listening", app.Config.Port)
	err = http.ListenAndServe(app.Config.Port, r)
	if err != nil {
		log.Fatalf("Server error: %v", err)
	}

}
