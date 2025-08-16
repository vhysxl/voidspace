package domain

import (
	"context"
	"time"
)

type FeedParams struct {
	Cursor string
	Limit  int32
}

type FeedResponse struct {
	Posts      []*Post
	NextCursor string
	HasMore    bool
}

type Post struct {
	ID         int32
	Content    string
	UserID     int32
	PostImages []string
	LikesCount int32
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type PostRepository interface {
	Create(ctx context.Context, post *Post) error
	GetByID(ctx context.Context, id int32) (*Post, error)
	Update(ctx context.Context, post *Post) error
	Delete(ctx context.Context, id int32) error
	GetAllPosts(ctx context.Context, userID int32) ([]*Post, error)

	// Feed methods
	GetGlobalFeed(ctx context.Context, params FeedParams) (*FeedResponse, error)
	GetFeedByUserIDs(ctx context.Context, userIDs []int32, params FeedParams) (*FeedResponse, error)
}
