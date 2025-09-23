package service

import (
	"context"
	"time"
	"voidspace/posts/internal/domain"
	pb "voidspace/posts/proto/generated/posts"
	utils "voidspace/posts/utils/common"
	errorutils "voidspace/posts/utils/error"
	"voidspace/posts/utils/interceptor"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
)

type PostHandler struct {
	pb.UnimplementedPostServiceServer
	PostUsecase    domain.PostUsecase
	Logger         *zap.Logger
	validator      *validator.Validate
	contextTimeout time.Duration
}

func NewPostHandler(
	postUsecase domain.PostUsecase,
	validator *validator.Validate,
	timeout time.Duration,
	logger *zap.Logger,
) *PostHandler {
	return &PostHandler{
		PostUsecase:    postUsecase,
		validator:      validator,
		contextTimeout: timeout,
		Logger:         logger,
	}
}

func (ph *PostHandler) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.PostResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, ph.contextTimeout)
	defer cancel()

	userId, err := errorutils.GetUserIDFromContext(ctx, interceptor.CtxKeyUserID)
	if err != nil {
		return nil, errorutils.HandleAuthError(userId, ph.Logger)
	}

	data := &domain.Post{
		UserID:     int32(userId),
		Content:    req.Content,
		PostImages: req.PostImages,
	}

	post, err := ph.PostUsecase.CreatePost(ctx, data)
	if err != nil {
		return nil, errorutils.HandleError(err, ph.Logger, "CreatePost")
	}

	return utils.PostMapper(post), nil
}

func (ph *PostHandler) GetPost(ctx context.Context, req *pb.GetPostRequest) (*pb.PostResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, ph.contextTimeout)
	defer cancel()

	userId, _ := ctx.Value(interceptor.CtxKeyUserID).(int)

	post, err := ph.PostUsecase.GetByID(ctx, int32(userId), req.Id)
	if err != nil {
		return nil, errorutils.HandleError(err, ph.Logger, "GetPost")
	}

	return utils.PostMapper(post), nil
}

func (ph *PostHandler) UpdatePost(ctx context.Context, req *pb.UpdatePostRequest) (*emptypb.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, ph.contextTimeout)
	defer cancel()

	userId, err := errorutils.GetUserIDFromContext(ctx, interceptor.CtxKeyUserID)
	if err != nil {
		return nil, errorutils.HandleAuthError(userId, ph.Logger)
	}

	data := &domain.Post{
		ID:         req.Id,
		UserID:     int32(userId),
		Content:    req.Content,
		PostImages: req.PostImages,
	}

	err = ph.PostUsecase.UpdatePost(ctx, data)
	if err != nil {
		return nil, errorutils.HandleError(err, ph.Logger, "UpdatePost")
	}

	return &emptypb.Empty{}, nil
}

func (ph *PostHandler) DeletePost(ctx context.Context, req *pb.DeletePostRequest) (*emptypb.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, ph.contextTimeout)
	defer cancel()

	userId, err := errorutils.GetUserIDFromContext(ctx, interceptor.CtxKeyUserID)
	if err != nil {
		return nil, errorutils.HandleAuthError(userId, ph.Logger)
	}

	err = ph.PostUsecase.DeletePost(ctx, req.Id, int32(userId))
	if err != nil {
		return nil, errorutils.HandleError(err, ph.Logger, "DeletePost")
	}

	return &emptypb.Empty{}, nil
}

func (ph *PostHandler) GetAllPosts(ctx context.Context, req *pb.GetAllPostsRequest) (*pb.GetFeedResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, ph.contextTimeout)
	defer cancel()

	posts, err := ph.PostUsecase.GetAllUserPosts(ctx, int32(req.UserId))
	if err != nil {
		return nil, errorutils.HandleError(err, ph.Logger, "GetAllPosts")
	}

	postResponses := make([]*pb.PostResponse, 0, len(posts))
	for _, post := range posts {
		postResponses = append(postResponses, utils.PostMapper(post))
	}

	return &pb.GetFeedResponse{Posts: postResponses}, nil
}

func (ph *PostHandler) GetGlobalFeed(ctx context.Context, req *pb.GetGlobalFeedRequest) (*pb.GetFeedResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, ph.contextTimeout)
	defer cancel()

	var cursorTime *time.Time
	var cursorID *int32

	if req.Cursor != nil {
		t := req.Cursor.AsTime()
		cursorTime = &t
	}
	if req.CursorID != 0 {
		cursorID = &req.CursorID
	}

	userId, _ := ctx.Value(interceptor.CtxKeyUserID).(int)

	posts, hasNext, err := ph.PostUsecase.GetGlobalFeed(ctx, cursorTime, cursorID, int32(userId))
	if err != nil {
		return nil, errorutils.HandleError(err, ph.Logger, "GetGlobalFeed")
	}

	postResponses := make([]*pb.PostResponse, 0, len(posts))
	for _, post := range posts {
		postResponses = append(postResponses, utils.PostMapper(post))
	}

	return &pb.GetFeedResponse{
		Posts:   postResponses,
		HasMore: hasNext,
	}, nil
}

func (ph *PostHandler) GetFeedByUserIDs(ctx context.Context, req *pb.GetFeedByUserIDsRequest) (*pb.GetFeedResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, ph.contextTimeout)
	defer cancel()

	var cursorTime *time.Time
	var cursorID *int32

	if req.Cursor != nil {
		t := req.Cursor.AsTime()
		cursorTime = &t
	}

	userId, err := errorutils.GetUserIDFromContext(ctx, interceptor.CtxKeyUserID)
	if err != nil {
		return nil, errorutils.HandleAuthError(userId, ph.Logger)
	}

	cursorID = &req.CursorID
	posts, hasMore, err := ph.PostUsecase.GetFollowFeed(ctx, req.UserIds, cursorTime, cursorID, int32(userId))
	if err != nil {
		return nil, errorutils.HandleError(err, ph.Logger, "GetFeedByUserIDs")
	}

	postResponses := make([]*pb.PostResponse, 0, len(posts))
	for _, post := range posts {
		postResponses = append(postResponses, utils.PostMapper(post))
	}

	return &pb.GetFeedResponse{
		Posts:   postResponses,
		HasMore: hasMore,
	}, nil
}

func (ph *PostHandler) AccountDeletionHandle(ctx context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, ph.contextTimeout)
	defer cancel()

	userId, err := errorutils.GetUserIDFromContext(ctx, interceptor.CtxKeyUserID)
	if err != nil {
		return nil, errorutils.HandleAuthError(userId, ph.Logger)
	}

	err = ph.PostUsecase.AccountDeletionHandle(ctx, int32(userId))
	if err != nil {
		return nil, errorutils.HandleError(err, ph.Logger, "AccountDeletionHandle")
	}

	return &emptypb.Empty{}, nil
}
