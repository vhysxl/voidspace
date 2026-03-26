package handler

import (
	"time"
	"voidspace/posts/internal/domain"
	pb "voidspace/posts/proto/generated/posts/v1"

	"go.uber.org/zap"
)

type PostHandler struct {
	pb.UnimplementedPostServiceServer

	PostUsecase    domain.PostUsecase
	LikeUsecase    domain.LikeUsecase
	Logger         *zap.Logger
	ContextTimeout time.Duration
}

func NewPostHandler(
	postUsecase domain.PostUsecase,
	likeUsecase domain.LikeUsecase,
	logger *zap.Logger,
	timeout time.Duration,
) pb.PostServiceServer {
	return &PostHandler{
		PostUsecase:    postUsecase,
		LikeUsecase:    likeUsecase,
		Logger:         logger,
		ContextTimeout: timeout,
	}
}
