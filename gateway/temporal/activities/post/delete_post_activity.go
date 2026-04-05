package post

import (
	"context"

	postpb "voidspaceGateway/proto/generated/posts/v1"
	temporal_dto "voidspaceGateway/temporal/dto"
	"voidspaceGateway/utils"

	"go.temporal.io/sdk/temporal"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const DeletePostActivity = "DeletePostActivity"

func (pa *PostActivities) DeletePostActivity(
	ctx context.Context,
	req temporal_dto.DeletePostReq,
) error {
	pa.Logger.Info(
		"Starting Delete Post Activity",
		zap.Int64("postID", req.PostID),
		zap.String("userID", req.UserID),
		zap.String("username", req.Username),
	)

	md := utils.MetaDataHandler(req.UserID, req.Username)
	ctx = metadata.NewOutgoingContext(ctx, md)

	_, err := pa.PostClient.DeletePost(ctx, &postpb.DeletePostRequest{
		PostId: req.PostID,
	})
	if err != nil {
		pa.Logger.Error("failed to call PostService.DeletePost", zap.Error(err))
		st, ok := status.FromError(err)
		if ok {
			return temporal.NewApplicationErrorWithCause(st.Message(), st.Code().String(), err)
		}
	}

	return nil
}
