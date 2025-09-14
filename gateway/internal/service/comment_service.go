package service

import (
	"context"
	"time"
	"voidspaceGateway/internal/models"
	commentpb "voidspaceGateway/proto/generated/comments"
	postpb "voidspaceGateway/proto/generated/posts"
	userpb "voidspaceGateway/proto/generated/users"
	"voidspaceGateway/utils"

	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

type CommentsService struct {
	ContextTimeout time.Duration
	Logger         *zap.Logger
	PostClient     postpb.PostServiceClient
	UserClient     userpb.UserServiceClient
	CommentClient  commentpb.CommentServiceClient
}

func NewCommentService(
	timeout time.Duration,
	logger *zap.Logger,
	postClient postpb.PostServiceClient,
	userClient userpb.UserServiceClient,
	CommentClient commentpb.CommentServiceClient) *CommentsService {
	return &CommentsService{
		ContextTimeout: timeout,
		Logger:         logger,
		PostClient:     postClient,
		UserClient:     userClient,
		CommentClient:  CommentClient,
	}
}

func (cs *CommentsService) Create(ctx context.Context, req *models.CreateCommentReq, userID, username string) (*models.Comments, error) {
	ctx, cancel := context.WithTimeout(ctx, cs.ContextTimeout)
	defer cancel()

	md := metadata.New(map[string]string{
		"user_id":  userID,
		"username": username,
	})

	ctx = metadata.NewOutgoingContext(ctx, md)

	data := &commentpb.CreateCommentRequest{
		PostID:  req.PostID,
		Content: req.Content,
	}

	res, err := cs.CommentClient.CreateComment(ctx, data)
	if err != nil {
		cs.Logger.Error("failed to call CommentService.CreateComment", zap.Error(err))
		return nil, err
	}

	userRes, err := cs.UserClient.GetUserById(ctx, &userpb.GetUserByIDRequest{
		UserID: res.UserId,
	})
	if err != nil {
		cs.Logger.Error("failed to call Userservice.GetUserById", zap.Error(err))
		return nil, err
	}

	user := utils.UserMapper(userRes)

	return &models.Comments{
		CommentID: res.Id,
		PostID:    res.PostId,
		Content:   res.Content,
		Author:    &user,
		CreatedAt: res.CreatedAt.AsTime(),
	}, nil
}

func (cs *CommentsService) Delete(ctx context.Context, commentID int32, userID, username string) error {
	ctx, cancel := context.WithTimeout(ctx, cs.ContextTimeout)
	defer cancel()

	md := metadata.New(map[string]string{
		"user_id":  userID,
		"username": username,
	})

	ctx = metadata.NewOutgoingContext(ctx, md)

	_, err := cs.CommentClient.DeleteComment(ctx, &commentpb.DeleteCommentRequest{
		CommentId: commentID,
	})
	if err != nil {
		cs.Logger.Error("failed to call CommentService.DeleteComment", zap.Error(err))
		return err
	}

	return nil
}

func (cs *CommentsService) GetAllByPostID(ctx context.Context, postID int32) ([]*models.Comments, error) {
	ctx, cancel := context.WithTimeout(ctx, cs.ContextTimeout)
	defer cancel()

	res, err := cs.CommentClient.GetAllCommentsByPostID(ctx, &commentpb.GetAllCommentsByPostIDRequest{
		PostId: postID,
	})
	if err != nil {
		cs.Logger.Error("failed to call CommentService.GetAllCommentsByPostID", zap.Error(err))
		return nil, err
	}

	userIDs := make([]int32, 0, len(res.Comments))
	for _, c := range res.Comments {
		userIDs = append(userIDs, c.UserId)
	}

	// Batch fetch user
	usersRes, err := cs.UserClient.GetUsersByIds(ctx, &userpb.GetUserByUserIDsRequest{
		UserID: userIDs,
	})
	if err != nil {
		cs.Logger.Error("failed to call UserService.GetUsersByIds", zap.Error(err))
		return nil, err
	}

	userMap := make(map[int32]models.User)
	for _, u := range usersRes.GetUsers() {
		userMap[u.GetId()] = *utils.UserMapperFromUser(u)
	}

	comments := make([]*models.Comments, 0, len(res.GetComments()))
	for _, c := range res.GetComments() {
		u := userMap[c.GetUserId()]

		comments = append(comments, &models.Comments{
			CommentID: c.GetId(),
			PostID:    c.GetPostId(),
			Content:   c.GetContent(),
			CreatedAt: c.CreatedAt.AsTime(),
			Author:    &u,
		})
	}

	return comments, nil
}

func (cs *CommentsService) GetAllByUserID(ctx context.Context, userID int32) ([]*models.Comments, error) {
	ctx, cancel := context.WithTimeout(ctx, cs.ContextTimeout)
	defer cancel()

	res, err := cs.CommentClient.GetAllCommentsByUserID(ctx, &commentpb.GetAllCommentsByUserIDRequest{
		UserId: userID,
	})
	if err != nil {
		cs.Logger.Error("failed to call CommentService.GetAllCommentsByUserID", zap.Error(err))
		return nil, err
	}

	userRes, err := cs.UserClient.GetUserById(ctx, &userpb.GetUserByIDRequest{
		UserID: userID,
	})
	if err != nil {
		cs.Logger.Error("failed to call Userservice.GetUserById", zap.Error(err))
		return nil, err
	}

	user := *utils.UserMapperFromUser(userRes.GetUser())

	comments := make([]*models.Comments, 0, len(res.GetComments()))
	for _, c := range res.GetComments() {
		comments = append(comments, &models.Comments{
			CommentID: c.GetId(),
			PostID:    c.GetPostId(),
			Content:   c.GetContent(),
			CreatedAt: c.CreatedAt.AsTime(),
			Author:    &user,
		})
	}

	return comments, nil
}
