package post

import (
	"context"
	"strconv"
	"voidspaceGateway/internal/models"
	postpb "voidspaceGateway/proto/generated/posts/v1"
	userpb "voidspaceGateway/proto/generated/users/v1"
	"voidspaceGateway/utils"

	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

func (ps *PostService) GetFollowingFeed(ctx context.Context, req *postpb.GetFollowingFeedRequest, reqUserID string, reqUsername string) (*models.GetFeedResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, ps.ContextTimeout)
	defer cancel()

	md := utils.MetaDataHandler(reqUserID, reqUsername)
	ctx = metadata.NewOutgoingContext(ctx, md)

	userID, err := strconv.ParseInt(reqUserID, 10, 64)
	if err != nil {
		ps.Logger.Error("failed to parse reqUserID", zap.String("userID", reqUserID), zap.Error(err))
		return nil, err
	}

	// 1. Get the list of users following by the current user
	followingRes, err := ps.UserClient.ListFollowing(ctx, &userpb.GetUserByIdRequest{
		UserId: userID,
	})
	if err != nil {
		ps.Logger.Error("failed to call UserClient.ListFollowing", zap.Error(err))
		return nil, err
	}

	// 2. If the user follows no one, return an empty feed
	if len(followingRes.GetUsers()) == 0 {
		return &models.GetFeedResponse{
			Posts:   []models.Post{},
			HasMore: false,
		}, nil
	}

	// 3. Extract IDs and update the request
	followedIDs := make([]int64, 0, len(followingRes.GetUsers()))
	for _, user := range followingRes.GetUsers() {
		followedIDs = append(followedIDs, user.GetId())
	}
	req.UserIds = followedIDs

	// 4. Call PostService with the list of followed IDs
	res, err := ps.PostClient.GetFollowingFeed(ctx, req)
	if err != nil {
		ps.Logger.Error("failed to call PostService.GetFollowingFeed", zap.Error(err))
		return nil, err
	}

	posts, err := utils.EnrichPosts(ctx, res.GetPosts(), ps.UserClient, ps.CommentClient, ps.Logger)
	if err != nil {
		return nil, err
	}

	return &models.GetFeedResponse{
		Posts:   posts,
		HasMore: res.GetHasMore(),
	}, nil
}
