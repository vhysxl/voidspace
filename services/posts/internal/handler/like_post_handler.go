package handler

import (
	"context"
	"voidspace/posts/internal/domain"
	pb "voidspace/posts/proto/generated/posts/v1"

	"github.com/vhysxl/voidspace/shared/utils/helper"
	"github.com/vhysxl/voidspace/shared/utils/interceptor"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *PostHandler) LikePost(ctx context.Context, req *pb.LikePostRequest) (*emptypb.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, h.ContextTimeout)
	defer cancel()

	userID, err := helper.GetUserIDFromContext(ctx, interceptor.CtxKeyUserID)
	if err != nil {
		return nil, err
	}

	like := &domain.Like{
		PostID: int(req.GetPostId()),
		UserID: userID,
	}

	err = h.LikeUsecase.LikePost(ctx, like)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
