package handler

import (
	"context"
	"voidspace/posts/internal/domain"
	pb "voidspace/posts/proto/generated/posts/v1"
	"voidspace/posts/utils"

	"github.com/vhysxl/voidspace/shared/utils/helper"
	"github.com/vhysxl/voidspace/shared/utils/interceptor"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *PostHandler) UpdatePost(ctx context.Context, req *pb.UpdatePostRequest) (*emptypb.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, h.ContextTimeout)
	defer cancel()

	userID, err := helper.GetUserIDFromContext(ctx, interceptor.CtxKeyUserID)
	if err != nil {
		return nil, helper.HandleError(err, h.Logger, "Update Post")
	}

	post := &domain.Post{
		ID:         int(req.GetPostId()),
		Content:    req.GetContent(),
		PostImages: utils.MapPbPostImageToDomain(req.GetImages()),
	}

	err = h.PostUsecase.UpdatePost(ctx, post, userID)
	if err != nil {
		return nil, helper.HandleError(err, h.Logger, "Update Post")
	}

	return &emptypb.Empty{}, nil
}
