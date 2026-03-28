package handler

import (
	"time"
	"voidspace/posts/internal/domain"
	pb "voidspace/posts/proto/generated/posts/v1"
)

type PostHandler struct {
	pb.UnimplementedPostServiceServer

	PostUsecase domain.PostUsecase
	LikeUsecase domain.LikeUsecase
	ContextTimeout time.Duration
}

func NewPostHandler(
	postUsecase domain.PostUsecase,
	likeUsecase domain.LikeUsecase,
	timeout time.Duration,
) pb.PostServiceServer {
	return &PostHandler{
		PostUsecase:    postUsecase,
		LikeUsecase:    likeUsecase,
		ContextTimeout: timeout,
	}
}
