package post

import (
	"context"
	"voidspace/posts/internal/domain"

	"github.com/vhysxl/voidspace/shared/utils/constants"
)

// UpdatePost implements [domain.PostUsecase].
func (p *postUsecase) UpdatePost(
	ctx context.Context,
	post *domain.Post,
	loggedInUserID int,
) error {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	existingPost, err := p.postRepository.GetByID(ctx, post.ID)
	if err != nil {
		return err
	}

	if existingPost.UserID != loggedInUserID {
		return constants.ErrUnauthorized
	}

	existingPost.Content = post.Content
	existingPost.PostImages = post.PostImages

	err = p.postRepository.Update(ctx, existingPost)
	if err != nil {
		return err
	}

	return nil
}
