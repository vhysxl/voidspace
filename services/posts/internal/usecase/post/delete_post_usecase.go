package post

import (
	"context"

	"github.com/vhysxl/voidspace/shared/utils/constants"
)

// DeletePost implements [domain.PostUsecase].
func (p *postUsecase) DeletePost(
	ctx context.Context,
	postID int,
	loggedInUserID int,
) error {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	post, err := p.postRepository.GetByID(ctx, postID)
	if err != nil {
		return err
	}

	if post.UserID != loggedInUserID {
		return constants.ErrUnauthorized
	}

	return p.postRepository.SoftDelete(ctx, postID)
}
