package user

import (
	"context"
	"errors"
	"regexp"
	"voidspace/users/internal/domain/views"
)

var searchRegex = regexp.MustCompile(`^[a-zA-Z0-9 \-_.' ,]{2,100}$`)

func (u *UserUsecase) SearchUsers(
	ctx context.Context,
	query string,
) ([]views.UserBanner, error) {
	if !searchRegex.MatchString(query) {
		return nil, errors.New("invalid search query: must be 2-100 characters and only contain letters, digits, spaces, and -_.' ,")
	}

	return u.userRepository.SearchUsers(ctx, query)
}
