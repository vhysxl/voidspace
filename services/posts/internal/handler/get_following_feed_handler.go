package handler

import (
	"context"
	"time"
	pb "voidspace/posts/proto/generated/posts/v1"
	"voidspace/posts/utils"

	"github.com/vhysxl/voidspace/shared/utils/helper"
	"github.com/vhysxl/voidspace/shared/utils/interceptor"
)

func (h *PostHandler) GetFollowingFeed(ctx context.Context, req *pb.GetFollowingFeedRequest) (*pb.GetFeedResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, h.ContextTimeout)
	defer cancel()

	userID, err := helper.GetUserIDFromContext(ctx, interceptor.CtxKeyUserID)
	if err != nil {
		return nil, err
	}

	var cursorTime *time.Time
	if req.GetCursorTime() != nil {
		t := req.GetCursorTime().AsTime()
		cursorTime = &t
	}

	userIDs := make([]int, len(req.GetUserIds()))
	for i, id := range req.GetUserIds() {
		userIDs[i] = int(id)
	}

	posts, hasMore, err := h.PostUsecase.GetFollowingFeed(ctx, cursorTime, int(req.GetCursorId()), userID, userIDs)
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
