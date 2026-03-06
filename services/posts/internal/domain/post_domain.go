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

type PostUsecase interface {
	CreatePost(ctx context.Context, post *Post) (*Post, error)
	GetByID(ctx context.Context, userID, id int32) (*Post, error)
	GetAllUserPosts(ctx context.Context, userID int32) ([]*Post, error)
	UpdatePost(ctx context.Context, post *Post) error
	DeletePost(ctx context.Context, id int32, userID int32) error
	GetGlobalFeed(ctx context.Context, cursorTime *time.Time, cursorID *int32, userID int32) ([]*Post, bool, error)
	GetFollowFeed(ctx context.Context, userIDs []int32, cursorTime *time.Time, cursorID *int32, userID int32) ([]*Post, bool, error)
	AccountDeletionHandle(ctx context.Context, userId int32) error
}

type PostRepository interface {
	Create(ctx context.Context, post *Post) (*Post, error)
	GetByID(ctx context.Context, id int32) (*Post, error)
	Update(ctx context.Context, post *Post) error
	Delete(ctx context.Context, id int32) error
	GetAllUserPosts(ctx context.Context, userID int32) ([]*Post, error)

	// Feed
	GetGlobalFeed(ctx context.Context, cursor time.Time, cursorID int32) ([]*Post, bool, error)
	GetFollowFeed(ctx context.Context, userIDs []int32, cursorTime time.Time, cursorID int32) ([]*Post, bool, error)
	DeleteAllPosts(ctx context.Context, userID int32) error
}
