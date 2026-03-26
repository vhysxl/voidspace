package like

import (
	"time"
	"voidspace/posts/internal/domain"
)

type likeUsecase struct {
	likeRepository domain.LikeRepository
	contextTimeout time.Duration
}

func NewLikeUsecase(
	likeRepository domain.LikeRepository,
	contextTimeout time.Duration,
) domain.LikeUsecase {
	return &likeUsecase{
		likeRepository: likeRepository,
		contextTimeout: contextTimeout,
	}
}
