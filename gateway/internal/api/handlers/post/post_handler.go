package post

import (
	"time"
	post_service "voidspaceGateway/internal/service/post"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type PostHandler struct {
	ContextTimeout time.Duration
	Logger         *zap.Logger
	Validator      *validator.Validate
	PostService    *post_service.PostService
}

func NewPostHandler(
	timeout time.Duration,
	logger *zap.Logger,
	validator *validator.Validate,
	postService *post_service.PostService,
) *PostHandler {
	return &PostHandler{
		ContextTimeout: timeout,
		Logger:         logger,
		Validator:      validator,
		PostService:    postService,
	}
}
