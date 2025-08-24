package service

import (
	"context"
	"crypto/rsa"
	"time"

	"voidspaceGateway/internal/models"
	userpb "voidspaceGateway/proto/generated/users"

	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AuthService struct {
	ContextTimeout time.Duration
	Logger         *zap.Logger
	AuthClient     userpb.AuthServiceClient
	PublicKey      rsa.PublicKey
}

func NewAuthService(timeout time.Duration, logger *zap.Logger, authClient userpb.AuthServiceClient, publicKey rsa.PublicKey) *AuthService {
	return &AuthService{
		ContextTimeout: timeout,
		Logger:         logger,
		AuthClient:     authClient,
		PublicKey:      publicKey,
	}
}

func (as *AuthService) Login(ctx context.Context, req *models.LoginRequest) (*models.AuthResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, as.ContextTimeout)
	defer cancel()

	res, err := as.AuthClient.Login(ctx, &userpb.LoginRequest{
		EmailOrUsername: req.UsernameOrEmail,
		Password:        req.Password,
	})

	if err != nil {
		as.Logger.Error("failed to call AuthService.Login", zap.Error(err))
		return nil, err
	}

	return &models.AuthResponse{
		AccessToken:  res.GetAccessToken(),
		RefreshToken: res.GetRefreshToken(),
		ExpiresIn:    int64(res.GetExpiresIn()),
	}, nil
}

func (as *AuthService) Register(ctx context.Context, req *models.RegisterRequest) (*models.AuthResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, as.ContextTimeout)
	defer cancel()

	res, err := as.AuthClient.Register(ctx, &userpb.RegisterRequest{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		as.Logger.Error("failed to call AuthService.Login", zap.Error(err))
		return nil, err
	}

	return &models.AuthResponse{
		AccessToken:  res.GetAccessToken(),
		RefreshToken: res.GetRefreshToken(),
		ExpiresIn:    int64(res.GetExpiresIn()),
	}, nil
}

func (as *AuthService) RefreshToken(ctx context.Context, userId string, username string) (*models.AuthResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, as.ContextTimeout)
	defer cancel()

	md := metadata.New(map[string]string{
		"user_id":  userId,
		"username": username,
	})

	ctx = metadata.NewOutgoingContext(ctx, md)

	res, err := as.AuthClient.RefreshToken(ctx, &emptypb.Empty{})
	if err != nil {
		as.Logger.Error("failed to call AuthService.RefreshToken", zap.Error(err))
		return nil, err
	}

	return &models.AuthResponse{
		AccessToken:  res.GetAccessToken(),
		RefreshToken: res.GetRefreshToken(),
		ExpiresIn:    int64(res.GetExpiresIn()),
	}, nil
}
