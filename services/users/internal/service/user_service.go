package handler

import (
	"context"

	"time"
	"voidspace/users/internal/domain"
	"voidspace/users/internal/usecase"
	pb "voidspace/users/proto/generated/users"
	"voidspace/users/utils/interceptor"

	"voidspace/users/utils/validations"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GetUserReq struct {
	Username string `validate:"required,min=3,max=32,alphanum"`
}

// UpdateUserReq represents a partial update payload for updating user profile.
// All fields are optional and will only be updated if provided in the request.
// Use pointers to differentiate between "field not provided" (nil) and "empty value" ("").
type UpdateUserReq struct {
	DisplayName string `validate:"omitempty,max=50"`
	Bio         string `validate:"omitempty,max=160"`
	AvatarUrl   string `validate:"omitempty"`
	BannerUrl   string `validate:"omitempty"`
	Location    string `validate:"omitempty,max=100"`
}

type UserHandler struct {
	pb.UnimplementedUserServiceServer

	UserUsecase    usecase.UserUsecase
	ProfileUsecase usecase.ProfileUsecase
	FollowUsecase  usecase.FollowUsecase
	Validator      *validator.Validate
	Logger         *zap.Logger
	ContextTimeout time.Duration
}

func NewUserHandler(
	userUsecase usecase.UserUsecase,
	profileUsecase usecase.ProfileUsecase,
	followUsecase usecase.FollowUsecase,
	validator *validator.Validate,
	timeout time.Duration,
	logger *zap.Logger,
) *UserHandler {
	return &UserHandler{
		UserUsecase:    userUsecase,
		ProfileUsecase: profileUsecase,
		FollowUsecase:  followUsecase,
		Validator:      validator,
		Logger:         logger,
		ContextTimeout: timeout,
	}
}

func (uh *UserHandler) GetCurrentUser(ctx context.Context, _ *emptypb.Empty) (*pb.GetUserResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, uh.ContextTimeout)
	defer cancel()

	userID, ok := ctx.Value(interceptor.CtxKeyUserID).(int)
	if !ok {
		uh.Logger.Error(ErrFailedGetUserID)
		return nil, status.Error(codes.Internal, ErrFailedGetUserID)
	}

	existingUser, err := uh.UserUsecase.GetCurrentUser(ctx, int(userID))
	if err != nil {
		uh.Logger.Error(ErrUsecase, zap.Error(err))
		switch err {
		case ctx.Err():
			return nil, status.Error(codes.DeadlineExceeded, ErrRequestTimeout)
		case domain.ErrUserNotFound:
			return nil, status.Error(codes.NotFound, err.Error())
		default:
			return nil, status.Error(codes.Internal, ErrInternalServer)
		}
	}

	profile := &pb.Profile{
		DisplayName: existingUser.DisplayName,
		Bio:         existingUser.Bio,
		AvatarUrl:   existingUser.AvatarUrl,
		BannerUrl:   existingUser.BannerUrl,
		Location:    existingUser.Location,
		Followers:   int32(existingUser.Followers),
		Following:   int32(existingUser.Following),
	}

	userIDInt32 := int32(userID)
	user := &pb.User{
		Id:        &userIDInt32,
		Username:  existingUser.Username,
		Profile:   profile,
		CreatedAt: timestamppb.New(existingUser.CreatedAt),
	}
	return &pb.GetUserResponse{
		Message: "success get current user",
		User:    user,
	}, nil
}

