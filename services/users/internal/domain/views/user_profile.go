package views

import "time"

type UserProfile struct {
	ID          int       `db:"id"`
	Username    string    `db:"username"`
	DisplayName string    `db:"display_name"`
	Bio         string    `db:"bio"`
	AvatarURL   string    `db:"avatar_url"`
	BannerURL   string    `db:"banner_url"`
	Location    string    `db:"location"`
	Follower    int       `db:"follower"`
	Following   int       `db:"following"`
	CreatedAt   time.Time `db:"created_at"`
	IsFollowed  bool      `db:"is_followed"`
}
