package handler

import (
	"context"
	"time"
	pb "voidspace/posts/proto/generated/posts/v1"
	"voidspace/posts/utils"

	"github.com/vhysxl/voidspace/shared/utils/helper"
	"github.com/vhysxl/voidspace/shared/utils/interceptor"
)

func (h *PostHandler) GetGlobalFeed(ctx context.Context, req *pb.GetGlobalFeedRequest) (*pb.GetFeedResponse, error) {
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

	var cursorTime *time.Time
	if req.GetCursorTime() != nil {
		t := req.GetCursorTime().AsTime()
		cursorTime = &t
	}

	posts, hasMore, err := h.PostUsecase.GetGlobalFeed(ctx, cursorTime, int(req.GetCursorId()), loggedInUserID)
	if err != nil {
		return nil, err
	}

	pbPosts := make([]*pb.Post, len(posts))
	for i, p := range posts {
		pbPosts[i] = utils.MapDomainPostToPb(&p)
	}

	return &pb.GetFeedResponse{
		Posts:   pbPosts,
		HasMore: hasMore,
	}, nil
}
