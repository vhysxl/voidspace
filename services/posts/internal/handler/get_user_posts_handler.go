package handler

import (
	"context"
	pb "voidspace/posts/proto/generated/posts/v1"
	"voidspace/posts/utils"

	"github.com/vhysxl/voidspace/shared/utils/helper"
	"github.com/vhysxl/voidspace/shared/utils/interceptor"
)

func (h *PostHandler) GetUserPosts(ctx context.Context, req *pb.GetUserPostsRequest) (*pb.GetPostsResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, h.ContextTimeout)
	defer cancel()

	userID, err := helper.GetOptionalUserIDFromContext(ctx, interceptor.CtxKeyUserID)
	if err != nil {
		return nil, err
	}

	var loggedInUserID *int
	if userID != 0 {
		loggedInUserID = &userID
	}

	posts, err := h.PostUsecase.GetUserPosts(ctx, int(req.GetUserId()), loggedInUserID)
	if err != nil {
		return nil, err
	}

	pbPosts := make([]*pb.Post, len(posts))
	for i, p := range posts {
		pbPosts[i] = utils.MapDomainPostToPb(&p)
	}

	return &pb.GetPostsResponse{Posts: pbPosts}, nil
}
