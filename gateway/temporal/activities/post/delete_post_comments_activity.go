package post

import (
	"context"

	commentpb "voidspaceGateway/proto/generated/comments/v1"
	temporal_dto "voidspaceGateway/temporal/dto"

	"go.uber.org/zap"
)

const DeletePostCommentsActivity = "DeletePostCommentsActivity"

func (pa *PostActivities) DeletePostCommentsActivity(
	ctx context.Context,
	req temporal_dto.DeletePostReq,
) error {
	pa.Logger.Info(
		"Starting Delete Post's Comments Activity",
		zap.Int64("postID", req.PostID),
	)

	_, err := pa.CommentClient.HandlePostDeletion(ctx, &commentpb.HandlePostDeletionRequest{
		PostId: req.PostID,
	})
	if err != nil {
		pa.Logger.Error("failed to delete post's comments", zap.Error(err))
		return err
	}

	pa.Logger.Info("Post's comments deleted successfully")
	return nil
}
