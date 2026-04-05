package comment

import (
	"time"

	commentpb "voidspaceGateway/proto/generated/comments/v1"
	postpb "voidspaceGateway/proto/generated/posts/v1"
	userpb "voidspaceGateway/proto/generated/users/v1"

	"go.uber.org/zap"
)

type CommentService struct {
	ContextTimeout time.Duration
	Logger         *zap.Logger
	PostClient     postpb.PostServiceClient
	CommentClient  commentpb.CommentServiceClient
	UserClient     userpb.UserServiceClient
}

func NewCommentService(
	contextTimeout time.Duration,
	logger *zap.Logger,
	postClient postpb.PostServiceClient,
	commentClient commentpb.CommentServiceClient,
	userClient userpb.UserServiceClient,
) *CommentService {
	return &CommentService{
		ContextTimeout: contextTimeout,
		Logger:         logger,
		PostClient:     postClient,
		CommentClient:  commentClient,
		UserClient:     userClient,
	}
}
