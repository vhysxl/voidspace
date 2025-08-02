package router

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

func Router(r *chi.Mux, db *sql.DB, validator *validator.Validate, contextTimeout time.Duration) {

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("user service online"))
		})
		r.Route("/auth", func(r chi.Router) {
			AuthRoutes(r, db, validator, contextTimeout)
		})
	})
}
