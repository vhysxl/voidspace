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

func NewCommentUsecase(commentRepository domain.CommentRepository, contextTimeout time.Duration) domain.CommentUsecase {
	return &commentUsecase{
		commentRepository: commentRepository,
		contextTimeout:    contextTimeout,
	}
}

// CreateComment creates a new comment
func (c *commentUsecase) CreateComment(ctx context.Context, comment *domain.Comment) (*domain.Comment, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	return c.commentRepository.Create(ctx, comment)
}

// DeleteComment deletes a comment by ID (only owner can delete)
func (c *commentUsecase) DeleteComment(ctx context.Context, commentID, userID int) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	comment, err := c.commentRepository.GetCommentByID(ctx, commentID)
	if err != nil {
		return 0, err
	}

	if comment.UserID != userID {
		return 0, domain.ErrUnauthorizedAction
	}

	return c.commentRepository.Delete(ctx, commentID)
}

// AccountDeletionHandle removes all comments for a given user
func (c *commentUsecase) AccountDeletionHandle(ctx context.Context, userID int) error {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	return c.commentRepository.DeleteAllComments(ctx, userID)
}

// GetAllCommentsByPostID returns all comments for a post
func (c *commentUsecase) GetAllCommentsByPostID(ctx context.Context, postID int) ([]*domain.Comment, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	return c.commentRepository.GetAllByPostID(ctx, postID)
}

// GetAllCommentsByUserID returns all comments made by a user
func (c *commentUsecase) GetAllCommentsByUserID(ctx context.Context, userID int) ([]*domain.Comment, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	return c.commentRepository.GetAllByUserID(ctx, userID)
}

// CountCommentsByPostID returns total comments for a post
func (c *commentUsecase) CountCommentsByPostID(ctx context.Context, postID int) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	return c.commentRepository.CountCommentsByPostID(ctx, postID)
}

// GetCommentsCountByPostIDs returns total comments for multiple posts
func (c *commentUsecase) GetCommentsCountByPostIDs(ctx context.Context, postIDs []int) (map[int]int, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	return c.commentRepository.CountCommentsByPostIDs(ctx, postIDs)
}
