package views

type UserProfile struct {
	ID          int    `json:"-"`
	Username    string `json:"username"`
	DisplayName string `json:"display_name"`
	Bio         string `json:"bio"`
	AvatarUrl   string `json:"avatar_url"`
	BannerUrl   string `json:"banner_url"`
	Location    string `json:"location"`
	//following
	//followers
}
