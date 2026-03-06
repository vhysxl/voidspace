package domain

import (
	"context"
	"time"
)

type Comment struct {
	ID        int
	UserID    int
	PostID    int
	Content   string
	CreatedAt time.Time
}

type CommentUsecase interface {
	CreateComment(ctx context.Context, comment *Comment) (*Comment, error)
	AccountDeletionHandle(ctx context.Context, userID int) error
	DeleteComment(ctx context.Context, commentID, userID int) (int, error)
	GetAllCommentsByPostID(ctx context.Context, postID int) ([]*Comment, error)
	GetAllCommentsByUserID(ctx context.Context, userID int) ([]*Comment, error)
	CountCommentsByPostID(ctx context.Context, postID int) (int, error)
	GetCommentsCountByPostIDs(ctx context.Context, postIDs []int) (map[int]int, error)
}

type CommentRepository interface {
	Create(ctx context.Context, comment *Comment) (*Comment, error)
	Delete(ctx context.Context, commentID int) (int, error)
	DeleteAllComments(ctx context.Context, userID int) error
	GetCommentByID(ctx context.Context, commentID int) (*Comment, error)
	GetAllByPostID(ctx context.Context, postID int) ([]*Comment, error)
	GetAllByUserID(ctx context.Context, userID int) ([]*Comment, error)
	CountCommentsByPostID(ctx context.Context, postID int) (int, error)
	CountCommentsByPostIDs(ctx context.Context, postIDs []int) (map[int]int, error)
}
