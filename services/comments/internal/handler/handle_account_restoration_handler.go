package handler

import (
	"context"
	pb "voidspace/comments/proto/generated/comments/v1"

	"github.com/vhysxl/voidspace/shared/utils/helper"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (ch *CommentHandler) HandleAccountRestoration(ctx context.Context, req *pb.HandleAccountRestorationRequest) (*emptypb.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, ch.ContextTimeout)
	defer cancel()

	err := ch.CommentUsecase.HandleAccountRestoration(ctx, int(req.GetUserId()))
	if err != nil {
		return nil, helper.HandleError(err, ch.Logger, "HandleAccountRestoration")
	}

	return &emptypb.Empty{}, nil
}
