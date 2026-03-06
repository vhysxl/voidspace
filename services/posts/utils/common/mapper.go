package utils

import (
	"voidspace/posts/internal/domain"
	pb "voidspace/posts/proto/generated/posts"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func PostMapper(post *domain.Post) *pb.PostResponse {
	if post == nil {
		return &pb.PostResponse{}
	}

	return &pb.PostResponse{
		Id:         post.ID,
		Content:    post.Content,
		UserId:     post.UserID,
		PostImages: post.PostImages,
		LikesCount: post.LikesCount,
		CreatedAt:  timestamppb.New(post.CreatedAt),
		UpdatedAt:  timestamppb.New(post.UpdatedAt),
		IsLiked:    post.IsLiked,
	}
}
