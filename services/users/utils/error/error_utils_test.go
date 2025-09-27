package error

import (
	"context"
	"errors"
	"testing"
	"voidspace/users/internal/domain"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type contextKey string

const (
	CtxKeyUserID   contextKey = "user_id"
	CtxKeyUsername contextKey = "username"
)

func TestHandleError(t *testing.T) {
	logger := zaptest.NewLogger(t)

	tests := []struct {
		name      string
		err       error
		wantCode  codes.Code
		wantError string
	}{
		{
			name:      "nil error",
			err:       nil,
			wantCode:  codes.OK,
			wantError: "",
		},
		{
			name:      "deadline exceeded",
			err:       context.DeadlineExceeded,
			wantCode:  codes.DeadlineExceeded,
			wantError: ErrRequestTimeout,
		},
		{
			name:      "user not found",
			err:       domain.ErrUserNotFound,
			wantCode:  codes.NotFound,
			wantError: domain.ErrUserNotFound.Error(),
		},
		{
			name:      "default error",
			err:       errors.New("unknown error"),
			wantCode:  codes.Internal,
			wantError: ErrInternalServer,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := HandleError(tt.err, logger, "test")
			if tt.err == nil {
				assert.Nil(t, err)
			} else {
				st, ok := status.FromError(err)
				assert.True(t, ok)
				assert.Equal(t, tt.wantCode, st.Code())
				assert.Equal(t, tt.wantError, st.Message())
			}
		})
	}
}

func TestHandleAuthError(t *testing.T) {
	logger := zaptest.NewLogger(t)

	tests := []struct {
		name      string
		userID    any
		wantError bool
	}{
		{
			name:      "valid user ID",
			userID:    123,
			wantError: false,
		},
		{
			name:      "nil user ID",
			userID:    nil,
			wantError: true,
		},
		{
			name:      "invalid user ID type",
			userID:    "invalid",
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := HandleAuthError(tt.userID, logger)
			if tt.wantError {
				assert.Error(t, err)
				st, ok := status.FromError(err)
				assert.True(t, ok)
				assert.Equal(t, codes.Unauthenticated, st.Code())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetUserIDFromContext(t *testing.T) {
	tests := []struct {
		name      string
		ctx       context.Context
		key       any
		wantID    int32
		wantError bool
	}{
		{
			name:      "valid user ID",
			ctx:       context.WithValue(context.Background(), CtxKeyUserID, 123),
			key:       CtxKeyUserID,
			wantID:    123,
			wantError: false,
		},
		{
			name:      "missing user ID",
			ctx:       context.Background(),
			key:       CtxKeyUserID,
			wantID:    0,
			wantError: true,
		},
		{
			name:      "invalid user ID type",
			ctx:       context.WithValue(context.Background(), CtxKeyUserID, "invalid"),
			key:       CtxKeyUserID,
			wantID:    0,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, err := GetUserIDFromContext(tt.ctx, tt.key)
			if tt.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantID, id)
			}
		})
	}
}
