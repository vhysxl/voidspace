package handler

import (
	"context"
	pb "voidspace/posts/proto/generated/posts/v1"

	"github.com/vhysxl/voidspace/shared/utils/helper"
	"github.com/vhysxl/voidspace/shared/utils/interceptor"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *PostHandler) DeletePost(ctx context.Context, req *pb.DeletePostRequest) (*emptypb.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, h.ContextTimeout)
	defer cancel()

	userID, err := helper.GetUserIDFromContext(ctx, interceptor.CtxKeyUserID)
	if err != nil {
		return nil, helper.HandleError(err, h.Logger, "Delete Post")
	}

	err = h.PostUsecase.DeletePost(ctx, int(req.GetPostId()), userID)
	if err != nil {
		return nil, helper.HandleError(err, h.Logger, "Delete Post")
	}

	return &emptypb.Empty{}, nil
}
