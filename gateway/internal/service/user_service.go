package service

import (
	"time"
	userpb "voidspaceGateway/proto/generated/users"

	"go.uber.org/zap"
)

type UserService struct {
	ContextTimeout time.Duration
	Logger         *zap.Logger
	UserClient     userpb.UserServiceClient
}

func NewUserService(timeout time.Duration, logger *zap.Logger, userClient userpb.UserServiceClient) *UserService {
	return &UserService{
		ContextTimeout: timeout,
		Logger:         logger,
		UserClient:     userClient,
	}
}
