package models

import "time"

type GetUserRequest struct {
	Username string `json:"username" validate:"required,alphanum,min=3,max=30"`
}

type GetProfileRequest struct {
	ID       string `json:"id" validate:"required"`
	Username string `json:"username" validate:"required,alphanum,min=3,max=30"`
}

type FollowRequest struct {
	TargetUsername string `json:"target_username" validate:"required,alphanum,min=3,max=30"`
}

type UpdateProfileRequest struct {
	DisplayName string `json:"display_name" validate:"omitempty,max=50"`
	Bio         string `json:"bio" validate:"omitempty,max=160"`
	AvatarURL   string `json:"avatar_url" validate:"omitempty,url"`
	BannerURL   string `json:"banner_url" validate:"omitempty,url"`
	Location    string `json:"location" validate:"omitempty,max=100"`
}

type User struct {
	ID         int       `json:"id"`
	Username   string    `json:"username"`
	Profile    Profile   `json:"profile"`
	CreatedAt  time.Time `json:"created_at"`
	IsFollowed bool      `json:"is_followed"`
}

type Profile struct {
	DisplayName string `json:"display_name"`
	Bio         string `json:"bio"`
	AvatarURL   string `json:"avatar_url"`
	BannerURL   string `json:"banner_url"`
	Location    string `json:"location"`
	Followers   int    `json:"followers"`
	Following   int    `json:"following"`
}
