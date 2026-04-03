package post

import "context"

// HandleAccountRestoration implements [domain.PostUsecase].
func (p *postUsecase) HandleAccountRestoration(
	ctx context.Context,
	userID int,
) error {
	err := p.postRepository.HandleAccountRestoration(ctx, userID)
	if err != nil {
		return err
	}

	return nil
}
