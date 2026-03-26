package post

import "context"

// HandleAccountRestoration implements [domain.PostUsecase].
func (p *postUsecase) HandleAccountRestoration(ctx context.Context, userID int) error {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	err := p.postRepository.HandleAccountRestoration(ctx, userID)
	if err != nil {
		return err
	}

	return nil
}
