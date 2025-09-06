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

func (fs *FeedService) GetGlobalFeed(ctx context.Context, req *models.GetGlobalFeedReq) (*models.GetFeedResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, fs.ContextTimeout)
	defer cancel()

	// 1. Call PostService to fetch the global feed posts.
	//    This returns only post data (id, content, user_id, etc.),
	//    but it does not include the full user info (username, profile).
	postRes, err := fs.PostClient.GetGlobalFeed(ctx, &postpb.GetGlobalFeedRequest{
		CursorID: int32(req.CursorUserID),
		Cursor:   timestamppb.New(req.Cursor),
	})
	if err != nil {
		fs.Logger.Error("failed to call PostService.GetGlobalFeed", zap.Error(err))
		return nil, err
	}

	// Guard if post is empty
	if len(postRes.Posts) == 0 {
		return &models.GetFeedResponse{
			Posts:   nil,
			HasMore: postRes.GetHasMore(),
		}, nil
	}

	// 2. Collect unique user IDs from the posts so we know which users we need
	//    to enrich the feed with. Using a map as a set avoids duplicates.
	userIDSet := make(map[int32]struct{})
	for _, post := range postRes.GetPosts() {
		userIDSet[post.GetUserId()] = struct{}{}
	}

	// 3. Convert the set into a slice, since gRPC request requires a slice.
	userIDs := make([]int32, 0, len(userIDSet))
	for id := range userIDSet {
		userIDs = append(userIDs, id)
	}

	// 4. Call UserService in batch to fetch all user details for the collected IDs.
	//    This is the critical optimization: instead of calling UserService N times
	//    (once per post), we only call it ONCE with all needed IDs.
	userRes, err := fs.UserClient.GetUsersByIds(ctx, &userpb.GetUserByUserIDsRequest{
		UserID: userIDs,
	})
	if err != nil {
		fs.Logger.Error("failed to call UserService.GetUsersByIds", zap.Error(err))
		return nil, err
	}

	// 5. Build a map[user_id]User for quick lookup when merging user info into posts.
	userMap := make(map[int32]*userpb.User)
	for _, u := range userRes.GetUsers() {
		userMap[u.GetId()] = u
	}

	// 6. Merge posts with user data.
	//    For each post, we lookup its author in userMap.
	//    If found, we map the gRPC User protobuf into our domain model (models.User).
	posts := make([]models.Post, 0, len(postRes.GetPosts()))
	for _, post := range postRes.GetPosts() {
		u := userMap[post.GetUserId()]

		var user *models.User
		if u != nil {
			// Map profile fields if present.
			var profile models.Profile
			if u.GetProfile() != nil {
				profile = models.Profile{
					DisplayName: u.GetProfile().GetDisplayName(),
					Bio:         u.GetProfile().GetBio(),
					AvatarURL:   u.GetProfile().GetAvatarUrl(),
					BannerURL:   u.GetProfile().GetBannerUrl(),
					Location:    u.GetProfile().GetLocation(),
					Followers:   u.GetProfile().GetFollowers(),
					Following:   u.GetProfile().GetFollowing(),
				}
			}

			// Construct the user model with profile attached.
			user = &models.User{
				ID:        u.GetId(),
				Username:  u.GetUsername(),
				Profile:   profile,
				CreatedAt: u.GetCreatedAt().AsTime(),
			}
		}

		// Combine post data with enriched author info.
		posts = append(posts, models.Post{
			ID:         int(post.GetId()),
			Content:    post.GetContent(),
			UserID:     int(post.GetUserId()),
			Author:     user, // This is the enriched user info
			PostImages: post.GetPostImages(),
			LikesCount: int(post.GetLikesCount()),
			CreatedAt:  post.GetCreatedAt().AsTime(),
			UpdatedAt:  post.GetUpdatedAt().AsTime(),
		})
	}

	// 7. Return aggregated response containing posts + author details.
	//    At this point, the response is "fully enriched" with user information
	//    so the frontend doesn't need to make additional calls.
	return &models.GetFeedResponse{
		Posts:   posts,
		HasMore: postRes.GetHasMore(),
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
