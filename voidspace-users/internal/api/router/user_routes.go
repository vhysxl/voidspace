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
	userUsecase := usecase.NewUserUsecase(userRepository, app.DBContextTimeout)
	profileRepository := repository.NewProfileRepository(app.DB)
	profileUsecase := usecase.NewProfileUsecase(profileRepository, userRepository, app.DBContextTimeout)

	userHandler := handler.NewUserHandler(
		userUsecase, profileUsecase, app.Validator, app.HandlerContextTimeout,
		app.Logger,
	)

	r.Get("/me", userHandler.HandleGetCurrentUser)
	r.Get("/{username}", userHandler.HandleGetUser)
	r.Patch("/profile", userHandler.HandleUpdateProfile)
}
