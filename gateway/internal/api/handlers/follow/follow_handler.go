package follow

import (
	"crypto/rsa"
	"time"
	user_service "voidspaceGateway/internal/service/user"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type FollowHandler struct {
	ContextTimeout time.Duration
	Logger         *zap.Logger
	Validator      *validator.Validate
	UserService    *user_service.UserService
	PublicKey      *rsa.PublicKey
}

func NewFollowHandler(
	timeout time.Duration,
	logger *zap.Logger,
	validator *validator.Validate,
	userService *user_service.UserService,
	PublicKey *rsa.PublicKey,
) *FollowHandler {
	return &FollowHandler{
		ContextTimeout: timeout,
		Logger:         logger,
		Validator:      validator,
		UserService:    userService,
		PublicKey:      PublicKey,
	}
}
