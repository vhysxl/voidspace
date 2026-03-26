package post

import (
	"context"
	"time"
	"voidspace/posts/internal/domain"
)

// GetGlobalFeed implements [domain.PostUsecase].
func (p *postUsecase) GetGlobalFeed(
	ctx context.Context,
	cursorTime *time.Time,
	cursorID int,
	loggedInUserID *int,
) ([]domain.Post, bool, error) {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	var cursor time.Time
	if cursorTime != nil {
		cursor = *cursorTime
	} else {
		cursor = time.Now()
	}

	posts, hasMore, err := p.postRepository.GetGlobalFeed(ctx, cursor, cursorID)
	if err != nil {
		return []domain.Post{}, false, err
	}

	if len(posts) == 0 {
		return []domain.Post{}, false, nil
	}

	if loggedInUserID != nil {
		postIDs := make([]int, 0, len(posts))
		for _, post := range posts {
			postIDs = append(postIDs, post.ID)
		}

		likedMap, err := p.likeRepository.IsPostsLikedByUser(ctx, *loggedInUserID, postIDs)
		if err != nil {
			return nil, false, err
		}

		for i := range posts {
			posts[i].IsLiked = likedMap[posts[i].ID]
			posts[i].IsOwner = posts[i].UserID == *loggedInUserID
		}
	} else {
		for i := range posts {
			posts[i].IsLiked = false
			posts[i].IsOwner = false
		}
	}

	return posts, hasMore, nil
}
