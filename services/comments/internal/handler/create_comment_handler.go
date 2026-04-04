package handler

import (
	"context"
	"voidspace/comments/internal/domain"
	pb "voidspace/comments/proto/generated/comments/v1"
	utils "voidspace/comments/utils/common"

	"github.com/vhysxl/voidspace/shared/utils/helper"
	"github.com/vhysxl/voidspace/shared/utils/interceptor"
)

func (ch *CommentHandler) CreateComment(ctx context.Context, req *pb.CreateCommentRequest) (*pb.Comment, error) {
	userId, err := helper.GetUserIDFromContext(ctx, interceptor.CtxKeyUserID)
	if err != nil {
		return nil, helper.HandleAuthError(nil, ch.Logger)
	}

	data := &domain.Comment{
		UserID:  userId,
		PostID:  int(req.GetPostId()),
		Content: req.GetContent(),
	}

	comment, err := ch.CommentUsecase.CreateComment(ctx, data)
	if err != nil {
		return nil, helper.HandleError(err, ch.Logger, "CreateComment")
	}

	return utils.CommentMapper(comment), nil
}
