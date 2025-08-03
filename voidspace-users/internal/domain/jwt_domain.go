package domain

import "github.com/golang-jwt/jwt/v5"

type AccessTokenClaims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type RefreshTokenClaims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}
