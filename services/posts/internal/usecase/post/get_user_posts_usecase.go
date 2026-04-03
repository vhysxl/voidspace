package post

import (
	"context"
	"voidspace/posts/internal/domain"
)

// GetUserPosts implements [domain.PostUsecase].
func (p *postUsecase) GetUserPosts(
	ctx context.Context,
	userID int,
	loggedInUserID *int,
) ([]domain.Post, error) {
	posts, err := p.postRepository.GetByUserID(ctx, userID)
	if err != nil {
		return []domain.Post{}, err
	}

	if len(posts) == 0 {
		return []domain.Post{}, nil
	}

	if loggedInUserID != nil {
		postIDs := make([]int, 0, len(posts))
		for _, post := range posts {
			postIDs = append(postIDs, post.ID)
		}

		likedMap, err := p.likeRepository.IsPostsLikedByUser(ctx, *loggedInUserID, postIDs)
		if err != nil {
			return nil, err
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

	return posts, nil
}
