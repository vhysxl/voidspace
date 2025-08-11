package domain

import "github.com/golang-jwt/jwt/v5"

type AccessTokenClaims struct {
	ID       string
	Username string
	jwt.RegisteredClaims
}

type RefreshTokenClaims struct {
	ID       string
	Username string
	jwt.RegisteredClaims
}
