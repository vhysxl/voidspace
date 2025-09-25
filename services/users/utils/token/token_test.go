package token

import (
	"crypto/rand"
	"crypto/rsa"
	"testing"
	"time"
	"voidspace/users/internal/domain"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccessToken(t *testing.T) {
	// Generate test private key
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	assert.NoError(t, err)

	testCases := []struct {
		name        string
		user        *domain.User
		expiry      time.Duration
		shouldError bool
	}{
		{
			name: "Valid token creation",
			user: &domain.User{
				Id:       1,
				Username: "testuser",
			},
			expiry:      time.Hour,
			shouldError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			token, err := CreateAccessToken(tc.user, privateKey, tc.expiry)
			if tc.shouldError {
				assert.Error(t, err)
				assert.Empty(t, token)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, token)
			}
		})
	}
}

func TestCreateRefreshToken(t *testing.T) {
	// Generate test private key
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	assert.NoError(t, err)

	testCases := []struct {
		name        string
		user        *domain.User
		expiry      time.Duration
		shouldError bool
	}{
		{
			name: "Valid token creation",
			user: &domain.User{
				Id:       1,
				Username: "testuser",
			},
			expiry:      time.Hour * 24,
			shouldError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			token, err := CreateRefreshToken(tc.user, privateKey, tc.expiry)
			if tc.shouldError {
				assert.Error(t, err)
				assert.Empty(t, token)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, token)
			}
		})
	}
}
