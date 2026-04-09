package domain

import (
	"context"
	"time"
)

type Post struct {
	ID         int
	Content    string
	UserID     int
	PostImages []PostImage
	LikesCount int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	IsLiked    bool
	IsOwner    bool
}

type PostImage struct {
	Url    string
	Order  int
	Width  int
	Height int
}

type PostUsecase interface {
	// Single post operations
	CreatePost(ctx context.Context, post *Post) (*Post, error)
	GetPost(ctx context.Context, postID int, loggedInUserID *int) (*Post, error)
	UpdatePost(ctx context.Context, post *Post, loggedInUserID int) error
	DeletePost(ctx context.Context, postID int, loggedInUserID int) error

	// User posts operations
	GetUserPosts(ctx context.Context, userID int, loggedInUserID *int) ([]Post, error)
	GetLikedPosts(ctx context.Context, userID int, loggedInUserID *int) ([]Post, error)

	// Feed operations
	GetGlobalFeed(ctx context.Context, cursorTime *time.Time, cursorID int, loggedInUserID *int) ([]Post, bool, error)
	GetFollowingFeed(ctx context.Context, cursorTime *time.Time, cursorID int, loggedInUserID int, userIDs []int) ([]Post, bool, error)

	// Account lifecycle
	HandleAccountDeletion(ctx context.Context, userID int) error
	HandleAccountRestoration(ctx context.Context, userID int) error

	SearchPosts(ctx context.Context, query string) ([]Post, error)
}

type PostRepository interface {
	// Single post CRUD
	Create(ctx context.Context, post *Post) (*Post, error)
	GetByID(ctx context.Context, postID int) (*Post, error)
	Update(ctx context.Context, post *Post) error
	Delete(ctx context.Context, postID int) error

	// Bulk query operations
	GetByUserID(ctx context.Context, userID int) ([]Post, error)
	GetLikedByUserID(ctx context.Context, userID int) ([]Post, error)

	// Feed operations
	GetGlobalFeed(ctx context.Context, cursorTime time.Time, cursorID int) ([]Post, bool, error)
	GetFollowingFeed(ctx context.Context, userIDs []int, cursorTime time.Time, cursorID int) ([]Post, bool, error)

	// Account lifecycle (atomic operations with transaction)
	HandleAccountDeletion(ctx context.Context, userID int) error
	HandleAccountRestoration(ctx context.Context, userID int) error

	SearchPosts(ctx context.Context, query string) ([]Post, error)
}
