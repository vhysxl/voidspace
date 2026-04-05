package post

import (
	"context"

	postpb "voidspaceGateway/proto/generated/posts/v1"
	"voidspaceGateway/utils"

	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

func (ps *PostService) UnlikePost(ctx context.Context, postID int, username string, userID string) error {
	ctx, cancel := context.WithTimeout(ctx, ps.ContextTimeout)
	defer cancel()

	md := utils.MetaDataHandler(userID, username)
	ctx = metadata.NewOutgoingContext(ctx, md)

	_, err := ps.PostClient.UnlikePost(ctx, &postpb.UnlikePostRequest{
		PostId: int64(postID),
	})
	if err != nil {
		ps.Logger.Error("failed to call PostService.UnlikePost", zap.Error(err))
		return err
	}

	return nil
}
