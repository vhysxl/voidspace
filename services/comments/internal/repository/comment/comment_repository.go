package comment

import (
	"voidspace/comments/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CommentRepository struct {
	db *pgxpool.Pool
}

func NewCommentRepository(db *pgxpool.Pool) domain.CommentRepository {
	return &CommentRepository{
		db: db,
	}
}
