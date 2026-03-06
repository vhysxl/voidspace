package domain

import "context"

type Profile struct {
	UserID      int
	DisplayName string
	Bio         string
	AvatarUrl   string
	BannerUrl   string
	Location    string
}

type ProfileUsecase interface {
	UpdateProfile(ctx context.Context, userID int, updates *Profile) error
}

type ProfileRepository interface {
	Update(ctx context.Context, userID int, profile *Profile) error
}
