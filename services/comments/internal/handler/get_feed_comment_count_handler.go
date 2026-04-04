package handler

import (
	"context"
	pb "voidspace/comments/proto/generated/comments/v1"

	"github.com/vhysxl/voidspace/shared/utils/helper"
)

func (ch *CommentHandler) GetFeedCommentCount(ctx context.Context, req *pb.GetFeedCommentCountRequest) (*pb.GetFeedCommentCountResponse, error) {
	postIds := make([]int, len(req.GetPostIds()))
	for i, postId := range req.GetPostIds() {
		postIds[i] = int(postId)
	}

	countMap, err := ch.CommentUsecase.GetFeedCommentCount(ctx, postIds)
	if err != nil {
		return nil, helper.HandleError(err, ch.Logger, "GetFeedCommentCount")
	}

	var pbCounts []*pb.CommentCount
	for postID, count := range countMap {
		pbCounts = append(pbCounts, &pb.CommentCount{
			PostId: int64(postID),
			Count:  int64(count),
		})
	}

	return &pb.GetFeedCommentCountResponse{
		PostCommentsCount: pbCounts,
	}, nil
}
