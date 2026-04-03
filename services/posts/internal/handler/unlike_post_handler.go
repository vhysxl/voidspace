package handler

import (
	"context"
	"voidspace/posts/internal/domain"
	pb "voidspace/posts/proto/generated/posts/v1"

	"github.com/vhysxl/voidspace/shared/utils/helper"
	"github.com/vhysxl/voidspace/shared/utils/interceptor"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *PostHandler) UnlikePost(
	ctx context.Context,
	req *pb.UnlikePostRequest,
) (*emptypb.Empty, error) {
	userID, err := helper.GetUserIDFromContext(ctx, interceptor.CtxKeyUserID)
	if err != nil {
		return nil, helper.HandleError(err, h.Logger, "Unlike Post")
	}

	like := &domain.Like{
		PostID: int(req.GetPostId()),
		UserID: userID,
	}

	err = h.LikeUsecase.UnlikePost(ctx, like)
	if err != nil {
		return nil, helper.HandleError(err, h.Logger, "Unlike Post")
	}

	return &emptypb.Empty{}, nil
}
