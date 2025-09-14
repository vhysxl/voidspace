package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"time"
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
	// Ensure PostImages is never nil, convert to empty slice if needed
	if post.PostImages == nil {
		post.PostImages = []string{}
	}

	jsonImages, err := json.Marshal(post.PostImages)
	if err != nil {
		return nil, err
	}

	// Use sql.NullString to handle potential NULL values from database
	var jsonRaw sql.NullString

	err = p.db.QueryRowContext(
		ctx,
		`INSERT INTO posts (content, user_id, post_images)
        VALUES ($1, $2, $3) RETURNING id, content, user_id, post_images, created_at, updated_at`,
		post.Content,
		post.UserID,
		jsonImages,
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

	// Handle post_images JSON field from RETURNING clause
	if jsonRaw.Valid {
		// JSON data exists, unmarshal it into the PostImages slice
		err = json.Unmarshal([]byte(jsonRaw.String), &post.PostImages)
		if err != nil {
			return nil, err
		}
	} else {
		post.PostImages = []string{}
	}

	return post, nil
}

// GetByID implements domain.PostRepository.
func (p *postRepository) GetByID(ctx context.Context, id int32) (*domain.Post, error) {
	post := &domain.Post{}
	// Use sql.NullString to handle NULL values from the database
	var jsonRaw sql.NullString

	err := p.db.QueryRowContext(
		ctx,
		`SELECT id, content, user_id, post_images, likes_count, created_at, updated_at
        FROM posts WHERE id = $1`,
		id,
	).Scan(
		&post.ID,
		&post.Content,
		&post.UserID,
		&jsonRaw,
		&post.LikesCount,
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrPostNotFound
		}
		return nil, err
	}

	// Handle post_images JSON field
	if jsonRaw.Valid {
		// JSON data exists, unmarshal it into the PostImages slice
		err = json.Unmarshal([]byte(jsonRaw.String), &post.PostImages)
		if err != nil {
			return nil, err
		}
	} else {
		// post_images is NULL in database, set empty slice as default
		post.PostImages = []string{}
	}

	return post, nil
}

// GetAllPosts implements domain.PostRepository.
func (p *postRepository) GetAllUserPosts(ctx context.Context, userID int32) ([]*domain.Post, error) {

	var jsonRaw sql.NullString

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
			&jsonRaw,
			&post.LikesCount,
			&post.CreatedAt,
			&post.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		if jsonRaw.Valid {
			err = json.Unmarshal([]byte(jsonRaw.String), &post.PostImages)
			if err != nil {
				return nil, err
			}
		} else {
			post.PostImages = []string{}
		}

		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

// GetFollowFeed implements domain.PostRepository.
func (p *postRepository) GetFollowFeed(ctx context.Context, userIDs []int32, cursorTime time.Time, cursorID int32) ([]*domain.Post, bool, error) {
	var jsonRaw sql.NullString

	query := `
        SELECT id, content, user_id, post_images, likes_count, created_at, updated_at
        FROM posts
        WHERE user_id = ANY($1)
          AND ((created_at < $2) OR (created_at = $2 AND id < $3))
        ORDER BY created_at DESC, id DESC
        LIMIT $4
    `

	rows, err := p.db.QueryContext(
		ctx,
		query,
		pq.Array(userIDs),
		cursorTime,
		cursorID,
		10+1, // ambil extra 1 buat cek hasMore
	)
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
			&jsonRaw,
			&post.LikesCount,
			&post.CreatedAt,
			&post.UpdatedAt,
		)
		if err != nil {
			return nil, false, err
		}

		if jsonRaw.Valid {
			if err := json.Unmarshal([]byte(jsonRaw.String), &post.PostImages); err != nil {
				return nil, false, err
			}
		} else {
			post.PostImages = []string{}
		}

		posts = append(posts, post)
	}

	fmt.Println(posts)

	hasMore := len(posts) > int(10)
	if hasMore {
		posts = posts[:10]
	}

	return posts, hasMore, nil
}

// GetGlobalFeed implements domain.PostRepository.
func (p *postRepository) GetGlobalFeed(ctx context.Context, cursorTime time.Time, cursorID int32) ([]*domain.Post, bool, error) {
	var jsonRaw sql.NullString

	query := `
		SELECT id, content, user_id, post_images, likes_count, created_at, updated_at
		FROM posts
		WHERE (created_at < $1)
		   OR (created_at = $1 AND id < $2)
		ORDER BY created_at DESC, id DESC
		LIMIT $3
	`

	rows, err := p.db.QueryContext(ctx, query, cursorTime, cursorID, 10+1)
	if err != nil {
		return nil, false, err
	}
	defer rows.Close()

	posts := []*domain.Post{}

	for rows.Next() {
		post := &domain.Post{}
		err := rows.Scan(
			&post.ID,
			&post.Content,
			&post.UserID,
			&jsonRaw,
			&post.LikesCount,
			&post.CreatedAt,
			&post.UpdatedAt,
		)
		if err != nil {
			return nil, false, err
		}
		if jsonRaw.Valid {
			if err := json.Unmarshal([]byte(jsonRaw.String), &post.PostImages); err != nil {
				return nil, false, err
			}
		} else {
			post.PostImages = []string{}
		}
		posts = append(posts, post)
	}

	hasMore := len(posts) > 10
	if hasMore {
		posts = posts[:10]
	}

	return posts, hasMore, nil
}

// Update implements domain.PostRepository.
func (p *postRepository) Update(ctx context.Context, post *domain.Post) error {
	if post.PostImages == nil {
		post.PostImages = []string{}
	}

	jsonImages, err := json.Marshal(post.PostImages)
	if err != nil {
		return err
	}
	res, err := p.db.ExecContext(
		ctx,
		`UPDATE posts SET content = $1, post_images = $2, updated_at = NOW()
    WHERE id = $3`,
		post.Content,
		jsonImages,
		post.ID,
	)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return domain.ErrPostNotFound
	}
	return err
}

// Delete implements domain.PostRepository.
func (p *postRepository) Delete(ctx context.Context, id int32) error {
	res, err := p.db.ExecContext(
		ctx,
		`DELETE FROM posts WHERE id = $1`,
		id,
	)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return domain.ErrPostNotFound
	}

	return err
}

func (p *postRepository) DeleteAllPosts(ctx context.Context, userID int32) error {
	_, err := p.db.ExecContext(
		ctx,
		`DELETE FROM posts WHERE user_id = $1`,
		userID,
	)
	if err != nil {
		return err
	}

	return nil
}
