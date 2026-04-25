package post

import (
	"context"
	"errors"
	"regexp"
	"voidspace/posts/internal/domain"
)

var searchRegex = regexp.MustCompile(`^[a-zA-Z0-9 \-_.' ,]{2,100}$`)

func (p *postUsecase) SearchPosts(
	ctx context.Context,
	query string,
) ([]domain.Post, error) {
	if !searchRegex.MatchString(query) {
		return nil, errors.New("invalid search query: must be 2-100 characters and only contain letters, digits, spaces, and -_.' ,")
	}

	return p.postRepository.SearchPosts(ctx, query)
}
