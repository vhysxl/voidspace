package comment

import (
	"context"
	"voidspaceGateway/internal/models"
	commentpb "voidspaceGateway/proto/generated/comments/v1"
	userpb "voidspaceGateway/proto/generated/users/v1"
	"voidspaceGateway/utils"

	"go.uber.org/zap"
)

func (s *CommentService) GetAllByPostID(
	ctx context.Context,
	postID int64,
) ([]models.Comment, error) {
	ctx, cancel := context.WithTimeout(ctx, s.ContextTimeout)
	defer cancel()

	commentClientRes, err := s.CommentClient.GetAllCommentsByPostId(ctx, &commentpb.GetAllCommentsByPostIdRequest{
		PostId: postID,
	})
	if err != nil {
		s.Logger.Error("failed to call CommentService.GetAllCommentsByPostID", zap.Error(err))
		return nil, err
	}

	commentsClient := commentClientRes.GetComments()

	if len(commentsClient) == 0 {
		return []models.Comment{}, nil
	}

	userIDs := make([]int64, 0, len(commentClientRes.Comments))
	for _, comment := range commentsClient {
		userIDs = append(userIDs, int64(comment.UserId))
	}

	userClientRes, err := s.UserClient.GetUsers(ctx, &userpb.GetUsersRequest{
		UserIds: userIDs,
	})
	if err != nil {
		s.Logger.Error("failed to call UserService.GetUsersByIds", zap.Error(err))
		return nil, err
	}

	users := make(map[int64]models.User)
	for _, user := range userClientRes.GetUsers() {
		users[user.GetId()] = *utils.UserMapper(user)
	}

	comments := make([]models.Comment, 0, len(commentsClient))
	for _, comment := range commentsClient {
		user := users[comment.GetUserId()]

		comments = append(comments, *utils.CommentMapper(comment, &user))
	}

	return comments, nil
}
