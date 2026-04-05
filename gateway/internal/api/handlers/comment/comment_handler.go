package comment

import (
	"time"
	comment_service "voidspaceGateway/internal/service/comment"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type CommentHandler struct {
	ContextTimeout time.Duration
	Logger         *zap.Logger
	Validator      *validator.Validate
	CommentService *comment_service.CommentService
}

func NewCommentHandler(
	timeout time.Duration,
	logger *zap.Logger,
	validator *validator.Validate,
	commentService *comment_service.CommentService,
) *CommentHandler {
	return &CommentHandler{
		ContextTimeout: timeout,
		Logger:         logger,
		Validator:      validator,
		CommentService: commentService,
	}
}
