package handlers

import (
	"time"
	"voidspaceGateway/internal/service"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type UserHandler struct {
	ContextTimeout time.Duration
	Logger         *zap.Logger
	Validator      *validator.Validate
	UserService    *service.UserService
}

func NewUserHandler(
	timeout time.Duration,
	logger *zap.Logger,
	validator *validator.Validate,
	userService *service.UserService,
) *UserHandler {
	return &UserHandler{
		ContextTimeout: timeout,
		Logger:         logger,
		Validator:      validator,
		UserService:    userService,
	}
}