func (uh *UserHandler) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, uh.ContextTimeout)
	defer cancel()

	userReq := GetUserReq{
		Username: req.GetUsername(),
	}

	err := uh.Validator.Struct(userReq)
	if err != nil {
		uh.Logger.Debug(ErrValidation, zap.Error(err))
		return nil, status.Errorf(codes.InvalidArgument, "%s", validations.FormatValidationError(err))
	}

	existingUser, err := uh.UserUsecase.GetUser(ctx, userReq.Username)
	if err != nil {
		uh.Logger.Error(ErrUsecase, zap.Error(err))
		switch err {
		case ctx.Err():
			return nil, status.Error(codes.DeadlineExceeded, ErrRequestTimeout)
		case domain.ErrUserNotFound:
			return nil, status.Error(codes.NotFound, err.Error())
		default:
			return nil, status.Error(codes.Internal, ErrInternalServer)
		}
	}

	profile := &pb.Profile{
		DisplayName: existingUser.DisplayName,
		Bio:         existingUser.Bio,
		AvatarUrl:   existingUser.AvatarUrl,
		BannerUrl:   existingUser.BannerUrl,
		Location:    existingUser.Location,
		Followers:   int32(existingUser.Followers),
		Following:   int32(existingUser.Following),
	}

	userIDInt32 := int32(existingUser.ID)
	user := &pb.User{
		Id:        &userIDInt32,
		Username:  existingUser.Username,
		Profile:   profile,
		CreatedAt: timestamppb.New(existingUser.CreatedAt),
	}

	return &pb.GetUserResponse{
		Message: "success get user",
		User:    user,
	}, nil
}

func (uh *UserHandler) GetUsersByIds(ctx context.Context, req *pb.GetUserByUserIDsRequest) (*pb.GetUsersResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, uh.ContextTimeout)
	defer cancel()

	users, err := uh.UserUsecase.GetUserByIds(ctx, req.UserID)
	if err != nil {
		uh.Logger.Error(ErrUsecase, zap.Error(err))
		switch err {
		case ctx.Err():
			return nil, status.Error(codes.DeadlineExceeded, ErrRequestTimeout)
		case domain.ErrUserNotFound:
			return nil, status.Error(codes.NotFound, err.Error())
		default:
			return nil, status.Error(codes.Internal, ErrInternalServer)
		}
	}

	pbUsers := make([]*pb.User, 0, len(users))
	for _, u := range users {
		profile := &pb.Profile{
			DisplayName: u.DisplayName,
			Bio:         u.Bio,
			AvatarUrl:   u.AvatarUrl,
			BannerUrl:   u.BannerUrl,
			Location:    u.Location,
			Followers:   int32(u.Followers),
			Following:   int32(u.Following),
		}

		userIDInt32 := int32(u.ID)
		pbUser := &pb.User{
			Id:        &userIDInt32,
			Username:  u.Username,
			Profile:   profile,
			CreatedAt: timestamppb.New(u.CreatedAt),
		}
		pbUsers = append(pbUsers, pbUser)
	}

	return &pb.GetUsersResponse{
		Message: "success get users",
		Users:   pbUsers,
	}, nil
}

func (uh *UserHandler) UpdateProfile(ctx context.Context, req *pb.UpdateProfileRequest) (*pb.UserServiceResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, uh.ContextTimeout)
	defer cancel()

	userID, ok := ctx.Value(interceptor.CtxKeyUserID).(int)
	if !ok {
		uh.Logger.Error(ErrFailedGetUserID)
		return nil, status.Error(codes.Internal, ErrFailedGetUserID)
	}

	prepData := &UpdateUserReq{
		DisplayName: req.GetDisplayName(),
		Bio:         req.GetBio(),
		AvatarUrl:   req.GetAvatarUrl(),
		BannerUrl:   req.GetBannerUrl(),
		Location:    req.GetLocation(),
	}

	err := uh.Validator.Struct(prepData)
	if err != nil {
		uh.Logger.Debug(ErrValidation, zap.Error(err))
		return nil, status.Errorf(codes.InvalidArgument, "%s", validations.FormatValidationError(err))
	}

	profile := &domain.Profile{
		UserID:      int(userID),
		DisplayName: prepData.DisplayName,
		Bio:         prepData.Bio,
		AvatarUrl:   prepData.AvatarUrl,
		BannerUrl:   prepData.BannerUrl,
		Location:    prepData.Location,
	}

	err = uh.ProfileUsecase.UpdateProfile(ctx, userID, profile)
	if err != nil {
		uh.Logger.Error(ErrUsecase, zap.Error(err))
		switch err {
		case ctx.Err():
			return nil, status.Error(codes.DeadlineExceeded, ErrRequestTimeout)
		case domain.ErrUserNotFound:
			return nil, status.Error(codes.NotFound, err.Error())
		default:
			return nil, status.Error(codes.Internal, ErrInternalServer)
		}
	}

	return &pb.UserServiceResponse{
		Message: "Update profile success",
	}, nil
}

