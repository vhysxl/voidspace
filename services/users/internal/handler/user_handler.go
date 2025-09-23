package handler

import (
	"context"
	"time"
	"voidspace/users/internal/domain"
	pb "voidspace/users/proto/generated/users"
	errorutils "voidspace/users/utils/error"
	"voidspace/users/utils/interceptor"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer

	UserUsecase    domain.UserUsecase
	ProfileUsecase domain.ProfileUsecase
	FollowUsecase  domain.FollowUsecase
	Logger         *zap.Logger
	ContextTimeout time.Duration
}

func NewUserHandler(
	userUsecase domain.UserUsecase,
	profileUsecase domain.ProfileUsecase,
	followUsecase domain.FollowUsecase,
	timeout time.Duration,
	logger *zap.Logger,
) pb.UserServiceServer {
	return &UserHandler{
		UserUsecase:    userUsecase,
		ProfileUsecase: profileUsecase,
		FollowUsecase:  followUsecase,
		Logger:         logger,
		ContextTimeout: timeout,
	}
}

func (uh *UserHandler) GetCurrentUser(ctx context.Context, _ *emptypb.Empty) (*pb.GetUserResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, uh.ContextTimeout)
	defer cancel()

	userId, err := errorutils.GetUserIDFromContext(ctx, interceptor.CtxKeyUserID)
	if err != nil {
		return nil, errorutils.HandleAuthError(nil, uh.Logger)
	}

	existingUser, err := uh.UserUsecase.GetCurrentUser(ctx, userId)
	if err != nil {
		return nil, errorutils.HandleError(err, uh.Logger, "GetCurrentUser")
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

	user := &pb.User{
		Id:        &userId,
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

	userId, _ := ctx.Value(interceptor.CtxKeyUserID).(int32)

	existingUser, err := uh.UserUsecase.GetUser(ctx, req.GetUsername(), userId)
	if err != nil {
		return nil, errorutils.HandleError(err, uh.Logger, "GetUser")
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

	user := &pb.User{
		Id:         &existingUser.Id,
		Username:   existingUser.Username,
		Profile:    profile,
		CreatedAt:  timestamppb.New(existingUser.CreatedAt),
		IsFollowed: existingUser.IsFollowed,
	}

	return &pb.GetUserResponse{
		Message: "success get user",
		User:    user,
	}, nil
}

func (uh *UserHandler) GetUserById(ctx context.Context, req *pb.GetUserByIDRequest) (*pb.GetUserResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, uh.ContextTimeout)
	defer cancel()

	existingUser, err := uh.UserUsecase.GetUserById(ctx, req.GetUserID())
	if err != nil {
		return nil, errorutils.HandleError(err, uh.Logger, "GetUserById")
	}

	profile := &pb.Profile{
		DisplayName: existingUser.DisplayName,
		Bio:         existingUser.Bio,
		AvatarUrl:   existingUser.AvatarUrl,
		BannerUrl:   existingUser.BannerUrl,
		Location:    existingUser.Location,
		Followers:   existingUser.Followers,
		Following:   existingUser.Following,
	}

	user := &pb.User{
		Id:        &existingUser.Id,
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
		return nil, errorutils.HandleError(err, uh.Logger, "GetUsersByIds")
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

		pbUser := &pb.User{
			Id:        &u.Id,
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

func (uh *UserHandler) GetUsersFollowedById(ctx context.Context, req *pb.GetUserByIDRequest) (*pb.GetUsersFollowedResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, uh.ContextTimeout)
	defer cancel()

	userId, err := errorutils.GetUserIDFromContext(ctx, interceptor.CtxKeyUserID)
	if err != nil {
		return nil, errorutils.HandleAuthError(nil, uh.Logger)
	}

	res, err := uh.UserUsecase.GetUserFollowedById(ctx, userId)
	if err != nil {
		return nil, errorutils.HandleError(err, uh.Logger, "GetUsersFollowedById")
	}

	return &pb.GetUsersFollowedResponse{
		UserIds: res,
	}, nil
}

func (uh *UserHandler) UpdateProfile(ctx context.Context, req *pb.UpdateProfileRequest) (*pb.UserServiceResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, uh.ContextTimeout)
	defer cancel()

	userId, err := errorutils.GetUserIDFromContext(ctx, interceptor.CtxKeyUserID)
	if err != nil {
		return nil, errorutils.HandleAuthError(nil, uh.Logger)
	}

	profile := &domain.Profile{
		UserId:      userId,
		DisplayName: req.GetDisplayName(),
		Bio:         req.GetBio(),
		AvatarUrl:   req.GetAvatarUrl(),
		BannerUrl:   req.GetBannerUrl(),
		Location:    req.GetLocation(),
	}

	err = uh.ProfileUsecase.UpdateProfile(ctx, userId, profile)
	if err != nil {
		return nil, errorutils.HandleError(err, uh.Logger, "UpdateProfile")
	}

	return &pb.UserServiceResponse{
		Message: "Update profile success",
	}, nil
}

func (uh *UserHandler) DeleteUser(ctx context.Context, _ *emptypb.Empty) (*pb.UserServiceResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, uh.ContextTimeout)
	defer cancel()

	userId, err := errorutils.GetUserIDFromContext(ctx, interceptor.CtxKeyUserID)
	if err != nil {
		return nil, errorutils.HandleAuthError(nil, uh.Logger)
	}

	uh.Logger.Info("User account deletion requested", zap.Int32("user_id", userId))

	err = uh.UserUsecase.DeleteUser(ctx, userId)
	if err != nil {
		return nil, errorutils.HandleError(err, uh.Logger, "DeleteUser")
	}

	return &pb.UserServiceResponse{
		Message: "Delete user success",
	}, nil
}

func (uh *UserHandler) Follow(ctx context.Context, req *pb.FollowRequest) (*pb.UserServiceResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, uh.ContextTimeout)
	defer cancel()

	userId, err := errorutils.GetUserIDFromContext(ctx, interceptor.CtxKeyUserID)
	if err != nil {
		return nil, errorutils.HandleAuthError(nil, uh.Logger)
	}

	err = uh.FollowUsecase.Follow(ctx, userId, req.GetUsername())
	if err != nil {
		return nil, errorutils.HandleError(err, uh.Logger, "Follow")
	}

	return &pb.UserServiceResponse{
		Message: "Follow user success",
	}, nil
}

func (uh *UserHandler) Unfollow(ctx context.Context, req *pb.FollowRequest) (*pb.UserServiceResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, uh.ContextTimeout)
	defer cancel()

	userId, err := errorutils.GetUserIDFromContext(ctx, interceptor.CtxKeyUserID)
	if err != nil {
		return nil, errorutils.HandleAuthError(nil, uh.Logger)
	}

	err = uh.FollowUsecase.Unfollow(ctx, userId, req.GetUsername())
	if err != nil {
		return nil, errorutils.HandleError(err, uh.Logger, "Unfollow")
	}

	return &pb.UserServiceResponse{
		Message: "Unfollow user success",
	}, nil
}
