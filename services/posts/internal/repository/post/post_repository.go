package post

import (
	"voidspace/posts/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostRepository struct {
	db *pgxpool.Pool
}

func NewPostRepository(db *pgxpool.Pool) domain.PostRepository {
	return &PostRepository{
		db: db,
	}
}
