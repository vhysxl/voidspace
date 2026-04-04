package handler

import (
	"context"
	pb "voidspace/comments/proto/generated/comments/v1"
	utils "voidspace/comments/utils/common"

	"github.com/vhysxl/voidspace/shared/utils/helper"
)

func (ch *CommentHandler) GetAllCommentsByPostId(ctx context.Context, req *pb.GetAllCommentsByPostIdRequest) (*pb.GetBatchCommentsResponse, error) {
	res, err := ch.CommentUsecase.GetAllCommentsByPostID(ctx, int(req.GetPostId()))
	if err != nil {
		return nil, helper.HandleError(err, ch.Logger, "GetAllCommentsByPostId")
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
