package service

import (
	"context"
	"time"
	"voidspaceGateway/internal/models"
	postpb "voidspaceGateway/proto/generated/posts"
	userpb "voidspaceGateway/proto/generated/users"

	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
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

func (fs *FeedService) GetGlobalFeed(ctx context.Context, req models.GetGlobalFeed) (*models.GetFeedResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, fs.ContextTimeout)
	defer cancel()

	res, err := fs.PostClient.GetGlobalFeed(ctx, &postpb.GetGlobalFeedRequest{
		CursorID: int32(req.CursorUserID),
		Cursor:   timestamppb.New(req.Cursor),
	})
	if err != nil {
		fs.Logger.Error("failed to call PostService.GetGlobalFeed", zap.Error(err))
		return nil, err
	}

	posts := make([]models.Post, 0, len(res.GetPosts()))
	for _, post := range res.GetPosts() {
		posts = append(posts, models.Post{
			ID:         int(post.GetId()),
			Content:    post.GetContent(),
			UserID:     int(post.GetUserId()),
			PostImages: post.GetPostImages(),
			LikesCount: int(post.GetLikesCount()),
			CreatedAt:  post.GetCreatedAt().AsTime(),
			UpdatedAt:  post.GetUpdatedAt().AsTime(),
		})
	}
	return &models.GetFeedResponse{
		Posts:   posts,
		HasMore: res.GetHasMore(),
	}, nil
}

func (fs *FeedService) GetFollowFeed(ctx context.Context, username string, userID string, req models.GetFollowFeedReq) (*models.GetFeedResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, fs.ContextTimeout)
	defer cancel()

	md := metadata.New(map[string]string{
		"user_id":  userID,
		"username": username,
	})

	ctx = metadata.NewOutgoingContext(ctx, md)

	data := &postpb.GetFeedByUserIDsRequest{
		UserIds:  req.UserIDs,
		Cursor:   timestamppb.New(req.Cursor),
		CursorID: int32(req.CursorUserID),
	}

	res, err := fs.PostClient.GetFeedByUserIDs(ctx, data)
	if err != nil {
		fs.Logger.Error("failed to call PostService.GetGlobalFeed", zap.Error(err))
		return nil, err
	}

	posts := make([]models.Post, 0, len(res.GetPosts()))
	for _, post := range res.GetPosts() {
		posts = append(posts, models.Post{
			ID:         int(post.GetId()),
			Content:    post.GetContent(),
			UserID:     int(post.GetUserId()),
			PostImages: post.GetPostImages(),
			LikesCount: int(post.GetLikesCount()),
			CreatedAt:  post.GetCreatedAt().AsTime(),
			UpdatedAt:  post.GetUpdatedAt().AsTime(),
		})
	}
	return &models.GetFeedResponse{
		Posts:   posts,
		HasMore: res.GetHasMore(),
	}, nil
}
