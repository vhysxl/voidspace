package handler

import (
	"context"
	"crypto/rsa"
	"time"
	"voidspace/users/internal/domain"
	"voidspace/users/internal/usecase"
	pb "voidspace/users/proto/generated/users"
	errorutils "voidspace/users/utils/error"
	"voidspace/users/utils/interceptor"
	"voidspace/users/utils/token"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AuthHandler struct {
	pb.UnimplementedAuthServiceServer

	AuthUsecase          usecase.AuthUsecase
	PrivateKey           *rsa.PrivateKey
	ContextTimeout       time.Duration
	AccessTokenDuration  time.Duration
	RefreshTokenDuration time.Duration
	Logger               *zap.Logger
}

func NewAuthHandler(
	authUsecase usecase.AuthUsecase,
	privateKey *rsa.PrivateKey,
	contextTimeout time.Duration,
	accessTokenDuration time.Duration,
	refreshTokenDuration time.Duration,
	logger *zap.Logger,
) pb.AuthServiceServer {
	return &AuthHandler{
		AuthUsecase:          authUsecase,
		PrivateKey:           privateKey,
		ContextTimeout:       contextTimeout,
		AccessTokenDuration:  accessTokenDuration,
		RefreshTokenDuration: refreshTokenDuration,
		Logger:               logger,
	}
}

func (ah *AuthHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.AuthResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, ah.ContextTimeout)
	defer cancel()

	user, err := ah.AuthUsecase.Login(ctx, req.GetEmailOrUsername(), req.GetPassword())
	if err != nil {
		return nil, errorutils.HandleError(err, ah.Logger, "Login")
	}

	accessToken, err := token.CreateAccessToken(user, ah.PrivateKey, ah.AccessTokenDuration)
	if err != nil {
		ah.Logger.Error("Generate access token error", zap.Error(err))
		return nil, status.Error(codes.Internal, "token generation failed")
	}

	refreshToken, err := token.CreateRefreshToken(user, ah.PrivateKey, ah.RefreshTokenDuration)
	if err != nil {
		ah.Logger.Error("Generate refresh token error", zap.Error(err))
		return nil, status.Error(codes.Internal, "token generation failed")
	}

	return &pb.AuthResponse{
		RefreshToken: &refreshToken,
		AccessToken:  accessToken,
		ExpiresIn:    int32(ah.AccessTokenDuration.Seconds()),
		Message:      "Login Success",
	}, nil
}

func (ah *AuthHandler) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.AuthResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, ah.ContextTimeout)
	defer cancel()

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		ah.Logger.Error("Hashing password error", zap.Error(err))
		return nil, status.Error(codes.Internal, "password hashing failed")
	}

	hashed := string(hashedPassword)

	user, err := ah.AuthUsecase.Register(ctx, req.Username, req.Email, hashed)
	if err != nil {
		return nil, errorutils.HandleError(err, ah.Logger, "Register")
	}

	accessToken, err := token.CreateAccessToken(user, ah.PrivateKey, ah.AccessTokenDuration)
	if err != nil {
		ah.Logger.Error("Generate access token error", zap.Error(err))
		return nil, status.Error(codes.Internal, "token generation failed")
	}

	refreshToken, err := token.CreateRefreshToken(user, ah.PrivateKey, ah.RefreshTokenDuration)
	if err != nil {
		ah.Logger.Error("Generate refresh token error", zap.Error(err))
		return nil, status.Error(codes.Internal, "token generation failed")
	}

	return &pb.AuthResponse{
		RefreshToken: &refreshToken,
		AccessToken:  accessToken,
		ExpiresIn:    int32(ah.AccessTokenDuration.Seconds()),
		Message:      "Register Success",
	}, nil
}

func (ah *AuthHandler) RefreshToken(ctx context.Context, _ *emptypb.Empty) (*pb.AuthResponse, error) {
	userId, ok := ctx.Value(interceptor.CtxKeyUserID).(int)
	if !ok {
		ah.Logger.Error("Failed to get userId from context")
		return nil, status.Error(codes.Internal, "failed to get user ID from context")
	}

	username, ok := ctx.Value(interceptor.CtxKeyUsername).(string)
	if !ok {
		ah.Logger.Error("Failed to get username from context")
		return nil, status.Error(codes.Internal, "failed to get username from context")
	}

	user := &domain.User{
		Id:       int32(userId),
		Username: username,
	}

	accessToken, err := token.CreateAccessToken(user, ah.PrivateKey, ah.AccessTokenDuration)
	if err != nil {
		return nil, errorutils.HandleError(err, ah.Logger, "Refresh")
	}

	return &pb.AuthResponse{
		AccessToken: accessToken,
		ExpiresIn:   int32(ah.AccessTokenDuration.Seconds()),
		Message:     "Token refreshed successfully",
	}, nil
}
