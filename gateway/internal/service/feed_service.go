package service

import (
	"context"
	"time"
	"voidspaceGateway/internal/models"
	commentpb "voidspaceGateway/proto/generated/comments"
	postpb "voidspaceGateway/proto/generated/posts"
	userpb "voidspaceGateway/proto/generated/users"
	"voidspaceGateway/utils"

	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// TODO: Extract user enrichment logic to reduce duplication between GetGlobalFeed and GetFollowFeed
type FeedService struct {
	ContextTimeout time.Duration
	Logger         *zap.Logger
	PostClient     postpb.PostServiceClient
	UserClient     userpb.UserServiceClient
	CommentClient  commentpb.CommentServiceClient
}

func NewFeedService(timeout time.Duration, logger *zap.Logger, postClient postpb.PostServiceClient, userClient userpb.UserServiceClient, commentClient commentpb.CommentServiceClient) *FeedService {
	return &FeedService{
		ContextTimeout: timeout,
		Logger:         logger,
		PostClient:     postClient,
		UserClient:     userClient,
		CommentClient:  commentClient,
	}
}

func (fs *FeedService) GetGlobalFeed(ctx context.Context, req *models.GetGlobalFeedRequest, userID string, username string) (*models.GetFeedResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, fs.ContextTimeout)
	defer cancel()

	md := utils.MetaDataHandler(userID, username)
	ctx = metadata.NewOutgoingContext(ctx, md)

	//    Call PostService to fetch the global feed posts.
	//    This returns only post data (id, content, user_id, etc.),
	//    but it does not include the full user info (username, profile).
	postRes, err := fs.PostClient.GetGlobalFeed(ctx, &postpb.GetGlobalFeedRequest{
		CursorID: int32(req.CursorID),
		Cursor:   timestamppb.New(req.Cursor),
	})
	if err != nil {
		fs.Logger.Error("failed to call PostService.GetGlobalFeed", zap.Error(err))
		return nil, err
	}

	posts, err := utils.EnrichPosts(ctx, postRes.GetPosts(), fs.CommentClient, fs.UserClient, fs.Logger)
	if err != nil {
		return nil, err
	}

	if userID != "" {
		// create map for lookup
		followedUsers := make(map[int32]bool)
		folRes, err := fs.UserClient.GetUsersFollowedById(ctx, &userpb.GetUserByIDRequest{})
		if err == nil {
			for _, userID := range folRes.UserIds {
				followedUsers[userID] = true
			}
		}

		for _, post := range posts {
			post.Author.IsFollowed = followedUsers[int32(post.Author.ID)]
		}
	}

	return &models.GetFeedResponse{
		Posts:   posts,
		HasMore: postRes.GetHasMore(),
	}, nil

}

func (fs *FeedService) GetFollowFeed(ctx context.Context, userID string, username string, req *models.GetFollowFeedRequest) (*models.GetFeedResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, fs.ContextTimeout)
	defer cancel()

	md := utils.MetaDataHandler(userID, username)
	ctx = metadata.NewOutgoingContext(ctx, md)

	folRes, err := fs.UserClient.GetUsersFollowedById(ctx, &userpb.GetUserByIDRequest{})
	if err != nil {
		fs.Logger.Error("failed to call UserService.GetUserFollowedById", zap.Error(err))
		return nil, err
	}

	data := &postpb.GetFeedByUserIDsRequest{
		UserIds:  folRes.UserIds,
		Cursor:   timestamppb.New(req.Cursor),
		CursorID: int32(req.CursorID),
	}

	postRes, err := fs.PostClient.GetFeedByUserIDs(ctx, data)
	if err != nil {
		fs.Logger.Error("failed to call PostService.GetFeedByUserIDs", zap.Error(err))
		return nil, err
	}

	posts, err := utils.EnrichPosts(ctx, postRes.GetPosts(), fs.CommentClient, fs.UserClient, fs.Logger)
	if err != nil {
		return nil, err
	}

	for _, post := range posts {
		post.Author.IsFollowed = true
	}

	return &models.GetFeedResponse{
		Posts:   posts,
		HasMore: postRes.GetHasMore(),
	}, nil
}
