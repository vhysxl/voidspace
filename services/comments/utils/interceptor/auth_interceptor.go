package interceptor

import (
	"context"
	"fmt"
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

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			fmt.Println("meta data not found")
		}

		skippedMethods := map[string]bool{ //change this later
			"/comments.v1.CommentService/GetAllCommentsByPostId":    true,
			"/comments.v1.CommentService/GetAllCommentsByUserId":    true,
			"/comments.v1.CommentService/CountCommentsByPostId":     true,
			"/comments.v1.CommentService/GetCommentsCountByPostIds": true,
		}

		isSkippedMethod := skippedMethods[info.FullMethod]

		// Metadata optional for skipped methods
		if isSkippedMethod {
			// try to get metadata
			if md != nil {
				if userIDArr := md.Get("user_id"); len(userIDArr) > 0 {
					if userID, err := strconv.Atoi(userIDArr[0]); err == nil {
						ctx = context.WithValue(ctx, CtxKeyUserID, userID)
					}
				}

				if usernameArr := md.Get("username"); len(usernameArr) > 0 {
					ctx = context.WithValue(ctx, CtxKeyUsername, usernameArr[0])
				}
			}
			return handler(ctx, req)
		}

		userIDArr := md.Get("user_id")
		if len(userIDArr) == 0 {
			return nil, status.Error(codes.Unauthenticated, "missing user_id in metadata")
		}

		userID, err := strconv.Atoi(userIDArr[0])
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "invalid user_id format")
		}

		usernameArr := md.Get("username")
		if len(usernameArr) == 0 {
			return nil, status.Error(codes.Unauthenticated, "missing username in metadata")
		}
		username := usernameArr[0]

		ctx = context.WithValue(ctx, CtxKeyUserID, userID)
		ctx = context.WithValue(ctx, CtxKeyUsername, username)

		return handler(ctx, req)
	}
}
