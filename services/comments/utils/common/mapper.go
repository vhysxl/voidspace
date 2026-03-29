package utils

import (
	"voidspace/comments/internal/domain"
	pb "voidspace/comments/proto/generated/comments/v1"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func CommentMapper(comment *domain.Comment) *pb.Comment {
	if comment == nil {
		return &pb.Comment{}
	}

	return &pb.Comment{
		Id:        int64(comment.ID),
		UserId:    int64(comment.UserID),
		PostId:    int64(comment.PostID),
		Content:   comment.Content,
		CreatedAt: timestamppb.New(comment.CreatedAt),
	}
}
