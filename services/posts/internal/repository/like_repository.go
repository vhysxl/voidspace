package repository

import (
	"context"
	"database/sql"
	"voidspace/posts/internal/domain"

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

// LikePost implements domain.LikeRepository.
func (l *likeRepository) LikePost(ctx context.Context, like *domain.Like) (int32, error) {
	tx, err := l.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

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
	defer tx.Rollback()

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
