package post

import (
	"time"
	"voidspace/posts/internal/domain"
)

type postUsecase struct {
	postRepository domain.PostRepository
	likeRepository domain.LikeRepository
	contextTimeout time.Duration
}

func NewPostUsecase(
	postRepository domain.PostRepository,
	likeRepository domain.LikeRepository,
	contextTimeout time.Duration,
) domain.PostUsecase {
	return &postUsecase{
		postRepository: postRepository,
		likeRepository: likeRepository,
		contextTimeout: contextTimeout,
	}
}
