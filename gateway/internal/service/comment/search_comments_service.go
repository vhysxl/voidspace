package comment

import (
	"context"
	commentpb "voidspaceGateway/proto/generated/comments/v1"
)

func (s *CommentService) SearchComments(ctx context.Context, query string) (*commentpb.SearchCommentsResponse, error) {
	return s.CommentClient.SearchComments(ctx, &commentpb.SearchCommentsRequest{
		Query: query,
	})
}
