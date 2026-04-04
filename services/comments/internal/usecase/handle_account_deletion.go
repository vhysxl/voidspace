package usecase

import "context"

// HandleAccountDeletion implements [domain.CommentUsecase].
func (c *commentUsecase) HandleAccountDeletion(
	ctx context.Context,
	userID int) error {

	err := c.commentRepository.HandleAccountDeletion(ctx, userID)
	if err != nil {
		return err
	}

	return nil
}
