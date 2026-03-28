package handler

import (
	"context"
	pb "voidspace/posts/proto/generated/posts/v1"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *PostHandler) HandleAccountRestoration(ctx context.Context, req *pb.HandleAccountRestorationRequest) (*emptypb.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, h.ContextTimeout)
	defer cancel()

	err := h.PostUsecase.HandleAccountRestoration(ctx, int(req.GetUserId()))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
