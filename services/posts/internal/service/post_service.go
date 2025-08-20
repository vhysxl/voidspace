package service

import (
	"time"
	"voidspace/posts/internal/usecase"
	pb "voidspace/posts/proto/generated/posts"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type PostHandler struct {
	pb.UnimplementedPostServiceServer

	PostUsecase    usecase.PostUsecase
	Logger         *zap.Logger
	validator      *validator.Validate
	contextTimeout time.Duration
}

func NewPostHandler(
	postUsecase usecase.PostUsecase,
	validator *validator.Validate,
	timeout time.Duration,
	logger *zap.Logger,
) *PostHandler {
	return &PostHandler{
		PostUsecase:    postUsecase,
		validator:      validator,
		contextTimeout: timeout,
		Logger:         logger,
	}
}
