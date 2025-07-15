package api

import (
	"database/sql"
	"log"
	"net/http"

	"voidspace-api/services/user"

	"github.com/go-chi/chi/v5"
)

type APIServer struct {//server struct
	addr string
	db   *sql.DB 
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
		return &APIServer{
			addr: addr,
			db: db,
		}
} 

func (s *APIServer) Run() error { // router mapping
	router := chi.NewRouter()

	router.Route("/api/v1", func(r chi.Router) {
		router.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Voidspace online!"))
		})

		userHandler := user.NewHandler()
		r.Mount("/auth", userHandler.AuthRoutes())
	})


	log.Println("Listening on", s.addr)
	return http.ListenAndServe(s.addr, router)
}