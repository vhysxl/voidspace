package handler

import (
	"context"
	pb "voidspace/comments/proto/generated/comments/v1"

	"github.com/vhysxl/voidspace/shared/utils/helper"
	"github.com/vhysxl/voidspace/shared/utils/interceptor"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (ch *CommentHandler) DeleteComment(ctx context.Context, req *pb.DeleteCommentRequest) (*emptypb.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, ch.ContextTimeout)
	defer cancel()

	userId, err := helper.GetUserIDFromContext(ctx, interceptor.CtxKeyUserID)
	if err != nil {
		return nil, helper.HandleAuthError(nil, ch.Logger)
	}

	err = ch.CommentUsecase.DeleteComment(ctx, int(req.GetCommentId()), userId)
	if err != nil {
		return nil, helper.HandleError(err, ch.Logger, "DeleteComment")
	}

	return &emptypb.Empty{}, nil
}
