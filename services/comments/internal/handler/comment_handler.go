package handler

import (
	"time"
	"voidspace/comments/internal/domain"
	pb "voidspace/comments/proto/generated/comments/v1"

	"go.uber.org/zap"
)

type CommentHandler struct {
	pb.UnimplementedCommentServiceServer

	CommentUsecase domain.CommentUsecase
	Logger         *zap.Logger
	ContextTimeout time.Duration
}

func NewCommentHandler(
	timeout time.Duration,
	logger *zap.Logger,
	commentUsecase domain.CommentUsecase,
) pb.CommentServiceServer {
	return &CommentHandler{
		CommentUsecase: commentUsecase,
		Logger:         logger,
		ContextTimeout: timeout,
	}
}
