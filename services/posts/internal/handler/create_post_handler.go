package handler

import (
	"context"
	"voidspace/posts/internal/domain"
	pb "voidspace/posts/proto/generated/posts/v1"
	"voidspace/posts/utils"

	"github.com/vhysxl/voidspace/shared/utils/helper"
	"github.com/vhysxl/voidspace/shared/utils/interceptor"
)

func (h *PostHandler) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.Post, error) {
	ctx, cancel := context.WithTimeout(ctx, h.ContextTimeout)
	defer cancel()

	userID, err := helper.GetUserIDFromContext(ctx, interceptor.CtxKeyUserID)
	if err != nil {
		return nil, err
	}

	post := &domain.Post{
		Content:    req.GetContent(),
		UserID:     userID,
		PostImages: utils.MapPbPostImageToDomain(req.GetImages()),
	}

	createdPost, err := h.PostUsecase.CreatePost(ctx, post)
	if err != nil {
		return nil, err
	}

	return utils.MapDomainPostToPb(createdPost), nil
}
