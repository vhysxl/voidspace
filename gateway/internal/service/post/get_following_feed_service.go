package post

import (
	"context"

	"voidspaceGateway/internal/models"
	postpb "voidspaceGateway/proto/generated/posts/v1"
	"voidspaceGateway/utils"

	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

func (ps *PostService) GetFollowingFeed(ctx context.Context, req *postpb.GetFollowingFeedRequest, reqUserID string, reqUsername string) (*models.GetFeedResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, ps.ContextTimeout)
	defer cancel()

	md := utils.MetaDataHandler(reqUserID, reqUsername)
	ctx = metadata.NewOutgoingContext(ctx, md)

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