func (uh *UserHandler) DeleteUser(ctx context.Context, _ *emptypb.Empty) (*pb.UserServiceResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, uh.ContextTimeout)
	defer cancel()

	userID, ok := ctx.Value(interceptor.CtxKeyUserID).(int)
	if !ok {
		uh.Logger.Error(ErrFailedGetUserID)
		return nil, status.Error(codes.Internal, ErrFailedGetUserID)
	}

	uh.Logger.Info("User account deletion requested", zap.Int("user_id", userID))

	err := uh.UserUsecase.DeleteUser(ctx, userID)
	if err != nil {
		uh.Logger.Error(ErrUsecase, zap.Error(err))
		switch err {
		case ctx.Err():
			return nil, status.Error(codes.DeadlineExceeded, ErrRequestTimeout)
		case domain.ErrUserNotFound:
			return nil, status.Error(codes.NotFound, err.Error())
		default:
			return nil, status.Error(codes.Internal, ErrInternalServer)
		}
	}

	return &pb.UserServiceResponse{
		Message: "Delete user success",
	}, nil
}

func (uh *UserHandler) Follow(ctx context.Context, req *pb.FollowRequest) (*pb.UserServiceResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, uh.ContextTimeout)
	defer cancel()

	userID, ok := ctx.Value(interceptor.CtxKeyUserID).(int)
	if !ok {
		uh.Logger.Error(ErrFailedGetUserID)
		return nil, status.Error(codes.Internal, ErrFailedGetUserID)
	}

	usernameTarget := &GetUserReq{
		Username: req.GetUsername(),
	}

	err := uh.Validator.Struct(usernameTarget)
	if err != nil {
		uh.Logger.Debug(ErrValidation, zap.Error(err))
		return nil, status.Errorf(codes.InvalidArgument, "%s", validations.FormatValidationError(err))
	}

	err = uh.FollowUsecase.Follow(ctx, userID, usernameTarget.Username)
	if err != nil {
		uh.Logger.Error(ErrUsecase, zap.Error(err))
		switch err {
		case ctx.Err():
			return nil, status.Error(codes.DeadlineExceeded, ErrRequestTimeout)
		case domain.ErrAlreadyFollow:
			return nil, status.Error(codes.AlreadyExists, err.Error())
		case domain.ErrSelfFollow:
			return nil, status.Error(codes.InvalidArgument, err.Error())
		case domain.ErrUserNotFound:
			return nil, status.Error(codes.NotFound, err.Error())
		default:
			return nil, status.Error(codes.Internal, ErrInternalServer)
		}
	}

	return &pb.UserServiceResponse{
		Message: "Follow user success",
	}, nil
}

func (uh *UserHandler) Unfollow(ctx context.Context, req *pb.FollowRequest) (*pb.UserServiceResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, uh.ContextTimeout)
	defer cancel()

	userID, ok := ctx.Value(interceptor.CtxKeyUserID).(int)
	if !ok {
		uh.Logger.Error("Failed to get userID from context")
		return nil, status.Error(codes.Internal, "failed to get user ID from context")
	}

	usernameTarget := &GetUserReq{
		Username: req.GetUsername(),
	}

	err := uh.Validator.Struct(usernameTarget)
	if err != nil {
		uh.Logger.Debug(ErrValidation, zap.Error(err))
		return nil, status.Errorf(codes.InvalidArgument, "%s", validations.FormatValidationError(err))
	}

	err = uh.FollowUsecase.Unfollow(ctx, userID, usernameTarget.Username)
	if err != nil {
		uh.Logger.Error(ErrUsecase, zap.Error(err))
		switch err {
		case ctx.Err():
			return nil, status.Error(codes.DeadlineExceeded, ErrRequestTimeout)
		case domain.ErrUserNotFound:
			return nil, status.Error(codes.NotFound, err.Error())
		default:
			return nil, status.Error(codes.Internal, ErrInternalServer)
		}
	}

	return &pb.UserServiceResponse{
		Message: "Unfollow user success",
	}, nil
}
