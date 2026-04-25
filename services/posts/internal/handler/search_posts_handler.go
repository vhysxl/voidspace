package handler

import (
	"context"
	pb "voidspace/posts/proto/generated/posts/v1"
	"voidspace/posts/utils"

	"github.com/vhysxl/voidspace/shared/utils/helper"
)

func (h *PostHandler) SearchPosts(
	ctx context.Context,
	req *pb.SearchPostsRequest,
) (*pb.SearchPostsResponse, error) {
	posts, err := h.PostUsecase.SearchPosts(ctx, req.GetQuery())
	if err != nil {
		return nil, helper.HandleError(err, h.Logger, "Search Posts")
	}

	pbPosts := make([]*pb.Post, len(posts))
	for i, p := range posts {
		pbPosts[i] = utils.MapDomainPostToPb(&p)
	}

	return &pb.SearchPostsResponse{
		Posts: pbPosts,
	}, nil
}
