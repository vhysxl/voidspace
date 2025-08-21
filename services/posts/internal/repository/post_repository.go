package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
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
func (p *postRepository) GetFollowFeed(ctx context.Context, userIDs []int32, limit int32, offset int32) ([]*domain.Post, bool, error) {
	//variable initializer
	var jsonRaw sql.NullString

	//query to get data
	//LIMIT + 1 to check hasMore, offset for skip rows
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
	//make sure the conn close after function done
	defer rows.Close()

	// slice initializer to store posts
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
		// marshall the jsonraw to PostImages, set [] if null
		if jsonRaw.Valid {
			err = json.Unmarshal([]byte(jsonRaw.String), &post.PostImages)
			if err != nil {
				return nil, false, err
			}
		} else {
			post.PostImages = []string{}
		}
		posts = append(posts, post)
	}

	// check has more compare it using limit
	hasMore := len(posts) > int(limit)
	if hasMore {
		//remove last element, basically just use 10 (9 index)
		posts = posts[:limit]
	}

	return posts, hasMore, nil
}

// GetGlobalFeed implements domain.PostRepository.
func (p *postRepository) GetGlobalFeed(ctx context.Context, limit, offset int32) ([]*domain.Post, bool, error) {
	var jsonRaw sql.NullString

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
			&jsonRaw,
			&post.LikesCount,
			&post.CreatedAt,
			&post.UpdatedAt,
		)
		if err != nil {
			return nil, false, err
		}
		if jsonRaw.Valid {
			err = json.Unmarshal([]byte(jsonRaw.String), &post.PostImages)
			if err != nil {
				return nil, false, err
			}
		} else {
			post.PostImages = []string{}
		}
		posts = append(posts, post)
	}

	hasMore := len(posts) > int(limit)
	if hasMore {
		posts = posts[:limit]
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
