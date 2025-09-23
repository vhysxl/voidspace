package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
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
	comment.ID = int(id)

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
func (r *commentRepository) Delete(ctx context.Context, commentID int) (int, error) {
	var postID int
	err := r.db.QueryRowContext(
		ctx,
		`SELECT post_id FROM comments WHERE id = ?`,
		commentID,
	).Scan(&postID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, domain.ErrCommentsNotFound
		}
		return 0, err
	}

	result, err := r.db.ExecContext(
		ctx,
		`DELETE FROM comments WHERE id = ?`,
		commentID,
	)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	if rowsAffected == 0 {
		return 0, domain.ErrCommentsNotFound
	}

	return postID, nil
}

func (r *commentRepository) GetCommentByID(ctx context.Context, commentID int) (*domain.Comment, error) {
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
func (r *commentRepository) GetAllByPostID(ctx context.Context, postID int) ([]*domain.Comment, error) {
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
func (r *commentRepository) GetAllByUserID(ctx context.Context, userID int) ([]*domain.Comment, error) {
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

func (r *commentRepository) DeleteAllComments(ctx context.Context, userID int) error {
	_, err := r.db.ExecContext(
		ctx,
		`DELETE FROM comments WHERE user_id = ?`,
		userID,
	)

	if err != nil {
		return err
	}

	return nil
}

// CountCommentsByPostID implements domain.CommentRepository.
func (r *commentRepository) CountCommentsByPostID(ctx context.Context, postID int) (int, error) {
	var count int
	err := r.db.QueryRowContext(
		ctx,
		`SELECT COUNT(*) FROM comments WHERE post_id = ?`,
		postID,
	).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// CountCommentsByPostIDs implements domain.CommentRepository.
func (r *commentRepository) CountCommentsByPostIDs(ctx context.Context, postIDs []int) (map[int]int, error) {
	if len(postIDs) == 0 {
		return make(map[int]int), nil
	}

	// Build placeholders for IN clause
	placeholders := make([]string, len(postIDs))
	args := make([]interface{}, len(postIDs))
	for i, postID := range postIDs {
		placeholders[i] = "?"
		args[i] = postID
	}

	query := fmt.Sprintf(`
		SELECT post_id, COUNT(*) as comment_count 
		FROM comments 
		WHERE post_id IN (%s) 
		GROUP BY post_id
	`, strings.Join(placeholders, ","))

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make(map[int]int)

	// Initialize all postIDs with 0 count
	for _, postID := range postIDs {
		result[postID] = 0
	}

	// Fill in the actual counts
	for rows.Next() {
		var postID, count int
		if err := rows.Scan(&postID, &count); err != nil {
			return nil, err
		}
		result[postID] = count
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
