package views

type UserBanner struct {
	ID          int    `db:"user_id"`
	Username    string `db:"username"`
	DisplayName string `db:"display_name"`
	AvatarURL   string `db:"avatar_url"`
}
