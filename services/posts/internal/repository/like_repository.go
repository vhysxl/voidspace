package repository

import (
	"context"
	"database/sql"
	"log"
	"voidspace/posts/internal/domain"
	errUtil "voidspace/posts/utils/error"

	"github.com/lib/pq"
)

type likeRepository struct {
	db *sql.DB
}

func NewLikeRepository(db *sql.DB) domain.LikeRepository {
	return &likeRepository{
		db: db,
	}
}

// IsPostsLikedByUser implements domain.LikeRepository.
func (l *likeRepository) IsPostsLikedByUser(ctx context.Context, userID int32, postIDs []int32) (map[int32]bool, error) {
	// prep data
	result := make(map[int32]bool, len(postIDs))
	for _, id := range postIDs {
		result[id] = false
	}

	// query all
	rows, err := l.db.QueryContext(ctx, `
		SELECT post_id
		FROM post_likes
		WHERE user_id = $1 AND post_id = ANY($2)
	`, userID, pq.Array(postIDs))
	if err != nil {
		return nil, err
	}
	defer errUtil.SafeClose(rows)

	// mark found as true
	for rows.Next() {
		var postID int32
		if err := rows.Scan(&postID); err != nil {
			return nil, err
		}
		result[postID] = true
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

// IsPostLikedByUser implements domain.LikeRepository.
func (l *likeRepository) IsPostLikedByUser(ctx context.Context, userID int32, postID int32) (bool, error) {
	var dummy int
	err := l.db.QueryRowContext(ctx,
		`SELECT 1 FROM post_likes WHERE user_id = $1 AND post_id = $2`,
		userID, postID,
	).Scan(&dummy)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, err
	}
	return true, nil
}

// LikePost implements domain.LikeRepository.
func (l *likeRepository) LikePost(ctx context.Context, like *domain.Like) (int32, error) {
	tx, err := l.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}

	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Printf("failed to rollback : %v", err)
		}
	}()

	// Insert like (abaikan kalau sudah ada)
	result, err := tx.ExecContext(
		ctx,
		`INSERT INTO post_likes (user_id, post_id, created_at)
		 VALUES ($1, $2, $3)
		 ON CONFLICT (post_id, user_id) DO NOTHING`,
		like.UserID,
		like.PostID,
		like.CreatedAt,
	)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23503" {
			return 0, domain.ErrPostNotFound
		}
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	if rowsAffected == 0 {
		if err := tx.Commit(); err != nil {
			return 0, err
		}
		return 0, nil
	}

	var newLikesCount int32
	err = tx.QueryRowContext(
		ctx,
		`UPDATE posts
		 SET likes_count = likes_count + 1
		 WHERE id = $1
		 RETURNING likes_count`,
		like.PostID,
	).Scan(&newLikesCount)
	if err != nil {
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return newLikesCount, nil
}

// UnlikePost implements domain.LikeRepository.
func (l *likeRepository) UnlikePost(ctx context.Context, like *domain.Like) (int32, error) {
	tx, err := l.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}

	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Printf("failed to rollback : %v", err)
		}
	}()

	result, err := tx.ExecContext(
		ctx,
		`DELETE FROM post_likes WHERE user_id = $1 AND post_id = $2`,
		like.UserID,
		like.PostID,
	)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23503" {
			return 0, domain.ErrPostNotFound
		}
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	var newLikesCount int32
	if rowsAffected > 0 {
		err := tx.QueryRowContext(
			ctx,
			`UPDATE posts
             SET likes_count = likes_count - 1
             WHERE id = $1
             RETURNING likes_count`,
			like.PostID,
		).Scan(&newLikesCount)
		if err != nil {
			return 0, err
		}
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return newLikesCount, nil
}

func (l *likeRepository) DeleteAllLikes(ctx context.Context, userID int32) error {
	_, err := l.db.ExecContext(
		ctx,
		`DELETE FROM post_likes WHERE user_id = $1`,
		userID,
	)
	if err != nil {
		return err
	}

	return nil
}
