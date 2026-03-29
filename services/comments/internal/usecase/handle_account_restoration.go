package usecase

import "context"

// HandleAccountRestoration implements [domain.CommentUsecase].
func (c *commentUsecase) HandleAccountRestoration(
	ctx context.Context, userID int) error {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	err := c.commentRepository.HandleAccountRestoration(ctx, userID)
	if err != nil {
		return err
	}

	return nil
}
