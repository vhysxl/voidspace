package service

import (
	"context"
	"errors"
	"time"
	"voidspaceGateway/internal/models"
	commentpb "voidspaceGateway/proto/generated/comments"
	postpb "voidspaceGateway/proto/generated/posts"
	userpb "voidspaceGateway/proto/generated/users"
	temporal_constants "voidspaceGateway/temporal/constants"
	temporal_dto "voidspaceGateway/temporal/dto"
	"voidspaceGateway/utils"

	"go.temporal.io/sdk/client"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserService struct {
	ContextTimeout  time.Duration
	Logger          *zap.Logger
	UserClient      userpb.UserServiceClient
	PostClient      postpb.PostServiceClient
	CommentClient   commentpb.CommentServiceClient
	TemporalClient  client.Client
	TemporalService string
}

func NewUserService(
	timeout time.Duration,
	logger *zap.Logger,
	userClient userpb.UserServiceClient,
	postClient postpb.PostServiceClient,
	commentClient commentpb.CommentServiceClient,
	temporalClient client.Client,
	temporalService string,
) *UserService {
	return &UserService{
		ContextTimeout:  timeout,
		Logger:          logger,
		UserClient:      userClient,
		PostClient:      postClient,
		CommentClient:   commentClient,
		TemporalClient:  temporalClient,
		TemporalService: temporalService,
	}
}

func (us *UserService) GetCurrentUser(ctx context.Context, userID string, username string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, us.ContextTimeout)
	defer cancel()

	md := utils.MetaDataHandler(userID, username)

	ctx = metadata.NewOutgoingContext(ctx, md)

	res, err := us.UserClient.GetCurrentUser(ctx, &emptypb.Empty{})
	if err != nil {
		us.Logger.Error("failed to call UserService.GetCurrentUser", zap.Error(err))
		return nil, err
	}

	return utils.UserMapper(res), nil
}

func (us *UserService) GetUser(ctx context.Context, username string, userID string, usernameRequester string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, us.ContextTimeout)
	defer cancel()

	md := utils.MetaDataHandler(userID, username)

	ctx = metadata.NewOutgoingContext(ctx, md)

	res, err := us.UserClient.GetUser(ctx, &userpb.GetUserRequest{
		Username: username,
	})
	if err != nil {
		us.Logger.Error("failed to call UserService.GetUser", zap.Error(err))
		return nil, err
	}

	return utils.UserMapper(res), nil
}

func (us *UserService) UpdateProfile(ctx context.Context, userID string, username string, req *models.UpdateProfileRequest) error {
	ctx, cancel := context.WithTimeout(ctx, us.ContextTimeout)
	defer cancel()

	md := utils.MetaDataHandler(userID, username)

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

// todo: distributed trx
func (us *UserService) DeleteUser(ctx context.Context, userID string, username string) error {
	ctx, cancel := context.WithTimeout(ctx, us.ContextTimeout)
	defer cancel()

	param := temporal_dto.DeleteUserWorkflowParam{
		UserID:   userID,
		Username: username,
	}

	run, err := us.TemporalClient.ExecuteWorkflow(
		ctx,
		client.StartWorkflowOptions{
			ID:        "delete-user-" + userID,
			TaskQueue: us.TemporalService,
		},
		temporal_constants.DeleteUserWorkflowName,
		param,
	)
	if err != nil {
		us.Logger.Error("failed to execute workflow", zap.Error(err))
		return err
	}

	var res temporal_dto.DeleteUserWorkflowResult
	if err := run.Get(ctx, &res); err != nil {
		us.Logger.Error("workflow failed", zap.Error(err))
		return err
	}

	if !res.Success {
		return errors.New("delete user workflow failed")
	}

	return nil
}
func (us *UserService) Follow(ctx context.Context, userID string, username string, targetUsername string) error {
	ctx, cancel := context.WithTimeout(ctx, us.ContextTimeout)
	defer cancel()

	md := utils.MetaDataHandler(userID, username)

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

	md := utils.MetaDataHandler(userID, username)

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
