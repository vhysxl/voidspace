package comment

import (
	"context"
	"voidspaceGateway/internal/models"
	commentpb "voidspaceGateway/proto/generated/comments/v1"
	"voidspaceGateway/utils"

	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

func (s *CommentService) Create(
	ctx context.Context,
	req *models.CreateCommentRequest,
	userID, username string) error {
	ctx, cancel := context.WithTimeout(ctx, s.ContextTimeout)
	defer cancel()

	md := utils.MetaDataHandler(userID, username)
	ctx = metadata.NewOutgoingContext(ctx, md)

	_, err := s.CommentClient.CreateComment(ctx, &commentpb.CreateCommentRequest{
		PostId:  int64(req.PostID),
		Content: req.Content,
	})
	if err != nil {
		s.Logger.Error("failed to call CommentService.CreteComment", zap.Error(err))
		return err
	}

	return nil
}
