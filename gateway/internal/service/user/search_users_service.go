package user

import (
	"context"
	userpb "voidspaceGateway/proto/generated/users/v1"
)

func (s *UserService) SearchUsers(ctx context.Context, query string) (*userpb.SearchUsersResponse, error) {
	return s.UserClient.SearchUsers(ctx, &userpb.SearchUsersRequest{
		Query: query,
	})
}
