package handler

import (
	"context"
	pb "voidspace/comments/proto/generated/comments/v1"
	utils "voidspace/comments/utils/common"

	"github.com/vhysxl/voidspace/shared/utils/helper"
)

func (ch *CommentHandler) SearchComments(
	ctx context.Context,
	req *pb.SearchCommentsRequest,
) (*pb.SearchCommentsResponse, error) {
	comments, err := ch.CommentUsecase.SearchComments(ctx, req.GetQuery())
	if err != nil {
		return nil, helper.HandleError(err, ch.Logger, "SearchComments")
	}

	pbComments := make([]*pb.Comment, 0, len(comments))
	for _, c := range comments {
		pbComments = append(pbComments, utils.CommentMapper(c))
	}

	return &pb.SearchCommentsResponse{
		Comments: pbComments,
	}, nil
}
