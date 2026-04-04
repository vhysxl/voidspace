package handler

import (
	"context"
	pb "voidspace/comments/proto/generated/comments/v1"
	utils "voidspace/comments/utils/common"

	"github.com/vhysxl/voidspace/shared/utils/helper"
)

func (ch *CommentHandler) GetAllCommentsByUserId(ctx context.Context, req *pb.GetAllCommentsByUserIdRequest) (*pb.GetBatchCommentsResponse, error) {
	res, err := ch.CommentUsecase.GetAllCommentsByUserID(ctx, int(req.GetUserId()))
	if err != nil {
		return nil, helper.HandleError(err, ch.Logger, "GetAllCommentsByUserId")
	}

	pbComments := make([]*pb.Comment, 0, len(res.Comments))
	for _, c := range res.Comments {
		pbComments = append(pbComments, utils.CommentMapper(c))
	}

	return &pb.GetBatchCommentsResponse{
		Comments:      pbComments,
		CommentsCount: int64(res.CommentsCount),
	}, nil
}
