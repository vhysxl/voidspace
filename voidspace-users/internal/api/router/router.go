package router

import (
	"net/http"
	"voidspace/users/bootstrap"

	"github.com/go-chi/chi/v5"
)

func Router(r *chi.Mux, app *bootstrap.Application) {

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("user service online"))
		})
		r.Route("/auth", func(r chi.Router) {
			AuthRoutes(r, app)
		})
	})
}
