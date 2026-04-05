package utils

import (
	"context"
	"voidspaceGateway/internal/models"
	commentpb "voidspaceGateway/proto/generated/comments/v1"
	postpb "voidspaceGateway/proto/generated/posts/v1"
	userpb "voidspaceGateway/proto/generated/users/v1"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

// EnrichPosts takes a slice of proto posts and enriches them with user data
// and comment counts by batch-fetching both in parallel.
func EnrichPosts(
	ctx context.Context,
	posts []*postpb.Post,
	userClient userpb.UserServiceClient,
	commentClient commentpb.CommentServiceClient,
	logger *zap.Logger,
) ([]models.Post, error) {
	if len(posts) == 0 {
		return []models.Post{}, nil
	}

	// Collect unique user IDs and all post IDs in a single pass
	userIDSet := make(map[int64]struct{})
	postIDs := make([]int64, 0, len(posts))
	for _, p := range posts {
		userIDSet[p.GetUserId()] = struct{}{}
		postIDs = append(postIDs, p.GetId())
	}

	userIDs := make([]int64, 0, len(userIDSet))
	for id := range userIDSet {
		userIDs = append(userIDs, id)
	}

	// Batch-fetch users and comment counts in parallel
	g, gCtx := errgroup.WithContext(ctx)

	var usersRes *userpb.GetUsersResponse
	g.Go(func() error {
		var err error
		usersRes, err = userClient.GetUsers(gCtx, &userpb.GetUsersRequest{
			UserIds: userIDs,
		})
		if err != nil {
			logger.Error("failed to call UserService.GetUsers", zap.Error(err))
		}
		return err
	})

	var commentCountRes *commentpb.GetFeedCommentCountResponse
	g.Go(func() error {
		var err error
		commentCountRes, err = commentClient.GetFeedCommentCount(gCtx, &commentpb.GetFeedCommentCountRequest{
			PostIds: postIDs,
		})
		if err != nil {
			logger.Error("failed to call CommentService.GetFeedCommentCount", zap.Error(err))
		}
		return err
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}

	// Build user lookup map: user_id -> *models.User
	userMap := make(map[int64]*models.User, len(usersRes.GetUsers()))
	for _, u := range usersRes.GetUsers() {
		userMap[u.GetId()] = UserMapper(u)
	}

	// Build comment count lookup map: post_id -> count
	commentCountMap := make(map[int64]int64, len(commentCountRes.GetPostCommentsCount()))
	for _, cc := range commentCountRes.GetPostCommentsCount() {
		commentCountMap[cc.GetPostId()] = cc.GetCount()
	}

	// Merge everything into enriched posts
	enriched := make([]models.Post, 0, len(posts))
	for _, p := range posts {
		author := userMap[p.GetUserId()]
		commentCount := int(commentCountMap[p.GetId()])
		enriched = append(enriched, *PostMapper(p, author, commentCount))
	}

	return enriched, nil
}
