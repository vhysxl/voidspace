package handler

import (
	"context"
	pb "voidspace/posts/proto/generated/posts/v1"

	"github.com/vhysxl/voidspace/shared/utils/helper"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *PostHandler) HandleAccountDeletion(
	ctx context.Context,
	req *pb.HandleAccountDeletionRequest,
) (*emptypb.Empty, error) {
	err := h.PostUsecase.HandleAccountDeletion(ctx, int(req.GetUserId()))
	if err != nil {
		return nil, helper.HandleError(err, h.Logger, "Handle Account Deletion")
	}

	return &emptypb.Empty{}, nil
}
