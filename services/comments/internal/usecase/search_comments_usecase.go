package usecase

import (
	"context"
	"errors"
	"regexp"
	"voidspace/comments/internal/domain"
)

var searchRegex = regexp.MustCompile(`^[a-zA-Z0-9 \-_.' ,]{2,100}$`)

func (c *commentUsecase) SearchComments(
	ctx context.Context,
	query string,
) ([]*domain.Comment, error) {
	if !searchRegex.MatchString(query) {
		return nil, errors.New("invalid search query: must be 2-100 characters and only contain letters, digits, spaces, and -_.' ,")
	}

	return c.commentRepository.SearchComments(ctx, query)
}
