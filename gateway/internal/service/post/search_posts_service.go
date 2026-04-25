package post

import (
	"context"
	"voidspaceGateway/internal/models"
	commentpb "voidspaceGateway/proto/generated/comments/v1"
	postpb "voidspaceGateway/proto/generated/posts/v1"
	userpb "voidspaceGateway/proto/generated/users/v1"
	"voidspaceGateway/utils"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

func (s *PostService) SearchPosts(ctx context.Context, query string) ([]*models.Post, error) {
	ctx, cancel := context.WithTimeout(ctx, s.ContextTimeout)
	defer cancel()

	postRes, err := s.PostClient.SearchPosts(ctx, &postpb.SearchPostsRequest{
		Query: query,
	})
	if err != nil {
		s.Logger.Error("Failed to call PostService.SearchPosts", zap.Error(err))
		return nil, err
	}

	if len(postRes.GetPosts()) == 0 {
		return []*models.Post{}, nil
	}

	posts := postRes.GetPosts()
	userIDsMap := make(map[int64]bool)
	postIDs := make([]int64, 0, len(posts))

	for _, p := range posts {
		userIDsMap[p.GetUserId()] = true
		postIDs = append(postIDs, p.GetId())
	}

	userIDs := make([]int64, 0, len(userIDsMap))
	for id := range userIDsMap {
		userIDs = append(userIDs, id)
	}

	g, gCtx := errgroup.WithContext(ctx)

	var usersRes *userpb.GetUsersResponse
	g.Go(func() error {
		var err error
		usersRes, err = s.UserClient.GetUsers(gCtx, &userpb.GetUsersRequest{
			UserIds: userIDs,
		})
		if err != nil {
			s.Logger.Error("Failed to call UserService.GetUsers", zap.Error(err))
		}
		return err
	})

	var countsRes *commentpb.GetFeedCommentCountResponse
	g.Go(func() error {
		var err error
		countsRes, err = s.CommentClient.GetFeedCommentCount(gCtx, &commentpb.GetFeedCommentCountRequest{
			PostIds: postIDs,
		})
		if err != nil {
			s.Logger.Error("Failed to call CommentService.GetFeedCommentCount", zap.Error(err))
		}
		return err
	})

	if err := g.Wait(); err != nil {
		s.Logger.Warn("Hydration partially failed, some data might be missing", zap.Error(err))
	}

	// Create lookups
	userMap := make(map[int64]*userpb.UserProfile)
	if usersRes != nil {
		for _, u := range usersRes.GetUsers() {
			userMap[u.GetId()] = u
		}
	}

	countMap := make(map[int64]*commentpb.CommentCount)
	if countsRes != nil {
		for _, c := range countsRes.GetPostCommentsCount() {
			countMap[c.GetPostId()] = c
		}
	}

	// Map to hydrated models
	hydratedPosts := make([]*models.Post, 0, len(posts))
	for _, p := range posts {
		author := utils.UserMapper(userMap[p.GetUserId()])
		commentCount := 0
		if c, ok := countMap[p.GetId()]; ok {
			commentCount = int(c.GetCount())
		}
		hydratedPosts = append(hydratedPosts, utils.PostMapper(p, author, commentCount))
	}

	return hydratedPosts, nil
}
