package user

import (
	"time"

	commentpb "voidspaceGateway/proto/generated/comments/v1"
	postpb "voidspaceGateway/proto/generated/posts/v1"
	userpb "voidspaceGateway/proto/generated/users/v1"

	"go.uber.org/zap"
)

type UserActivities struct {
	ContextTimeout time.Duration
	Logger         *zap.Logger
	UserClient     userpb.UserServiceClient
	PostClient     postpb.PostServiceClient
	CommentClient  commentpb.CommentServiceClient
}

func NewUserActivities(
	contextTimeout time.Duration,
	logger *zap.Logger,
	userClient userpb.UserServiceClient,
	postClient postpb.PostServiceClient,
	commentClient commentpb.CommentServiceClient,
) *UserActivities {
	return &UserActivities{
		ContextTimeout: contextTimeout,
		Logger:         logger,
		UserClient:     userClient,
		PostClient:     postClient,
		CommentClient:  commentClient,
	}
}
