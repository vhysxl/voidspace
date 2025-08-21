package usecase

import (
	"context"
	"time"
	"voidspace/posts/internal/domain"
)

type postUsecase struct {
	postRepository domain.PostRepository
	contextTimeout time.Duration
}

type PostUsecase interface {
	CreatePost(ctx context.Context, post *domain.Post) (*domain.Post, error)
	GetByID(ctx context.Context, id int32) (*domain.Post, error)
	GetAllUserPosts(ctx context.Context, userID int32) ([]*domain.Post, error)
	UpdatePost(ctx context.Context, post *domain.Post) error
	DeletePost(ctx context.Context, id int32, userID int32) error
	GetGlobalFeed(ctx context.Context, limit, offset int32) ([]*domain.Post, bool, error)
	GetFollowFeed(ctx context.Context, userIDs []int32, limit, offset int32) ([]*domain.Post, bool, error)
}

func NewPostUsecase(postRepository domain.PostRepository, contextTimeout time.Duration) PostUsecase {
	return &postUsecase{
		postRepository: postRepository,
		contextTimeout: contextTimeout,
	}
}

// CreatePost implements PostUsecase.
func (p *postUsecase) CreatePost(ctx context.Context, post *domain.Post) (*domain.Post, error) {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	return p.postRepository.Create(ctx, post)
}

// DeletePost implements PostUsecase.
func (p *postUsecase) DeletePost(ctx context.Context, id int32, userID int32) error {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	//TO DO:  Ensure the post belongs to the user

	return p.postRepository.Delete(ctx, id)
}

// GetAllUserPosts implements PostUsecase.
func (p *postUsecase) GetAllUserPosts(ctx context.Context, userID int32) ([]*domain.Post, error) {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	posts, err := p.postRepository.GetAllUserPosts(ctx, userID)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

// GetByID implements PostUsecase.
func (p *postUsecase) GetByID(ctx context.Context, id int32) (*domain.Post, error) {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	post, err := p.postRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return post, nil
}

// GetFollowFeed implements PostUsecase.
func (p *postUsecase) GetFollowFeed(ctx context.Context, userIDs []int32, limit int32, offset int32) ([]*domain.Post, bool, error) {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	posts, hasNext, err := p.postRepository.GetFollowFeed(ctx, userIDs, limit, offset)
	if err != nil {
		return nil, false, err
	}

	return posts, hasNext, nil
}

// GetGlobalFeed implements PostUsecase.
func (p *postUsecase) GetGlobalFeed(ctx context.Context, limit int32, offset int32) ([]*domain.Post, bool, error) {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	posts, hasNext, err := p.postRepository.GetGlobalFeed(ctx, limit, offset)
	if err != nil {
		return nil, false, err
	}
	return posts, hasNext, nil
}

// UpdatePost implements PostUsecase.
func (p *postUsecase) UpdatePost(ctx context.Context, post *domain.Post) error {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	//tODO: Ensure the post belongs to the user
	return p.postRepository.Update(ctx, post)
}
