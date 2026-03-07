package domain

import (
	"context"
	"time"
)

type Post struct {
	postID        int
	Content       string
	UserID        int
	PostImages    []string
	LikesCount    int
	CommentsCount int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	IsLiked       bool
}

type PostUsecase interface {
	CreatePost(ctx context.Context, post *Post) (*Post, error)
	GetByID(ctx context.Context, userID, postID int) (*Post, error)
	GetAllUserPosts(ctx context.Context, userID int) ([]Post, error)
	GetLikedPosts(ctx context.Context, userID int) ([]Post, error)
	UpdatePost(ctx context.Context, post *Post) error
	DeletePost(ctx context.Context, postID, userID int) error
	GetGlobalFeed(ctx context.Context, cursorTime *time.Time, cursorID, userID int) ([]Post, bool, error)
	GetFollowFeed(ctx context.Context, userIDs []int, cursorTime *time.Time, cursorID, userID int) ([]Post, bool, error)
	HandleAccountDeletion(ctx context.Context, userID int) error
}

type PostRepository interface {
	// CRUD operations
	Create(ctx context.Context, post *Post) (*Post, error)
	GetByID(ctx context.Context, postID int) (*Post, error)
	GetAllUserPosts(ctx context.Context, userID int) ([]Post, error)
	GetLikedPosts(ctx context.Context, userID int) ([]Post, error)
	Update(ctx context.Context, post *Post) error
	Delete(ctx context.Context, postID int) error
	DeleteAllPosts(ctx context.Context, userID int) error

	// Soft delete operations
	SoftDeletePost(ctx context.Context, postID int) error
	SoftDeletePosts(ctx context.Context, userID int) error
	RestorePosts(ctx context.Context, userID int) error

	// Feed operations
	GetGlobalFeed(ctx context.Context, cursorTime time.Time, cursorID int) ([]Post, bool, error)
	GetFollowFeed(ctx context.Context, userIDs []int, cursorTime time.Time, cursorID int) ([]Post, bool, error)
}
