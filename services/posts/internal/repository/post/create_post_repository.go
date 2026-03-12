package post

import (
	"context"
	"encoding/json"
	"voidspace/posts/internal/domain"
)

// Create implements [domain.PostRepository].
func (p *PostRepository) Create(ctx context.Context, post *domain.Post) (*domain.Post, error) {
	if len(post.PostImages) == 0 {
		post.PostImages = nil
	}

	imagesJSON, err := json.Marshal(post.PostImages)
	if err != nil {
		return nil, err
	}

	var jsonRaw []byte

	err = p.db.QueryRow(
		ctx,
		`INSERT INTO posts (content, user_id, post_images)
        VALUES ($1, $2, $3) 
        RETURNING id, content, user_id, post_images, created_at, updated_at`,
		post.Content,
		post.UserID,
		imagesJSON,
	).Scan(
		&post.ID,
		&post.Content,
		&post.UserID,
		&jsonRaw,
		&post.CreatedAt,
		&post.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return post, nil
}
