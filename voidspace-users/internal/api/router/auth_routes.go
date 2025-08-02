package router

import (
	"database/sql"
	"time"
	"voidspace/users/internal/api/handler"
	"voidspace/users/internal/repository"
	"voidspace/users/internal/usecase"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

func AuthRoutes(r chi.Router, db *sql.DB, validator *validator.Validate, contextTimeout time.Duration) {
	userRepository := repository.NewUserRepository(db)
	registerUsecase := usecase.NewRegisterUsecase(userRepository, contextTimeout)
	loginUsecase := usecase.NewLoginUsecase(userRepository, contextTimeout)

	registerHandler := handler.NewRegisterHandler(registerUsecase, validator)
	loginHandler := handler.NewLoginHandler(loginUsecase, validator)

	r.Post("/register", registerHandler.HandleRegister)
	r.Post("/login", loginHandler.HandleLogin)
}
