package views

import "time"

type UserProfile struct {
	Id          int32
	Username    string
	DisplayName string
	Bio         string
	AvatarUrl   string
	BannerUrl   string
	Location    string
	Followers   int32
	Following   int32
	CreatedAt   time.Time
	IsFollowed  bool
}
