package handler

import (
	"crypto/rsa"
	"time"
	"voidspace/users/internal/domain"
	pb "voidspace/users/proto/users/v1"

	"go.uber.org/zap"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer

	UserUsecase          domain.UserUsecase
	ProfileUsecase       domain.ProfileUsecase
	FollowUsecase        domain.FollowUsecase
	Logger               *zap.Logger
	ContextTimeout       time.Duration
	PrivateKey           *rsa.PrivateKey
	AccessTokenDuration  time.Duration
	RefreshTokenDuration time.Duration
}

func NewUserHandler(
	userUsecase domain.UserUsecase,
	profileUsecase domain.ProfileUsecase,
	followUsecase domain.FollowUsecase,
	timeout time.Duration,
	logger *zap.Logger,
	privateKey *rsa.PrivateKey,
	accessTokenDuration time.Duration,
	refreshTimeDuration time.Duration,
) pb.UserServiceServer {
	return &UserHandler{
		UserUsecase:          userUsecase,
		ProfileUsecase:       profileUsecase,
		FollowUsecase:        followUsecase,
		Logger:               logger,
		ContextTimeout:       timeout,
		PrivateKey:           privateKey,
		AccessTokenDuration:  accessTokenDuration,
		RefreshTokenDuration: refreshTimeDuration,
	}
}
