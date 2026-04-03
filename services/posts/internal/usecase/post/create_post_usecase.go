package post

import (
	"context"
	"voidspace/posts/internal/domain"
)

// CreatePost implements [domain.PostUsecase].
func (p *postUsecase) CreatePost(
	ctx context.Context,
	post *domain.Post,
) (*domain.Post, error) {
	post, err := p.postRepository.Create(ctx, post)
	if err != nil {
		return nil, err
	}

	return post, nil
}
