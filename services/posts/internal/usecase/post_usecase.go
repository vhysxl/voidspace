package usecase

import (
	"context"
	"math"
	"time"
	"voidspace/posts/internal/domain"

	"golang.org/x/sync/errgroup"
)

type postUsecase struct {
	postRepository domain.PostRepository
	contextTimeout time.Duration
	likeRepository domain.LikeRepository
}

type PostUsecase interface {
	CreatePost(ctx context.Context, post *domain.Post) (*domain.Post, error)
	GetByID(ctx context.Context, userID, id int32) (*domain.Post, error)
	GetAllUserPosts(ctx context.Context, userID int32) ([]*domain.Post, error)
	UpdatePost(ctx context.Context, post *domain.Post) error
	DeletePost(ctx context.Context, id int32, userID int32) error
	GetGlobalFeed(ctx context.Context, cursorTime *time.Time, cursorID *int32, userID int32) ([]*domain.Post, bool, error)
	GetFollowFeed(ctx context.Context, userIDs []int32, cursorTime *time.Time, cursorID *int32, userID int32) ([]*domain.Post, bool, error)
	AccountDeletionHandle(ctx context.Context, userId int32) error
}

func NewPostUsecase(postRepository domain.PostRepository, likeRepository domain.LikeRepository, contextTimeout time.Duration) PostUsecase {
	return &postUsecase{
		likeRepository: likeRepository,
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

// UpdatePost implements PostUsecase.
func (p *postUsecase) UpdatePost(ctx context.Context, post *domain.Post) error {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	existingPost, err := p.postRepository.GetByID(ctx, post.ID)
	if err != nil {
		return err
	}

	if post.UserID != existingPost.UserID {
		return domain.ErrUnauthorizedAction
	}

	return p.postRepository.Update(ctx, post)
}

// DeletePost implements PostUsecase.
func (p *postUsecase) DeletePost(ctx context.Context, id int32, userID int32) error {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	existingPost, err := p.postRepository.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if userID != existingPost.UserID {
		return domain.ErrUnauthorizedAction
	}

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
func (p *postUsecase) GetByID(ctx context.Context, userID, id int32) (*domain.Post, error) {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	post, err := p.postRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	post.IsLiked = false

	if userID > 0 {
		isLiked, err := p.likeRepository.IsPostLikedByUser(ctx, userID, id)
		if err != nil {
			return nil, err
		}

		post.IsLiked = isLiked
	}

	return post, nil
}

// GetGlobalFeed implements PostUsecase.
func (p *postUsecase) GetGlobalFeed(ctx context.Context, cursorTime *time.Time, cursorID *int32, userID int32) ([]*domain.Post, bool, error) {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	var timeVal time.Time
	var idVal int32
	// using pointer because it will default 0 if no pointer
	if cursorTime != nil && cursorID != nil && !cursorTime.IsZero() && *cursorID > 0 {
		timeVal = *cursorTime
		idVal = *cursorID
	} else {
		timeVal = time.Now()
		idVal = math.MaxInt32
	}

	posts, hasNext, err := p.postRepository.GetGlobalFeed(ctx, timeVal, idVal)
	if err != nil {
		return nil, false, err
	}

	if userID > 0 {
		postIDs := make([]int32, 0, len(posts))
		for _, post := range posts {
			postIDs = append(postIDs, int32(post.ID))
		}

		likes, err := p.likeRepository.IsPostsLikedByUser(ctx, userID, postIDs)
		if err != nil {
			return nil, false, err
		}

		for i, post := range posts {
			if liked, ok := likes[int32(post.ID)]; ok {
				posts[i].IsLiked = liked
			} else {
				posts[i].IsLiked = false
			}
		}
	}

	return posts, hasNext, nil
}

// GetFollowFeed implements PostUsecase.
func (p *postUsecase) GetFollowFeed(ctx context.Context, userIDs []int32, cursorTime *time.Time, cursorID *int32, userID int32) ([]*domain.Post, bool, error) {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	var timeVal time.Time
	var idVal int32
	if cursorTime != nil && cursorID != nil && !cursorTime.IsZero() && *cursorID > 0 {
		timeVal = *cursorTime
		idVal = *cursorID
	} else {
		timeVal = time.Now()
		idVal = math.MaxInt32
	}

	posts, hasNext, err := p.postRepository.GetFollowFeed(ctx, userIDs, timeVal, idVal)
	if err != nil {
		return nil, false, err
	}

	postIDs := make([]int32, 0, len(posts))
	for _, post := range posts {
		postIDs = append(postIDs, int32(post.ID))
	}

	likes, err := p.likeRepository.IsPostsLikedByUser(ctx, userID, postIDs)
	if err != nil {
		return nil, false, err
	}

	for i, post := range posts {
		if liked, ok := likes[int32(post.ID)]; ok {
			posts[i].IsLiked = liked
		} else {
			posts[i].IsLiked = false
		}
	}

	return posts, hasNext, nil
}

// AccountDeletionHandle implements PostUsecase.
func (p *postUsecase) AccountDeletionHandle(ctx context.Context, userId int32) error {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return p.postRepository.DeleteAllPosts(ctx, userId)
	})
	g.Go(func() error {
		return p.likeRepository.DeleteAllLikes(ctx, userId)
	})

	return g.Wait()
}
