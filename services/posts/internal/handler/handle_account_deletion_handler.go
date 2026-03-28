package handler

import (
	"context"
	pb "voidspace/posts/proto/generated/posts/v1"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *PostHandler) HandleAccountDeletion(ctx context.Context, req *pb.HandleAccountDeletionRequest) (*emptypb.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, h.ContextTimeout)
	defer cancel()

	err := h.PostUsecase.HandleAccountDeletion(ctx, int(req.GetUserId()))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
