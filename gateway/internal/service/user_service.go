package service

import (
	"context"
	"time"
	"voidspaceGateway/internal/models"
	userpb "voidspaceGateway/proto/generated/users"

	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
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

func (us *UserService) GetCurrentUser(ctx context.Context, userID string, username string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, us.ContextTimeout)
	defer cancel()

	md := metadata.New(map[string]string{
		"user_id":  userID,
		"username": username,
	})

	ctx = metadata.NewOutgoingContext(ctx, md)

	res, err := us.UserClient.GetCurrentUser(ctx, &emptypb.Empty{})
	if err != nil {
		us.Logger.Error("failed to call UserService.GetCurrentUser", zap.Error(err))
		return nil, err
	}
	user := res.User
	profile := &models.Profile{}

	if userProfile := user.GetProfile(); userProfile != nil {
		profile.DisplayName = userProfile.GetDisplayName()
		profile.Bio = userProfile.GetBio()
		profile.AvatarURL = userProfile.GetAvatarUrl()
		profile.Location = userProfile.GetLocation()
		profile.BannerURL = userProfile.GetBannerUrl()
		profile.Followers = userProfile.GetFollowers()
		profile.Following = userProfile.GetFollowing()
	}

	return &models.User{
		ID:        res.User.GetId(),
		Username:  res.User.GetUsername(),
		Profile:   *profile,
		CreatedAt: res.User.GetCreatedAt().AsTime(),
	}, nil
}

func (us *UserService) GetUser(ctx context.Context, username string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, us.ContextTimeout)
	defer cancel()

	res, err := us.UserClient.GetUser(ctx, &userpb.GetUserRequest{
		Username: username,
	})
	if err != nil {
		us.Logger.Error("failed to call UserService.GetUser", zap.Error(err))
		return nil, err
	}
	user := res.User
	profile := &models.Profile{}

	if userProfile := user.GetProfile(); userProfile != nil {
		profile.DisplayName = userProfile.GetDisplayName()
		profile.Bio = userProfile.GetBio()
		profile.AvatarURL = userProfile.GetAvatarUrl()
		profile.Location = userProfile.GetLocation()
		profile.BannerURL = userProfile.GetBannerUrl()
		profile.Followers = userProfile.GetFollowers()
		profile.Following = userProfile.GetFollowing()
	}

	return &models.User{
		Username:  res.User.GetUsername(),
		Profile:   *profile,
		CreatedAt: res.User.GetCreatedAt().AsTime(),
	}, nil
}

func (us *UserService) UpdateProfile(ctx context.Context, userID string, username string, req *models.UpdateProfileRequest) error {
	ctx, cancel := context.WithTimeout(ctx, us.ContextTimeout)
	defer cancel()

	md := metadata.New(map[string]string{
		"user_id":  userID,
		"username": username,
	})

	ctx = metadata.NewOutgoingContext(ctx, md)

	_, err := us.UserClient.UpdateProfile(ctx, &userpb.UpdateProfileRequest{
		DisplayName: &req.DisplayName,
		Bio:         &req.Bio,
		AvatarUrl:   &req.AvatarURL,
		BannerUrl:   &req.BannerURL,
		Location:    &req.Location,
	})
	if err != nil {
		us.Logger.Error("failed to call UserService.UpdateProfile", zap.Error(err))
		return err
	}

	return err
}

func (us *UserService) DeleteUser(ctx context.Context, userID string, username string) error {
	ctx, cancel := context.WithTimeout(ctx, us.ContextTimeout)
	defer cancel()

	md := metadata.New(map[string]string{
		"user_id":  userID,
		"username": username,
	})

	ctx = metadata.NewOutgoingContext(ctx, md)

	_, err := us.UserClient.DeleteUser(ctx, &emptypb.Empty{})
	if err != nil {
		us.Logger.Error("failed to call UserService.DeleteUser", zap.Error(err))
		return err
	}

	return err
}

func (us *UserService) Follow(ctx context.Context, userID string, username string, targetUsername string) error {
	ctx, cancel := context.WithTimeout(ctx, us.ContextTimeout)
	defer cancel()

	md := metadata.New(map[string]string{
		"user_id":  userID,
		"username": username,
	})

	ctx = metadata.NewOutgoingContext(ctx, md)

	_, err := us.UserClient.Follow(ctx, &userpb.FollowRequest{
		Username: targetUsername,
	})
	if err != nil {
		us.Logger.Error("failed to call UserService.Follow", zap.Error(err))
		return err
	}

	return nil
}

func (us *UserService) Unfollow(ctx context.Context, userID string, username string, targetUsername string) error {
	ctx, cancel := context.WithTimeout(ctx, us.ContextTimeout)
	defer cancel()

	md := metadata.New(map[string]string{
		"user_id":  userID,
		"username": username,
	})

	ctx = metadata.NewOutgoingContext(ctx, md)

	_, err := us.UserClient.Unfollow(ctx, &userpb.FollowRequest{
		Username: targetUsername,
	})
	if err != nil {
		us.Logger.Error("failed to call UserService.Unfollow", zap.Error(err))
		return err
	}

	return nil
}
