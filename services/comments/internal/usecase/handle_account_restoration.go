package usecase

import "context"

// HandleAccountRestoration implements [domain.CommentUsecase].
func (c *commentUsecase) HandleAccountRestoration(
	ctx context.Context, userID int) error {
		
	err := c.commentRepository.HandleAccountRestoration(ctx, userID)
	if err != nil {
		return err
	}

	return nil
}
