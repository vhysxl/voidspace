package domain

import (
	"context"
	"time"
)

type Post struct {
	ID            int32
	Content       string
	UserID        int32
	PostImages    []string
	LikesCount    int32
	CommentsCount int32
	CreatedAt     time.Time
	UpdatedAt     time.Time
	IsLiked       bool
}

type PostRepository interface {
	Create(ctx context.Context, post *Post) (*Post, error)
	GetByID(ctx context.Context, id int32) (*Post, error)
	Update(ctx context.Context, post *Post) error
	IncrementCommentsCount(ctx context.Context, postId int) (int, error)
	DecrementCommentsCount(ctx context.Context, postId int) (int, error)
	Delete(ctx context.Context, id int32) error
	GetAllUserPosts(ctx context.Context, userID int32) ([]*Post, error)

	// Feed
	GetGlobalFeed(ctx context.Context, cursor time.Time, cursorID int32) ([]*Post, bool, error)
	GetFollowFeed(ctx context.Context, userIDs []int32, cursorTime time.Time, cursorID int32) ([]*Post, bool, error)
	DeleteAllPosts(ctx context.Context, userID int32) error
}
