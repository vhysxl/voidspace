package domain

import (
	"context"
	"time"
)

type Comment struct {
	ID        int32
	UserID    int32
	PostID    int32
	Content   string
	CreatedAt time.Time
}

type CommentRepository interface {
	Create(ctx context.Context, comment *Comment) (*Comment, error)
	Delete(ctx context.Context, commentID int32) (int, error)
	DeleteAllComments(ctx context.Context, userId int32) error
	GetCommentByID(ctx context.Context, commentID int32) (*Comment, error)
	GetAllByPostID(ctx context.Context, postID int32) ([]*Comment, error)
	GetAllByUserID(ctx context.Context, userID int32) ([]*Comment, error)
}
