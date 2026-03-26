package handler

import (
	"context"
	pb "voidspace/posts/proto/generated/posts/v1"
	"voidspace/posts/utils"

	"github.com/vhysxl/voidspace/shared/utils/helper"
	"github.com/vhysxl/voidspace/shared/utils/interceptor"
)

func (h *PostHandler) GetPost(ctx context.Context, req *pb.GetPostRequest) (*pb.Post, error) {
	ctx, cancel := context.WithTimeout(ctx, h.ContextTimeout)
	defer cancel()

	userID, err := helper.GetOptionalUserIDFromContext(ctx, interceptor.CtxKeyUserID)
	if err != nil {
		return nil, helper.HandleError(err, h.Logger, "Get Post")
	}

	var loggedInUserID *int
	if userID != 0 {
		loggedInUserID = &userID
	}

	post, err := h.PostUsecase.GetPost(ctx, int(req.GetPostId()), loggedInUserID)
	if err != nil {
		return nil, helper.HandleError(err, h.Logger, "Get Post")
	}

	return utils.MapDomainPostToPb(post), nil
}
