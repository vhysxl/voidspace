package utils

import (
	"voidspace/comments/internal/domain"
	pb "voidspace/comments/proto/generated/comments"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func CommentMapper(comment *domain.Comment) *pb.CommentResponse {
	if comment == nil {
		return &pb.CommentResponse{}
	}

	return &pb.CommentResponse{
		Id:        int32(comment.UserID),
		UserId:    int32(comment.UserID),
		PostId:    int32(comment.UserID),
		Content:   comment.Content,
		CreatedAt: timestamppb.New(comment.CreatedAt),
	}
}
