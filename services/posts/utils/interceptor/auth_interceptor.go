package interceptor

import (
	"context"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type ctxKey string

const (
	CtxKeyUserID   ctxKey = "userID"
	CtxKeyUsername ctxKey = "username"
)

func AuthInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {

		skipAuthMethods := map[string]bool{
			"/posts.v1.PostService/GetPost":       true,
			"/posts.v1.PostService/GetAllPosts":   true,
			"/posts.v1.PostService/GetGlobalFeed": true,
		}
		if skipAuthMethods[info.FullMethod] {
			return handler(ctx, req)
		}

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, "missing metadata")
		}

		//userid
		userIDArr := md.Get("user_id")
		if len(userIDArr) == 0 {
			return nil, status.Error(codes.Unauthenticated, "missing user_id in metadata")
		}
		userID, err := strconv.Atoi(userIDArr[0])
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "invalid user_id format")
		}

		//username
		usernameArr := md.Get("username")
		if len(usernameArr) == 0 {
			return nil, status.Error(codes.Unauthenticated, "missing username in metadata")
		}
		username := usernameArr[0]

		// Simpan ke context
		ctx = context.WithValue(ctx, CtxKeyUserID, userID)
		ctx = context.WithValue(ctx, CtxKeyUsername, username)

		return handler(ctx, req)
	}
}
