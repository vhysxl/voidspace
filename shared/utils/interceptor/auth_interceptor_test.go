package interceptor

import (
	"context"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func TestAuthInterceptor(t *testing.T) {
	tests := []struct {
		name          string
		method        string
		metadata      metadata.MD
		expectedError codes.Code
	}{
		{
			name:   "Skip auth for login",
			method: "/users.v1.AuthService/Login",
			metadata: metadata.New(map[string]string{
				"user_id":  "1",
				"username": "test",
			}),
			expectedError: codes.OK,
		},
		{
			name:   "Skip auth for register",
			method: "/users.v1.AuthService/Register",
			metadata: metadata.New(map[string]string{
				"user_id":  "1",
				"username": "test",
			}),
			expectedError: codes.OK,
		},
		{
			name:          "Missing metadata for protected endpoint",
			method:        "/users.v1.UserService/UpdateUser",
			metadata:      metadata.MD{},
			expectedError: codes.Unauthenticated,
		},
		{
			name:   "Invalid user_id format",
			method: "/users.v1.UserService/UpdateUser",
			metadata: metadata.New(map[string]string{
				"user_id":  "invalid",
				"username": "test",
			}),
			expectedError: codes.InvalidArgument,
		},
		{
			name:   "Missing username",
			method: "/users.v1.UserService/UpdateUser",
			metadata: metadata.New(map[string]string{
				"user_id": "1",
			}),
			expectedError: codes.Unauthenticated,
		},
		{
			name:   "Valid auth data",
			method: "/users.v1.UserService/UpdateUser",
			metadata: metadata.New(map[string]string{
				"user_id":  "1",
				"username": "test",
			}),
			expectedError: codes.OK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			interceptor := AuthInterceptor()
			ctx := metadata.NewIncomingContext(context.Background(), tt.metadata)

			info := &grpc.UnaryServerInfo{
				FullMethod: tt.method,
			}

			handler := func(ctx context.Context, req interface{}) (interface{}, error) {
				return nil, nil
			}

			_, err := interceptor(ctx, nil, info, handler)

			if tt.expectedError == codes.OK && err != nil {
				t.Errorf("expected no error, got %v", err)
				return
			}

			if tt.expectedError != codes.OK {
				if err == nil {
					t.Errorf("expected error with code %v, got no error", tt.expectedError)
					return
				}

				status, ok := status.FromError(err)
				if !ok {
					t.Errorf("expected grpc status error, got %v", err)
					return
				}

				if status.Code() != tt.expectedError {
					t.Errorf("expected error code %v, got %v", tt.expectedError, status.Code())
				}
			}
		})
	}
}
