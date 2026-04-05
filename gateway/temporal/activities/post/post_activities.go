package post

import (
	"time"

	commentpb "voidspaceGateway/proto/generated/comments/v1"
	postpb "voidspaceGateway/proto/generated/posts/v1"

	"go.uber.org/zap"
)

type PostActivities struct {
	ContextTimeout time.Duration
	Logger         *zap.Logger
	PostClient     postpb.PostServiceClient
	CommentClient  commentpb.CommentServiceClient
}

func NewPostActivities(
	contextTimeout time.Duration,
	logger *zap.Logger,
	postClient postpb.PostServiceClient,
	commentClient commentpb.CommentServiceClient,
) *PostActivities {
	return &PostActivities{
		ContextTimeout: contextTimeout,
		Logger:         logger,
		PostClient:     postClient,
		CommentClient:  commentClient,
	}
}
