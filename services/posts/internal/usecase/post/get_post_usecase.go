package post

import (
	"context"
	"voidspace/posts/internal/domain"
)

// GetPost implements [domain.PostUsecase].
func (p *postUsecase) GetPost(
	ctx context.Context,
	postID int,
	loggedInUserID *int,
) (*domain.Post, error) {
	post, err := p.postRepository.GetByID(ctx, postID)
	if err != nil {
		return nil, err
	}

	if loggedInUserID != nil {
		liked, err := p.likeRepository.IsPostLikedByUser(ctx, *loggedInUserID, postID)
		if err != nil {
			return nil, err
		}

		post.IsLiked = liked
		post.IsOwner = post.UserID == *loggedInUserID
	} else {
		// Guest user
		post.IsLiked = false
		post.IsOwner = false
	}

	return post, nil
}
