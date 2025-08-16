package repository

import (
	"context"
	"database/sql"
	"voidspace/posts/internal/domain"
)

type postRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) domain.PostRepository {
	return &postRepository{
		db: db,
	}
}

// Create implements domain.PostRepository.
func (p *postRepository) Create(ctx context.Context, post *domain.Post) error {
	_, err := p.db.ExecContext(
		ctx,
		`INSERT INTO posts
		 (content, user_id, post_images, created_at, updated_at)
		 VALUES ($1, $2, $3, $4, $5)`,
		post.Content,
		post.UserID,
		post.PostImages,
		post.CreatedAt,
		post.UpdatedAt,
	)

	return err
}

// Delete implements domain.PostRepository.
func (p *postRepository) Delete(ctx context.Context, id int32) error {
	_, err := p.db.ExecContext(
		ctx,
		`DELETE FROM posts WHERE id = $1`,
		id,
	)
	return err
}

// GetAllPosts implements domain.PostRepository.
func (p *postRepository) GetAllPosts(ctx context.Context, userID int32) ([]*domain.Post, error) {
	rows, err := p.db.QueryContext(
		ctx,
		`SELECT id, content, user_id, post_images, likes_count, created_at, updated_at
		 FROM posts WHERE user_id = $1 ORDER BY created_at DESC`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*domain.Post
	for rows.Next() {
		post := &domain.Post{}
		err := rows.Scan(
			&post.ID,
			&post.Content,
			&post.UserID,
			&post.PostImages,
			&post.LikesCount,
			&post.CreatedAt,
			&post.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

// GetByID implements domain.PostRepository.
func (p *postRepository) GetByID(ctx context.Context, id int32) (*domain.Post, error) {
	post := &domain.Post{}
	err := p.db.QueryRowContext(
		ctx,
		`SELECT id, content, user_id, post_images, likes_count, created_at, updated_at
		 FROM posts WHERE id = $1`,
		id,
	).Scan(
		&post.ID,
		&post.Content,
		&post.UserID,
		&post.PostImages,
		&post.LikesCount,
		&post.CreatedAt,
		&post.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return post, nil
}

// GetFeedByUserIDs implements domain.PostRepository.
func (p *postRepository) GetFeedByUserIDs(ctx context.Context, userIDs []int32, params domain.FeedParams) (*domain.FeedResponse, error) {
	panic("unimplemented")
}

// GetGlobalFeed implements domain.PostRepository.
func (p *postRepository) GetGlobalFeed(ctx context.Context, params domain.FeedParams) (*domain.FeedResponse, error) {
	panic("unimplemented")
}

// Update implements domain.PostRepository.
func (p *postRepository) Update(ctx context.Context, post *domain.Post) error {
	panic("unimplemented")
}
