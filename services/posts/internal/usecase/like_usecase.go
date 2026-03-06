package usecase

import (
	"context"
	"time"
	"voidspace/posts/internal/domain"
)

type likeUsecase struct {
	likeRepository domain.LikeRepository
	contextTimeout time.Duration
}

func NewLikeUsecase(likeRepository domain.LikeRepository, contextTimeout time.Duration) domain.LikeUsecase {
	return &likeUsecase{
		likeRepository: likeRepository,
		contextTimeout: contextTimeout,
	}
}

// LikePost implements LikeUsecase.
func (l *likeUsecase) LikePost(ctx context.Context, like *domain.Like) (int32, error) {
	ctx, cancel := context.WithTimeout(ctx, l.contextTimeout)
	defer cancel()

	return l.likeRepository.LikePost(ctx, like)
}

// UnlikePost implements LikeUsecase.
func (l *likeUsecase) UnlikePost(ctx context.Context, like *domain.Like) (int32, error) {
	ctx, cancel := context.WithTimeout(ctx, l.contextTimeout)
	defer cancel()

	return l.likeRepository.UnlikePost(ctx, like)
}
