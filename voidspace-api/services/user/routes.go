package user

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
}

func (h *Handler) Routes() http.Handler {
	panic("unimplemented")
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) AuthRoutes() http.Handler {
	router := chi.NewRouter()
	
	router.Post("/login", h.Login)
	router.Post("/register", h.Register)
	return router
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Login here"))
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Register here"))
}
