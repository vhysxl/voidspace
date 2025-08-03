package router

import (
	"voidspace/users/bootstrap"
	"voidspace/users/internal/api/handler"
	"voidspace/users/internal/repository"
	"voidspace/users/internal/usecase"

	"github.com/go-chi/chi/v5"
)

func AuthRoutes(r chi.Router, app *bootstrap.Application) {
	userRepository := repository.NewUserRepository(app.DB)
	registerUsecase := usecase.NewRegisterUsecase(userRepository, app.DBContextTimeout)
	loginUsecase := usecase.NewLoginUsecase(userRepository, app.DBContextTimeout)

	registerHandler := handler.NewRegisterHandler(registerUsecase, app.Validator, app.PrivateKey, app.HandlerContextTimeout, app.AccessTokenDuration, app.RefreshTokenDuration)
	loginHandler := handler.NewLoginHandler(loginUsecase, app.Validator)

	r.Post("/register", registerHandler.HandleRegister)
	r.Post("/login", loginHandler.HandleLogin)
}
