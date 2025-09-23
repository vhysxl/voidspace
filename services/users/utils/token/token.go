package token

import (
	"crypto/rsa"
	"strconv"
	"time"
	"voidspace/users/internal/domain"

	"github.com/golang-jwt/jwt/v5"
)

func CreateAccessToken(user *domain.User, privateKey *rsa.PrivateKey, expiry time.Duration) (string, error) {
	exp := time.Now().Add(expiry).Unix()
	claims := &domain.AccessTokenClaims{
		ID:       strconv.Itoa(int(user.Id)),
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(exp, 0)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	accessToken, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

func CreateRefreshToken(user *domain.User, privateKey *rsa.PrivateKey, expiry time.Duration) (string, error) {
	exp := time.Now().Add(expiry).Unix()
	claims := &domain.RefreshTokenClaims{
		ID:       strconv.Itoa(int(user.Id)),
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(exp, 0)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	refreshToken, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}
	return refreshToken, nil
}
