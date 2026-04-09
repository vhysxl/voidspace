package comment

import (
	"context"
	"voidspaceGateway/internal/models"
	commentpb "voidspaceGateway/proto/generated/comments/v1"
	userpb "voidspaceGateway/proto/generated/users/v1"
	"voidspaceGateway/utils"

	"go.uber.org/zap"
)

func (s *CommentService) SearchComments(ctx context.Context, query string) ([]*models.Comment, error) {
	ctx, cancel := context.WithTimeout(ctx, s.ContextTimeout)
	defer cancel()

	commentRes, err := s.CommentClient.SearchComments(ctx, &commentpb.SearchCommentsRequest{
		Query: query,
	})
	if err != nil {
		s.Logger.Error("Failed to call CommentService.SearchComments", zap.Error(err))
		return nil, err
	}

	if len(commentRes.GetComments()) == 0 {
		return []*models.Comment{}, nil
	}

	comments := commentRes.GetComments()
	userIDsMap := make(map[int64]bool)
	for _, c := range comments {
		userIDsMap[c.GetUserId()] = true
	}

	userIDs := make([]int64, 0, len(userIDsMap))
	for id := range userIDsMap {
		userIDs = append(userIDs, id)
	}

	usersRes, err := s.UserClient.GetUsers(ctx, &userpb.GetUsersRequest{
		UserIds: userIDs,
	})
	if err != nil {
		s.Logger.Error("Failed to fetch users for comment hydration", zap.Error(err))
		// Continue anyway, mapper will handle nil user
	}

	userMap := make(map[int64]*userpb.UserProfile)
	if usersRes != nil {
		for _, u := range usersRes.GetUsers() {
			userMap[u.GetId()] = u
		}
	}

	hydratedComments := make([]*models.Comment, 0, len(comments))
	for _, c := range comments {
		author := utils.UserMapper(userMap[c.GetUserId()])
		hydratedComments = append(hydratedComments, utils.CommentMapper(c, author))
	}

	return hydratedComments, nil
}
