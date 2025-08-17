package repository

import (
	"context"
	"database/sql"
	"voidspace/posts/internal/domain"
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
func (l *likeRepository) LikePost(ctx context.Context, like *domain.Like) error {
	tx, err := l.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	result, err := tx.ExecContext(
		ctx,
		`INSERT INTO post_likes (user_id, post_id, created_at)
		VALUES ($1, $2, $3) ON CONFLICT (post_id, user_id) DO NOTHING`,
		like.UserID,
		like.PostID,
		like.CreatedAt,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected > 0 {
		_, err = tx.ExecContext(
			ctx,
			`UPDATE posts SET likes_count = likes_count + 1 WHERE id = $1`,
			like.PostID,
		)
		if err != nil {
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return err
}

// UnlikePost implements domain.LikeRepository.
func (l *likeRepository) UnlikePost(ctx context.Context, like *domain.Like) error {
	tx, err := l.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	result, err := tx.ExecContext(
		ctx,
		`DELETE FROM post_likes WHERE user_id = $1 AND post_id = $2`,
		like.UserID,
		like.PostID,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected > 0 {
		_, err = tx.ExecContext(
			ctx,
			`UPDATE posts SET likes_count = likes_count - 1 WHERE id = $1`,
			like.PostID,
		)
		if err != nil {
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return err
}
