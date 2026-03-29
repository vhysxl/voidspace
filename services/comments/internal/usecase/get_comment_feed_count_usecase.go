package usecase

import "context"

// GetFeedCommentCount implements [domain.CommentUsecase].
func (c *commentUsecase) GetFeedCommentCount(
	ctx context.Context,
	postIDs []int) (map[int]int, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	return c.commentRepository.CountCommentsByPostIDs(ctx, postIDs)
}
