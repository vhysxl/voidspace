package service

import (
	"context"
	"time"
	"voidspaceGateway/internal/models"
	postpb "voidspaceGateway/proto/generated/posts"
	userpb "voidspaceGateway/proto/generated/users"

	"go.uber.org/zap"
)

type FeedService struct {
	ContextTimeout time.Duration
	Logger         *zap.Logger
	PostClient     postpb.PostServiceClient
	UserClient     userpb.UserServiceClient
}

func NewFeedService(timeout time.Duration, logger *zap.Logger, postClient postpb.PostServiceClient, userClient userpb.UserServiceClient) *FeedService {
	return &FeedService{
		ContextTimeout: timeout,
		Logger:         logger,
		PostClient:     postClient,
		UserClient:     userClient,
	}
}

func (fs *FeedService) GetGlobalFeed(ctx context.Context) (models.GetFeedResponse, error) {
	panic("not implemented")
}

func (fs *FeedService) GetFollowFeed(ctx context.Context) (models.GetFeedResponse, error) {
	panic("not implemented")
}
