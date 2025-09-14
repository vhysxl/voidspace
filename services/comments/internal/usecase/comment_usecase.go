package usecase

import (
	"context"
	"time"
	"voidspace/comments/internal/domain"
)

type commentUsecase struct {
	commentRepository domain.CommentRepository
	contextTimeout    time.Duration
}

// CommentUsecase defines the interface for comment-related business logic
type CommentUsecase interface {
	CreateComment(ctx context.Context, comment *domain.Comment) (*domain.Comment, error)
	DeleteComment(ctx context.Context, commentID, userID int32) error
	GetAllCommentsByPostID(ctx context.Context, postID int32) ([]*domain.Comment, error)
	GetAllCommentsByUserID(ctx context.Context, userID int32) ([]*domain.Comment, error)
}

// NewCommentUsecase creates a new CommentUsecase with DI
func NewCommentUsecase(commentRepository domain.CommentRepository, contextTimeout time.Duration) CommentUsecase {
	return &commentUsecase{
		commentRepository: commentRepository,
		contextTimeout:    contextTimeout,
	}
}

// Implementations
func (c *commentUsecase) CreateComment(ctx context.Context, comment *domain.Comment) (*domain.Comment, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	return c.commentRepository.Create(ctx, comment)
}

func (c *commentUsecase) DeleteComment(ctx context.Context, commentID, userID int32) error {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	comment, err := c.commentRepository.GetCommentByID(ctx, commentID)
	if err != nil {
		return err
	}

	if comment.UserID != userID {
		return domain.ErrUnauthorizedAction
	}

	return c.commentRepository.Delete(ctx, commentID)
}

func (c *commentUsecase) GetAllCommentsByPostID(ctx context.Context, postID int32) ([]*domain.Comment, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	return c.commentRepository.GetAllByPostID(ctx, postID)
}

func (c *commentUsecase) GetAllCommentsByUserID(ctx context.Context, userID int32) ([]*domain.Comment, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	return c.commentRepository.GetAllByUserID(ctx, userID)
}
