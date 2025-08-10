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

type ProfileRepository interface {
	GetProfileById(ctx context.Context, userID int) (*Profile, error)
	Update(ctx context.Context, userID int, profile *Profile) error
}
