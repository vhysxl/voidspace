package post

import (
	"context"

	commentpb "voidspaceGateway/proto/generated/comments/v1"
	postpb "voidspaceGateway/proto/generated/posts/v1"
	"voidspaceGateway/utils"

	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

func (ps *PostService) Delete(ctx context.Context, postID int64, username string, userID string) error {
	ctx, cancel := context.WithTimeout(ctx, ps.ContextTimeout)
	defer cancel()

	md := utils.MetaDataHandler(userID, username)
	ctx = metadata.NewOutgoingContext(ctx, md)

	// 1. Delete Post via gRPC
	_, err := ps.PostClient.DeletePost(ctx, &postpb.DeletePostRequest{
		PostId: postID,
	})
	if err != nil {
		ps.Logger.Error("failed to call PostService.DeletePost", zap.Error(err))
		return err
	}

	// 2. Delete Comments associated with post via gRPC
	_, err = ps.CommentClient.HandlePostDeletion(ctx, &commentpb.HandlePostDeletionRequest{
		PostId: postID,
	})
	if err != nil {
		ps.Logger.Error("failed to delete post's comments", zap.Error(err))
	}

	return nil
}
