package utils

import (
	"voidspaceGateway/internal/models"
	commentpb "voidspaceGateway/proto/generated/comments"
	postpb "voidspaceGateway/proto/generated/posts"
	userpb "voidspaceGateway/proto/generated/users"
)

// this is where all mapper stored, used across all services
func ProfileMapper(profile *userpb.GetUserResponse) *models.Profile {
	if profile == nil {
		return nil
	}

	return &models.Profile{
		Bio:         profile.GetUser().GetProfile().GetBio(),
		DisplayName: profile.GetUser().GetProfile().GetDisplayName(),
		AvatarURL:   profile.GetUser().GetProfile().GetAvatarUrl(),
		BannerURL:   profile.GetUser().GetProfile().GetBannerUrl(),
		Location:    profile.GetUser().GetProfile().GetLocation(),
		Followers:   int(profile.GetUser().GetProfile().GetFollowers()),
		Following:   int(profile.GetUser().GetProfile().GetFollowing()),
	}
}

func UserMapper(user *userpb.GetUserResponse) *models.User {
	if user == nil {
		return nil
	}

	profile := ProfileMapper(user)
	return &models.User{
		ID:        int(user.GetUser().GetId()),
		Username:  user.GetUser().GetUsername(),
		CreatedAt: user.GetUser().GetCreatedAt().AsTime(),
		Profile:   *profile,
	}
}

func UserMapperFromUser(user *userpb.User) *models.User {
	if user == nil {
		return nil
	}

	profile := ProfileMapperFromProfile(user.GetProfile())
	return &models.User{
		ID:        int(user.GetId()),
		Username:  user.GetUsername(),
		CreatedAt: user.GetCreatedAt().AsTime(),
		Profile:   *profile,
	}
}

func ProfileMapperFromProfile(profile *userpb.Profile) *models.Profile {
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

func PostMapper(postRes *postpb.PostResponse, user *models.User, commentCount int) *models.Post {
	if postRes == nil {
		return nil
	}
	return &models.Post{
		ID:            int(postRes.GetId()),
		Content:       postRes.GetContent(),
		UserID:        int(postRes.GetUserId()),
		PostImages:    postRes.GetPostImages(),
		LikesCount:    int(postRes.GetLikesCount()),
		CommentsCount: commentCount,
		CreatedAt:     postRes.GetCreatedAt().AsTime(),
		UpdatedAt:     postRes.GetUpdatedAt().AsTime(),
		IsLiked:       postRes.GetIsLiked(),
		Author:        user,
	}
}

func CommentMapper(commentRes *commentpb.CommentResponse, user *models.User) *models.Comment {
	if commentRes == nil {
		return nil
	}

	return &models.Comment{
		CommentID: int(commentRes.GetId()),
		PostID:    int(commentRes.GetPostId()),
		Content:   commentRes.GetContent(),
		Author:    user,
		CreatedAt: commentRes.GetCreatedAt().AsTime(),
	}
}

func AuthMapper(authRes *userpb.AuthResponse) *models.AuthResponse {
	return &models.AuthResponse{
		AccessToken:  authRes.GetAccessToken(),
		RefreshToken: authRes.GetRefreshToken(),
		ExpiresIn:    int64(authRes.GetExpiresIn()),
	}
}
