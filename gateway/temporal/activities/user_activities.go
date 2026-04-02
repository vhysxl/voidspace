package activities

// import (
// 	"context"
// 	"time"

// 	commentpb "voidspaceGateway/proto/generated/comments"
// 	postpb "voidspaceGateway/proto/generated/posts"
// 	userpb "voidspaceGateway/proto/generated/users"
// 	"voidspaceGateway/utils"

// 	"go.uber.org/zap"
// 	"google.golang.org/grpc/metadata"
// 	"google.golang.org/protobuf/types/known/emptypb"
// )

// type UserActivities struct {
// 	ContextTimeout time.Duration
// 	Logger         *zap.Logger
// 	UserClient     userpb.UserServiceClient
// 	PostClient     postpb.PostServiceClient
// 	CommentClient  commentpb.CommentServiceClient
// }

// func NewUserActivities(
// 	contextTimeout time.Duration,
// 	logger *zap.Logger,
// 	userClient userpb.UserServiceClient,
// 	postClient postpb.PostServiceClient,
// 	commentClient commentpb.CommentServiceClient,
// ) *UserActivities {
// 	return &UserActivities{
// 		ContextTimeout: contextTimeout,
// 		Logger:         logger,
// 		UserClient:     userClient,
// 		PostClient:     postClient,
// 		CommentClient:  commentClient,
// 	}
// }

// type DeleteUserReq struct {
// 	UserID   string
// 	Username string
// }

// func (ua *UserActivities) DeleteUserActivity(ctx context.Context, req DeleteUserReq) error {
// 	ua.Logger.Info(
// 		"Starting Delete User Activity",
// 		zap.String("userID", req.UserID),
// 		zap.String("username", req.Username))

// 	ctx, cancel := context.WithTimeout(ctx, ua.ContextTimeout)
// 	defer cancel()

// 	md := utils.MetaDataHandler(req.UserID, req.Username)
// 	ctx = metadata.NewOutgoingContext(ctx, md)

// 	_, err := ua.UserClient.DeleteUser(ctx, &emptypb.Empty{})
// 	if err != nil {
// 		ua.Logger.Error("Failed to delete user", zap.Error(err))
// 		return err
// 	}

// 	ua.Logger.Info("User deleted successfully")
// 	return nil
// }

// func (ua *UserActivities) DeleteUserPostsActivity(ctx context.Context, req DeleteUserReq) error {
// 	ua.Logger.Info(
// 		"Starting Delete User Posts Activity",
// 		zap.String("userID", req.UserID),
// 		zap.String("username", req.Username),
// 	)

// 	ctx, cancel := context.WithTimeout(ctx, ua.ContextTimeout)
// 	defer cancel()

// 	md := utils.MetaDataHandler(req.UserID, req.Username)
// 	ctx = metadata.NewOutgoingContext(ctx, md)

// 	_, err := ua.PostClient.AccountDeletionHandle(ctx, &emptypb.Empty{})
// 	if err != nil {
// 		ua.Logger.Error("Failed to delete user posts", zap.Error(err))
// 		return err
// 	}

// 	ua.Logger.Info("User posts deleted successfully")
// 	return nil
// }

// func (ua *UserActivities) DeleteUserCommentsActivity(ctx context.Context, req DeleteUserReq) error {
// 	ua.Logger.Info(
// 		"Starting Delete User Comments Activity",
// 		zap.String("userID", req.UserID),
// 		zap.String("username", req.Username),
// 	)

// 	ctx, cancel := context.WithTimeout(ctx, ua.ContextTimeout)
// 	defer cancel()

// 	md := utils.MetaDataHandler(req.UserID, req.Username)
// 	ctx = metadata.NewOutgoingContext(ctx, md)

// 	_, err := ua.CommentClient.AccountDeletionHandle(ctx, &emptypb.Empty{})
// 	if err != nil {
// 		ua.Logger.Error("Failed to delete user comments", zap.Error(err))
// 		return err
// 	}

// 	ua.Logger.Info("User comments deleted successfully")
// 	return nil
// }
