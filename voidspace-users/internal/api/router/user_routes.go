package router

import (
	"voidspace/users/bootstrap"
	"voidspace/users/internal/api/handler"
	"voidspace/users/internal/repository"
	"voidspace/users/internal/usecase"

	"github.com/go-chi/chi/v5"
)

func UserRoutes(r chi.Router, app *bootstrap.Application) {
	userRepository := repository.NewUserRepository(app.DB)
	userUsecase := usecase.NewUserUsecase(userRepository, app.DBContextTimeout) //dbctx to passed to repo

	userHandler := handler.NewUserHandler(
		userUsecase, app.Validator, app.HandlerContextTimeout,
		app.Logger,
	)

	r.Get("/me", userHandler.HandleGetCurrentUser)
}
