package post

import (
	"time"
	commentpb "voidspaceGateway/proto/generated/comments/v1"
	postpb "voidspaceGateway/proto/generated/posts/v1"
	userpb "voidspaceGateway/proto/generated/users/v1"
	"go.uber.org/zap"
)

type PostService struct {
	ContextTimeout time.Duration
	Logger         *zap.Logger
	UserClient     userpb.UserServiceClient
	PostClient     postpb.PostServiceClient
	CommentClient  commentpb.CommentServiceClient
}

func NewPostService(
	contextTimeout time.Duration,
	logger *zap.Logger,
	userClient userpb.UserServiceClient,
	postClient postpb.PostServiceClient,
	commentClient commentpb.CommentServiceClient,
) *PostService {
	return &PostService{
		ContextTimeout: contextTimeout,
		Logger:         logger,
		UserClient:     userClient,
		PostClient:     postClient,
		CommentClient:  commentClient,
	}
}
