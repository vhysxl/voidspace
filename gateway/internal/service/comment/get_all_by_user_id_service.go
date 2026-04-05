package comment

import (
	"context"
	"voidspaceGateway/internal/models"
	commentpb "voidspaceGateway/proto/generated/comments/v1"
	usersv1 "voidspaceGateway/proto/generated/users/v1"
	"voidspaceGateway/utils"

	"go.uber.org/zap"
)

func (s *CommentService) GetAllByUser(
	ctx context.Context,
	username string,
) ([]models.Comment, error) {
	ctx, cancel := context.WithTimeout(ctx, s.ContextTimeout)
	defer cancel()

	userClientRes, err := s.UserClient.GetUser(ctx, &usersv1.GetUserRequest{
		Username: username,
	})
	if err != nil {
		s.Logger.Error("failed to call UserService.GetUser", zap.Error(err))
		return nil, err
	}

	commentClientRes, err := s.CommentClient.GetAllCommentsByUserId(ctx, &commentpb.GetAllCommentsByUserIdRequest{
		UserId: userClientRes.User.GetId(),
	})
	if err != nil {
		s.Logger.Error("failed to call CommentService.GetAllCommentsByUserID", zap.Error(err))
		return nil, err
	}

	commentsClient := commentClientRes.GetComments()

	if len(commentsClient) == 0 {
		return []models.Comment{}, nil
	}

	user := utils.UserMapper(userClientRes.GetUser())

	comments := make([]models.Comment, 0, len(commentsClient))
	for _, comment := range commentsClient {
		comments = append(comments, *utils.CommentMapper(comment, user))
	}

	return comments, nil

}
