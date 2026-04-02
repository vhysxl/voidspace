package utils

import (
	"voidspaceGateway/internal/models"
	userpb "voidspaceGateway/proto/generated/users/v1"
)

// this is where all mapper stored, used across all services
func ProfileMapper(profile *userpb.UserProfile) *models.Profile {
	if profile == nil {
		return nil
	}

	return &models.Profile{
		Bio:         profile.GetBio(),
		DisplayName: profile.GetDisplayName(),
		AvatarURL:   profile.GetAvatarUrl(),
		BannerURL:   profile.GetBannerUrl(),
		Location:    profile.GetLocation(),
		Followers:   int(profile.GetFollowers()),
		Following:   int(profile.GetFollowing()),
	}
}

func UserMapper(user *userpb.UserProfile) *models.User {
	if user == nil {
		return nil
	}

	profile := ProfileMapper(user)
	return &models.User{
		ID:         int(user.GetId()),
		Username:   user.GetUsername(),
		CreatedAt:  user.GetCreatedAt().AsTime(),
		Profile:    *profile,
		IsFollowed: user.GetIsFollowed(),
	}
}

// func UserBannerMapper(user *userpb.UserBanner) *models.UserBanner {
// 	if user == nil {
// 		return nil
// 	}

// 	return &models.UserBanner{
// 		ID:          int(user.GetId()),
// 		Username:    user.GetUsername(),
// 		DisplayName: user.GetDisplayName(),
// 		AvatarURL:   user.GetAvatarUrl(),
// 	}
// }

// // func PostMapper(postRes *postpb.PostResponse, user *models.User, commentCount int) *models.Post {
// // 	if postRes == nil {
// // 		return nil
// // 	}
// // 	return &models.Post{
// // 		ID:            int(postRes.GetId()),
// // 		Content:       postRes.GetContent(),
// // 		UserID:        int(postRes.GetUserId()),
// // 		PostImages:    postRes.GetPostImages(),
// // 		LikesCount:    int(postRes.GetLikesCount()),
// // 		CommentsCount: commentCount,
// // 		CreatedAt:     postRes.GetCreatedAt().AsTime(),
// // 		UpdatedAt:     postRes.GetUpdatedAt().AsTime(),
// // 		IsLiked:       postRes.GetIsLiked(),
// // 		Author:        user,
// // 	}
// // }

// // func CommentMapper(commentRes *commentpb.CommentResponse, user *models.User) *models.Comment {
// // 	if commentRes == nil {
// // 		return nil
// // 	}

// // 	return &models.Comment{
// // 		CommentID: int(commentRes.GetId()),
// // 		PostID:    int(commentRes.GetPostId()),
// // 		Content:   commentRes.GetContent(),
// // 		Author:    user,
// // 		CreatedAt: commentRes.GetCreatedAt().AsTime(),
// // 	}
// // }

func AuthMapper(authRes *userpb.AuthResponse) *models.AuthResponseService {
	return &models.AuthResponseService{
		AccessToken:  authRes.GetAccessToken(),
		RefreshToken: authRes.GetRefreshToken(),
		ExpiresIn:    int64(authRes.GetExpiresIn()),
	}
}
