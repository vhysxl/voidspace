package handler

import (
	"context"
	pb "voidspace/comments/proto/generated/comments/v1"

	"github.com/vhysxl/voidspace/shared/utils/helper"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (ch *CommentHandler) HandleAccountDeletion(ctx context.Context, req *pb.HandleAccountDeletionRequest) (*emptypb.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, ch.ContextTimeout)
	defer cancel()

	err := ch.CommentUsecase.HandleAccountDeletion(ctx, int(req.GetUserId()))
	if err != nil {
		return nil, helper.HandleError(err, ch.Logger, "HandleAccountDeletion")
	}

	return &emptypb.Empty{}, nil
}
