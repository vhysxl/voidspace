package handler

import (
	"context"
	"crypto/rsa"
	"time"
	"voidspace/users/internal/domain"
	"voidspace/users/internal/usecase"
	pb "voidspace/users/proto/generated/voidspace/users/proto/users/v1"
	"voidspace/users/utils/interceptor"
	"voidspace/users/utils/token"
	"voidspace/users/utils/validations"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type LoginRequest struct {
	EmailOrUsername string `validate:"required,min=3,max=50"`
	Password        string `validate:"required,min=8"`
}

type RegisterRequest struct {
	Username string `validate:"required,min=3,max=50"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}

type RefreshRequest struct {
	UserID   int32  `validate:"required"`
	Username string `validate:"required"`
}

type AuthHandler struct {
	pb.UnimplementedAuthServiceServer

	AuthUsecase          usecase.AuthUsecase
	Validator            *validator.Validate
	PrivateKey           *rsa.PrivateKey
	ContextTimeout       time.Duration
	AccessTokenDuration  time.Duration
	RefreshTokenDuration time.Duration
	Logger               *zap.Logger
}

func NewAuthHandler(
	authUsecase usecase.AuthUsecase,
	validator *validator.Validate,
	privateKey *rsa.PrivateKey,
	contextTimeout time.Duration,
	accessTokenDuration time.Duration,
	refreshTokenDuration time.Duration,
	logger *zap.Logger,
) *AuthHandler {
	return &AuthHandler{
		AuthUsecase:          authUsecase,
		Validator:            validator,
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

	loginReq := LoginRequest{
		EmailOrUsername: req.GetEmailOrUsername(),
		Password:        req.GetPassword(),
	}

	err := ah.Validator.Struct(loginReq)
	if err != nil {
		ah.Logger.Debug("Validation failed", zap.Error(err))
		return nil, status.Errorf(codes.InvalidArgument, "%s", validations.FormatValidationError(err))
	}

	user, err := ah.AuthUsecase.Login(ctx, loginReq.EmailOrUsername, loginReq.Password)
	if err != nil {
		ah.Logger.Error("Usecase error", zap.Error(err))
		switch err {
		case ctx.Err():
			return nil, status.Error(codes.DeadlineExceeded, ErrRequestTimeout)
		case domain.ErrInvalidCredentials:
			return nil, status.Error(codes.Unauthenticated, err.Error())
		default:
			return nil, status.Error(codes.Internal, ErrInternalServer)
		}
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

	registerReq := RegisterRequest{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	err := ah.Validator.Struct(registerReq)
	if err != nil {
		ah.Logger.Debug("Validation failed", zap.Error(err))
		return nil, status.Errorf(codes.InvalidArgument, "%s", validations.FormatValidationError(err))
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(registerReq.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		ah.Logger.Error("Hashing password error", zap.Error(err))
		return nil, status.Error(codes.Internal, "password hashing failed")
	}

	hashed := string(hashedPassword)

	user, err := ah.AuthUsecase.Register(ctx, registerReq.Username, registerReq.Email, hashed)
	if err != nil {
		ah.Logger.Error("Usecase error", zap.Error(err))
		switch err {
		case ctx.Err():
			return nil, status.Error(codes.DeadlineExceeded, ErrRequestTimeout)
		case domain.ErrUserExists:
			return nil, status.Error(codes.AlreadyExists, err.Error())
		default:
			return nil, status.Error(codes.Internal, ErrInternalServer)
		}
	}

	//goroutines next
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

	userID, ok := ctx.Value(interceptor.CtxKeyUserID).(int)
	if !ok {
		ah.Logger.Error("Failed to get userID from context")
		return nil, status.Error(codes.Internal, "failed to get user ID from context")
	}

	username, ok := ctx.Value(interceptor.CtxKeyUsername).(string)
	if !ok {
		ah.Logger.Error("Failed to get username from context")
		return nil, status.Error(codes.Internal, "failed to get username from context")
	}

	refreshReq := RefreshRequest{
		UserID:   int32(userID),
		Username: username,
	}
	err := ah.Validator.Struct(refreshReq)
	if err != nil {
		ah.Logger.Debug("Validation failed", zap.Error(err))
		return nil, status.Errorf(codes.InvalidArgument, "%s", validations.FormatValidationError(err))
	}

	user := &domain.User{
		ID:       int(refreshReq.UserID),
		Username: refreshReq.Username,
	}

	accessToken, err := token.CreateAccessToken(user, ah.PrivateKey, ah.AccessTokenDuration)
	if err != nil {
		ah.Logger.Error("Generate access token error", zap.Error(err))
		return nil, status.Error(codes.Internal, "token generation failed")
	}

	return &pb.AuthResponse{
		AccessToken: accessToken,
		ExpiresIn:   int32(ah.AccessTokenDuration.Seconds()),
		Message:     "Token refreshed successfully",
	}, nil
}
