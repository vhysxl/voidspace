package domain

import "context"

type Profile struct {
	UserId      int32
	DisplayName string
	Bio         string
	AvatarUrl   string
	BannerUrl   string
	Location    string
}

type ProfileUsecase interface {
	UpdateProfile(ctx context.Context, userId int32, updates *Profile) error
}

type ProfileRepository interface {
	GetProfileById(ctx context.Context, userId int32) (*Profile, error)
	Update(ctx context.Context, userId int32, profile *Profile) error
}
