package post

import "context"

// HandleAccountDeletion implements [domain.PostUsecase].
func (p *postUsecase) HandleAccountDeletion(
	ctx context.Context,
	userID int,
) error {
	err := p.postRepository.HandleAccountDeletion(ctx, userID)
	if err != nil {
		return err
	}

	return nil
}
