package handler

import (
	"context"
	pb "voidspace/comments/proto/generated/comments/v1"

	"github.com/vhysxl/voidspace/shared/utils/helper"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *CommentHandler) HandlePostDeletion(
	ctx context.Context,
	req *pb.HandlePostDeletionRequest,
) (*emptypb.Empty, error) {
	postID := int(req.GetPostId())

	err := h.CommentUsecase.DeleteByPostID(ctx, postID)
	if err != nil {
		return nil, helper.HandleError(err, h.Logger, "HandlePostDeletion")
	}

	return &emptypb.Empty{}, nil
}
