package repository

import (
	"context"
	"database/sql"
	"voidspace/posts/internal/domain"

	"github.com/lib/pq"
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
func (p *postRepository) Create(ctx context.Context, post *domain.Post) (*domain.Post, error) {
	err := p.db.QueryRowContext(
		ctx,
		`INSERT INTO posts (content, user_id, post_images)
     VALUES ($1, $2, $3) RETURNING id, content, user_id, post_images, created_at, updated_at`,
		post.Content,
		post.UserID,
		pq.Array(post.PostImages), // <- ini penting
	).Scan(
		&post.ID,
		&post.Content,
		&post.UserID,
		pq.Array(&post.PostImages),
		&post.CreatedAt,
		&post.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return post, err
}

// GetAllPosts implements domain.PostRepository.
func (p *postRepository) GetAllUserPosts(ctx context.Context, userID int32) ([]*domain.Post, error) {
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

// GetFollowFeed implements domain.PostRepository.
func (p *postRepository) GetFollowFeed(ctx context.Context, userIDs []int32, limit int32, offset int32) ([]*domain.Post, bool, error) {
	rows, err := p.db.QueryContext(
		ctx,
		`SELECT id, content, user_id, post_images, likes_count, created_at, updated_at
         FROM posts 
         WHERE user_id = ANY($1) 
         ORDER BY created_at DESC 
         LIMIT $2 OFFSET $3`,
		pq.Array(userIDs), limit+1, offset)
	if err != nil {
		return nil, false, err
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
			return nil, false, err
		}
		posts = append(posts, post)
	}

	hasMore := len(posts) > int(limit)
	if hasMore {
		posts = posts[:limit]
	}

	return posts, hasMore, nil
}

// GetGlobalFeed implements domain.PostRepository.
func (p *postRepository) GetGlobalFeed(ctx context.Context, limit, offset int32) ([]*domain.Post, bool, error) {
	rows, err := p.db.QueryContext(
		ctx,
		`SELECT id, content, user_id, post_images, likes_count, created_at, updated_at
         FROM posts 
         ORDER BY created_at DESC 
         LIMIT $1 OFFSET $2`,
		limit+1, offset)
	if err != nil {
		return nil, false, err
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
			return nil, false, err
		}
		posts = append(posts, post)
	}

	hasMore := len(posts) > int(limit)
	if hasMore {
		posts = posts[:limit]
	}

	return posts, hasMore, nil
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

// Update implements domain.PostRepository.
func (p *postRepository) Update(ctx context.Context, post *domain.Post) error {
	_, err := p.db.ExecContext(
		ctx,
		`UPDATE posts SET content = $1, post_images = $2, updated_at = $3
    WHERE id = $4`,
		post.Content,
		post.PostImages,
		post.UpdatedAt,
		post.ID,
	)
	return err
}
