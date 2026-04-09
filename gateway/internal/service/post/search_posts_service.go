package post

import (
	"context"
	postpb "voidspaceGateway/proto/generated/posts/v1"
)

func (s *PostService) SearchPosts(ctx context.Context, query string) (*postpb.SearchPostsResponse, error) {
	return s.PostClient.SearchPosts(ctx, &postpb.SearchPostsRequest{
		Query: query,
	})
}
