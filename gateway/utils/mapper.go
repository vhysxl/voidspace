package utils

import (
	"voidspaceGateway/internal/models"
	postpb "voidspaceGateway/proto/generated/posts"
	userpb "voidspaceGateway/proto/generated/users"
)

// this is where all mapper stored, used across all services
func ProfileMapper(profile *userpb.GetUserResponse) models.Profile {
	if profile == nil {
		return models.Profile{}
	}

	return models.Profile{
		Bio:         profile.GetUser().GetProfile().GetBio(),
		DisplayName: profile.GetUser().GetProfile().GetDisplayName(),
		AvatarURL:   profile.GetUser().GetProfile().GetAvatarUrl(),
		BannerURL:   profile.GetUser().GetProfile().GetBannerUrl(),
		Location:    profile.GetUser().GetProfile().GetLocation(),
		Followers:   profile.GetUser().GetProfile().GetFollowers(),
		Following:   profile.GetUser().GetProfile().GetFollowing(),
	}
}

func UserMapper(user *userpb.GetUserResponse) models.User {
	if user == nil {
		return models.User{}
	}

	return models.User{
		ID:        user.GetUser().GetId(),
		Username:  user.GetUser().GetUsername(),
		CreatedAt: user.GetUser().GetCreatedAt().AsTime(),
		Profile:   ProfileMapper(user),
	}
}

func UserMapperFromUser(user *userpb.User) models.User {
	if user == nil {
		return models.User{}
	}

	return models.User{
		Username:  user.GetUsername(),
		CreatedAt: user.GetCreatedAt().AsTime(),
		Profile:   ProfileMapperFromProfile(user.GetProfile()),
	}
}

func ProfileMapperFromProfile(profile *userpb.Profile) models.Profile {
	if profile == nil {
		return models.Profile{}
	}

	return models.Profile{
		Bio:         profile.GetBio(),
		DisplayName: profile.GetDisplayName(),
		AvatarURL:   profile.GetAvatarUrl(),
		BannerURL:   profile.GetBannerUrl(),
		Location:    profile.GetLocation(),
		Followers:   profile.GetFollowers(),
		Following:   profile.GetFollowing(),
	}
}

func PostMapper(postRes *postpb.PostResponse, user *models.User) models.Post {
	if postRes == nil {
		return models.Post{}
	}
	return models.Post{
		ID:            int(postRes.GetId()),
		Content:       postRes.GetContent(),
		UserID:        int(postRes.GetUserId()),
		PostImages:    postRes.GetPostImages(),
		LikesCount:    int(postRes.GetLikesCount()),
		CommentsCount: int(postRes.GetCommentsCount()),
		CreatedAt:     postRes.GetCreatedAt().AsTime(),
		UpdatedAt:     postRes.GetUpdatedAt().AsTime(),
		IsLiked:       postRes.GetIsLiked(),
		Author:        user,
	}
}
