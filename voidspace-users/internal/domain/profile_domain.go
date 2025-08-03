package domain

import "context"

type Profile struct {
	UserID      int    `json:"user_id"`
	DisplayName string `json:"display_name"`
	Bio         string `json:"bio"`
	AvatarUrl   string `json:"avatar_url"`
	BannerUrl   string `json:"banner_url"`
	Location    string `json:"location"`
}

type ProfileRepository interface {
	Update(ctx context.Context, userID int, profile *Profile) error
}
