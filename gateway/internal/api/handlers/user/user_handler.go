package user

import (
	"crypto/rsa"
	"time"
	user_service "voidspaceGateway/internal/service/user"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type UserHandler struct {
	ContextTimeout time.Duration
	Logger         *zap.Logger
	Validator      *validator.Validate
	UserService    *user_service.UserService
	PublicKey      *rsa.PublicKey
}

func NewUserHandler(
	contextTimeout time.Duration,
	logger *zap.Logger,
	validator *validator.Validate,
	userService *user_service.UserService,
	publicKey *rsa.PublicKey,
) *UserHandler {
	return &UserHandler{
		ContextTimeout: contextTimeout,
		Logger:         logger,
		Validator:      validator,
		UserService:    userService,
		PublicKey:      publicKey,
	}
}
