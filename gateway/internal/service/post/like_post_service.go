package post

import (
	"context"

	postpb "voidspaceGateway/proto/generated/posts/v1"
	"voidspaceGateway/utils"

	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

func (ps *PostService) LikePost(ctx context.Context, postID int, username string, userID string) error {
	ctx, cancel := context.WithTimeout(ctx, ps.ContextTimeout)
	defer cancel()

	md := utils.MetaDataHandler(userID, username)
	ctx = metadata.NewOutgoingContext(ctx, md)

	_, err := ps.PostClient.LikePost(ctx, &postpb.LikePostRequest{
		PostId: int64(postID),
	})
	if err != nil {
		ps.Logger.Error("failed to call PostService.LikePost", zap.Error(err))
		return err
	}

	return nil
}
