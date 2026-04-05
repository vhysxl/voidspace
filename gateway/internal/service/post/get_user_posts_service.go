package post

import (
	"context"

	"voidspaceGateway/internal/models"
	postpb "voidspaceGateway/proto/generated/posts/v1"
	userpb "voidspaceGateway/proto/generated/users/v1"
	"voidspaceGateway/utils"

	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

func (ps *PostService) GetUserPosts(
	ctx context.Context,
	targetUsername string,
	reqUserID string,
	reqUsername string,
) ([]models.Post, error) {
	ctx, cancel := context.WithTimeout(ctx, ps.ContextTimeout)
	defer cancel()

	md := utils.MetaDataHandler(reqUserID, reqUsername)
	ctx = metadata.NewOutgoingContext(ctx, md)

	user, err := ps.UserClient.GetUser(ctx, &userpb.GetUserRequest{
		Username: targetUsername,
	})
	if err != nil {
		ps.Logger.Error("failed to call UserService.GetUser", zap.Error(err))
		return nil, err
	}

	res, err := ps.PostClient.GetUserPosts(ctx, &postpb.GetUserPostsRequest{
		UserId: user.GetUser().GetId(),
	})
	if err != nil {
		ps.Logger.Error("failed to call PostService.GetUserPosts", zap.Error(err))
		return nil, err
	}

	posts, err := utils.EnrichPosts(ctx, res.GetPosts(), ps.UserClient, ps.CommentClient, ps.Logger)
	if err != nil {
		return nil, err
	}

	return posts, nil
}
