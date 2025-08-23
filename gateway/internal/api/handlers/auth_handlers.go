package handlers

import (
	"time"
	"voidspaceGateway/bootstrap"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type AuthHandler struct {
	ContextTimeout time.Duration
	Logger         *zap.Logger
	App            *bootstrap.Application
	Validator      *validator.Validate
}

func NewAuthHandler(
	timeout time.Duration,
	logger *zap.Logger,
	validator *validator.Validate,
	app *bootstrap.Application,
) *AuthHandler {
	return &AuthHandler{
		ContextTimeout: timeout,
		Logger:         logger,
		Validator:      validator,
		App:            app,
	}
}

func (ah *AuthHandler) Login(c echo.Context) error {
	return c.String(200, "Login endpoint not implemented yet")
}
