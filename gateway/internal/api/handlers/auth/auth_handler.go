package auth

import (
	"crypto/rsa"
	"time"
	user_service "voidspaceGateway/internal/service/user"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type AuthHandler struct {
	ContextTimeout time.Duration
	Logger         *zap.Logger
	Validator      *validator.Validate
	UserService    *user_service.UserService
	PublicKey      *rsa.PublicKey
}

func NewAuthHandler(
	timeout time.Duration,
	logger *zap.Logger,
	validator *validator.Validate,
	userService *user_service.UserService,
	PublicKey *rsa.PublicKey,
) *AuthHandler {
	return &AuthHandler{
		ContextTimeout: timeout,
		Logger:         logger,
		Validator:      validator,
		UserService:    userService,
		PublicKey:      PublicKey,
	}
}
