package models

type LoginRequest struct {
	UsernameOrEmail string `json:"usernameoremail" validate:"required,min=3,max=50"`
	Password        string `json:"password" validate:"required"`
}

type RegisterRequest struct {
	Username string `json:"username" validate:"required,min=3,max=30,alphanum"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

// used in middleware
type AuthUser struct {
	ID       string
	Username string
}

// auth service generic response
type AuthResponseService struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
}

// ======================================== API RESPONSE ===================================
// Login Response
type LoginResponseAPI struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

// Register Response
type RegisterResponseAPI struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

// Refres Response
type RefreshTokenResponseAPI struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}
