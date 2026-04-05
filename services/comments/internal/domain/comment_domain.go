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
	DeleteComment(ctx context.Context, commentID, userID int) error

	GetAllCommentsByPostID(ctx context.Context, postID int) (CommentRes, error)
	GetAllCommentsByUserID(ctx context.Context, userID int) (CommentRes, error)
	GetFeedCommentCount(ctx context.Context, postIDs []int) (map[int]int, error)

	HandleAccountDeletion(ctx context.Context, userID int) error
	HandleAccountRestoration(ctx context.Context, userID int) error
	DeleteByPostID(ctx context.Context, postID int) error
}

type CommentRepository interface {
	Create(ctx context.Context, comment *Comment) (*Comment, error)
	Delete(ctx context.Context, commentID int) (int, error)

	GetCommentByID(ctx context.Context, commentID int) (*Comment, error)
	GetAllByPostID(ctx context.Context, postID int) ([]*Comment, error)
	GetAllByUserID(ctx context.Context, userID int) ([]*Comment, error)
	CountCommentsByPostIDs(ctx context.Context, postIDs []int) (map[int]int, error)

	HandleAccountDeletion(ctx context.Context, userID int) error
	HandleAccountRestoration(ctx context.Context, userID int) error
	DeleteByPostID(ctx context.Context, postID int) error
}
