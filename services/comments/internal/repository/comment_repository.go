package repository

import (
	"context"
	"database/sql"
	"errors"
	"voidspace/comments/internal/domain"
)

type commentRepository struct {
	db *sql.DB
}

func NewCommentRepository(db *sql.DB) domain.CommentRepository {
	return &commentRepository{
		db: db,
	}
}

func (r *commentRepository) Create(ctx context.Context, comment *domain.Comment) (*domain.Comment, error) {
	result, err := r.db.ExecContext(
		ctx,
		`INSERT INTO comments (user_id, post_id, content) VALUES (?, ?, ?)`,
		comment.UserID,
		comment.PostID,
		comment.Content,
	)
	if err != nil {
		return nil, err
	}

	// Get the auto-incremented ID
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	comment.ID = int32(id)

	// Fetch created_at timestamp
	err = r.db.QueryRowContext(
		ctx,
		`SELECT created_at FROM comments WHERE id = ?`,
		comment.ID,
	).Scan(&comment.CreatedAt)
	if err != nil {
		return nil, err
	}

	return comment, nil
}

// Delete removes a comment by its ID
func (r *commentRepository) Delete(ctx context.Context, commentID int32) error {
	result, err := r.db.ExecContext(
		ctx,
		`DELETE FROM comments WHERE id = ?`,
		commentID,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return domain.ErrCommentsNotFound
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *commentRepository) GetCommentByID(ctx context.Context, commentID int32) (*domain.Comment, error) {
	comment := &domain.Comment{}

	err := r.db.QueryRowContext(
		ctx,
		`SELECT id, user_id, post_id, content, created_at FROM comments WHERE id = ?`,
		commentID,
	).Scan(
		&comment.ID,
		&comment.UserID,
		&comment.PostID,
		&comment.Content,
		&comment.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrCommentsNotFound
		}
		return nil, err
	}

	return comment, nil
}

// GetAllByPostID implements domain.CommentRepository.
func (r *commentRepository) GetAllByPostID(ctx context.Context, postID int32) ([]*domain.Comment, error) {
	rows, err := r.db.QueryContext(
		ctx,
		`SELECT id, user_id, post_id, content, created_at FROM comments WHERE post_id = ? ORDER BY created_at ASC`,
		postID,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrCommentsNotFound
		}
		return nil, err
	}
	defer rows.Close()

	var comments []*domain.Comment
	for rows.Next() {
		comment := &domain.Comment{}
		if err := rows.Scan(
			&comment.ID,
			&comment.UserID,
			&comment.PostID,
			&comment.Content,
			&comment.CreatedAt,
		); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}

// GetAllByUserID implements domain.CommentRepository.
func (r *commentRepository) GetAllByUserID(ctx context.Context, userID int32) ([]*domain.Comment, error) {
	rows, err := r.db.QueryContext(
		ctx,
		`SELECT id, user_id, post_id, content, created_at FROM comments WHERE user_id = ? ORDER BY created_at ASC`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []*domain.Comment
	for rows.Next() {
		comment := &domain.Comment{}
		if err := rows.Scan(
			&comment.ID,
			&comment.UserID,
			&comment.PostID,
			&comment.Content,
			&comment.CreatedAt,
		); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}

func (r *commentRepository) DeleteAllComments(ctx context.Context, userId int32) error {
	_, err := r.db.ExecContext(
		ctx,
		`DELETE FROM comments WHERE user_id = ?`,
		userId,
	)

	if err != nil {
		return err
	}

	return nil
}
