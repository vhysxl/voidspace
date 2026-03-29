package comment

import (
	"context"

	"github.com/georgysavva/scany/v2/pgxscan"
)

type commentCount struct {
	PostID int `db:"post_id"`
	Count  int `db:"count"`
}

// CountCommentsByPostIDs implements [domain.CommentRepository].
func (c *CommentRepository) CountCommentsByPostIDs(
	ctx context.Context,
	postIDs []int) (map[int]int, error) {

	if len(postIDs) == 0 {
		return map[int]int{}, nil
	}

	query := `
		SELECT post_id, COUNT(*) as count
		FROM comments 
		WHERE post_id = ANY($1) 
		AND deleted_at IS NULL 
		GROUP BY post_id
	`

	var rows []commentCount
	err := pgxscan.Select(ctx, c.db, &rows, query, postIDs)
	if err != nil {
		return nil, err
	}

	counts := make(map[int]int, len(postIDs))
	for _, row := range rows {
		counts[row.PostID] = row.Count
	}

	return counts, nil
}
