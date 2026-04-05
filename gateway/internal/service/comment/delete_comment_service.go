package comment

import (
	"context"
	commentsv1 "voidspaceGateway/proto/generated/comments/v1"
	"voidspaceGateway/utils"

	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

func (s *CommentService) Delete(
	ctx context.Context,
	username, userID string,
	commentID int64,
) error {
	ctx, cancel := context.WithTimeout(ctx, s.ContextTimeout)
	defer cancel()

	md := utils.MetaDataHandler(userID, username)
	ctx = metadata.NewOutgoingContext(ctx, md)

	_, err := s.CommentClient.DeleteComment(ctx, &commentsv1.DeleteCommentRequest{
		CommentId: commentID,
	})
	if err != nil {
		s.Logger.Error("failed to call PostService.CreatePost", zap.Error(err))
		return err
	}

	return err
}
