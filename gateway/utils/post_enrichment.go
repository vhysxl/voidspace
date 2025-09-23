package utils

import (
	"context"
	"voidspaceGateway/internal/models"
	commentpb "voidspaceGateway/proto/generated/comments"
	postpb "voidspaceGateway/proto/generated/posts"
	userpb "voidspaceGateway/proto/generated/users"

	"go.uber.org/zap"
)

func EnrichPosts(
	ctx context.Context,
	posts []*postpb.PostResponse,
	commentClient commentpb.CommentServiceClient,
	userClient userpb.UserServiceClient,
	logger *zap.Logger,
) ([]models.Post, error) {

	// Guard if post is empty
	if len(posts) == 0 {
		return []models.Post{}, nil
	}

	postIds := make([]int32, 0, len(posts))
	for _, post := range posts {
		postIds = append(postIds, post.GetId())
	}

	commentsCountRes, err := commentClient.GetCommentsCountByPostIds(ctx, &commentpb.GetCommentsCountByPostIdsRequest{
		PostIds: postIds,
	})
	if err != nil {
		logger.Error("failed to call CommentService.GetCommentsCountByPostIds", zap.Error(err))
		return nil, err
	}

	//    Collect unique user IDs from the posts so we know which users we need
	//    to enrich the feed with. Using a map as a set avoids duplicates.
	userIDSet := make(map[int32]struct{})
	for _, post := range posts {
		userIDSet[post.GetUserId()] = struct{}{}
	}

	//   Convert the set into a slice, since gRPC request requires a slice.
	userIDs := make([]int32, 0, len(userIDSet))
	for id := range userIDSet {
		userIDs = append(userIDs, id)
	}

	//    Call UserService in batch to fetch all user details for the collected IDs.
	//    This is the critical optimization: instead of calling UserService N times
	//    (once per post), we only call it ONCE with all needed IDs.
	userRes, err := userClient.GetUsersByIds(ctx, &userpb.GetUserByUserIDsRequest{
		UserID: userIDs,
	})
	if err != nil {
		logger.Error("failed to call UserService.GetUsersByIds", zap.Error(err))
		return nil, err
	}

	//    Build a map[user_id]User for quick lookup when merging user info into posts.
	userMap := make(map[int32]*userpb.User)
	for _, u := range userRes.GetUsers() {
		userMap[u.GetId()] = u
	}

	//    Merge posts with user data.
	//    For each post, we lookup its author in userMap.
	//    If found, we map the gRPC User protobuf into our domain model (models.User).
	enriched := make([]models.Post, 0, len(posts))
	for _, p := range posts {
		u := userMap[p.GetUserId()]

		var user *models.User
		if u != nil {
			user = UserMapperFromUser(u)
		}

		commentsCount := int(commentsCountRes.GetPostCommentsCount()[p.GetId()])
		enriched = append(enriched, *PostMapper(p, user, commentsCount))
	}

	return enriched, nil
}
