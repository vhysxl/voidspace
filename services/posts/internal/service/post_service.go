package service

import (
	"context"
	"errors"
	"time"
	"voidspace/posts/internal/domain"
	"voidspace/posts/internal/usecase"
	pb "voidspace/posts/proto/generated/posts"
	"voidspace/posts/utils/interceptor"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type PostHandler struct {
	pb.UnimplementedPostServiceServer

	PostUsecase    usecase.PostUsecase
	Logger         *zap.Logger
	validator      *validator.Validate
	contextTimeout time.Duration
}

func NewPostHandler(
	postUsecase usecase.PostUsecase,
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

	userId, ok := ctx.Value(interceptor.CtxKeyUserID).(int)
	if !ok {
		ph.Logger.Error(ErrFailedGetUserID)
		return nil, status.Error(codes.Unauthenticated, ErrFailedGetUserID)
	}

	data := &domain.Post{
		UserID:     int32(userId),
		Content:    req.Content,
		PostImages: req.PostImages,
	}

	post, err := ph.PostUsecase.CreatePost(ctx, data)
	if err != nil {
		ph.Logger.Error(ErrUsecase, zap.Error(err))
		switch err {
		case ctx.Err():
			return nil, status.Error(codes.DeadlineExceeded, ErrRequestTimeout)
		default:
			return nil, status.Error(codes.Internal, ErrInternalServer)
		}
	}

	return &pb.PostResponse{
		Id:         post.ID,
		Content:    post.Content,
		UserId:     post.UserID,
		PostImages: post.PostImages,
		LikesCount: post.LikesCount,
		CreatedAt:  timestamppb.New(post.CreatedAt),
		UpdatedAt:  timestamppb.New(post.CreatedAt),
	}, nil
}

func (ph *PostHandler) GetPost(ctx context.Context, req *pb.GetPostRequest) (*pb.PostResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, ph.contextTimeout)
	defer cancel()

	post, err := ph.PostUsecase.GetByID(ctx, req.Id)
	if err != nil {
		ph.Logger.Error(ErrUsecase, zap.Error(err))
		switch err {
		case ctx.Err():
			return nil, status.Error(codes.DeadlineExceeded, ErrRequestTimeout)
		case domain.ErrPostNotFound:
			return nil, status.Error(codes.NotFound, err.Error())
		default:
			return nil, status.Error(codes.Internal, ErrInternalServer)
		}
	}

	return &pb.PostResponse{
		Id:         post.ID,
		Content:    post.Content,
		UserId:     post.UserID,
		PostImages: post.PostImages,
		LikesCount: post.LikesCount,
		CreatedAt:  timestamppb.New(post.CreatedAt),
		UpdatedAt:  timestamppb.New(post.UpdatedAt),
	}, nil
}

func (ph *PostHandler) UpdatePost(ctx context.Context, req *pb.UpdatePostRequest) (*emptypb.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, ph.contextTimeout)
	defer cancel()

	userId, ok := ctx.Value(interceptor.CtxKeyUserID).(int)
	if !ok {
		ph.Logger.Error(ErrFailedGetUserID)
		return nil, status.Error(codes.Unauthenticated, ErrFailedGetUserID)
	}

	data := &domain.Post{
		ID:         req.Id,
		UserID:     int32(userId),
		Content:    req.Content,
		PostImages: req.PostImages,
	}

	err := ph.PostUsecase.UpdatePost(ctx, data)
	if err != nil {
		ph.Logger.Error(ErrUsecase, zap.Error(err))
		switch err {
		case ctx.Err():
			return nil, status.Error(codes.DeadlineExceeded, ErrRequestTimeout)
		case domain.ErrPostNotFound:
			return nil, status.Error(codes.NotFound, err.Error())
		case domain.ErrUnauthorizedAction:
			return nil, status.Error(codes.PermissionDenied, err.Error())
		default:
			return nil, status.Error(codes.Internal, ErrInternalServer)
		}
	}

	return nil, err
}

func (ph *PostHandler) DeletePost(ctx context.Context, req *pb.DeletePostRequest) (*emptypb.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, ph.contextTimeout)
	defer cancel()

	userId, ok := ctx.Value(interceptor.CtxKeyUserID).(int)
	if !ok {
		ph.Logger.Error(ErrFailedGetUserID)
		return nil, status.Error(codes.Unauthenticated, ErrFailedGetUserID)
	}

	err := ph.PostUsecase.DeletePost(ctx, req.Id, int32(userId))
	if err != nil {
		ph.Logger.Error(ErrUsecase, zap.Error(err))
		switch err {
		case ctx.Err():
			return nil, status.Error(codes.DeadlineExceeded, ErrRequestTimeout)
		case domain.ErrPostNotFound:
			return nil, status.Error(codes.NotFound, err.Error())
		case domain.ErrUnauthorizedAction:
			return nil, status.Error(codes.PermissionDenied, err.Error())
		default:
			return nil, status.Error(codes.Internal, ErrInternalServer)
		}
	}

	return nil, err
}

func (ph *PostHandler) GetAllPosts(ctx context.Context, req *pb.GetAllPostsRequest) (*pb.GetFeedResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, ph.contextTimeout)
	defer cancel()

	posts, err := ph.PostUsecase.GetAllUserPosts(ctx, int32(req.UserId))
	if err != nil {
		ph.Logger.Error(ErrUsecase, zap.Error(err))
		switch err {
		case ctx.Err():
			return nil, status.Error(codes.DeadlineExceeded, ErrRequestTimeout)
		default:
			return nil, status.Error(codes.Internal, ErrInternalServer)
		}
	}

	var postResponses []*pb.PostResponse
	for _, post := range posts {
		postResponses = append(postResponses, &pb.PostResponse{
			Id:         post.ID,
			Content:    post.Content,
			UserId:     post.UserID,
			PostImages: post.PostImages,
			LikesCount: post.LikesCount,
			CreatedAt:  timestamppb.New(post.CreatedAt),
			UpdatedAt:  timestamppb.New(post.UpdatedAt),
		})
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

	posts, hasNext, err := ph.PostUsecase.GetGlobalFeed(ctx, cursorTime, cursorID)
	if err != nil {
		ph.Logger.Error(ErrUsecase, zap.Error(err))
		switch {
		case errors.Is(err, ctx.Err()):
			return nil, status.Error(codes.DeadlineExceeded, ErrRequestTimeout)
		default:
			return nil, status.Error(codes.Internal, ErrInternalServer)
		}
	}

	postResponses := make([]*pb.PostResponse, 0, len(posts))
	for _, post := range posts {
		postResponses = append(postResponses, &pb.PostResponse{
			Id:         post.ID,
			Content:    post.Content,
			UserId:     post.UserID,
			PostImages: post.PostImages,
			LikesCount: post.LikesCount,
			CreatedAt:  timestamppb.New(post.CreatedAt),
			UpdatedAt:  timestamppb.New(post.UpdatedAt),
		})
	}

	return &pb.GetFeedResponse{
		Posts:   postResponses,
		HasMore: hasNext,
	}, nil
}

func (ph *PostHandler) GetFeedByUserIDs(ctx context.Context, req *pb.GetFeedByUserIDsRequest) (*pb.GetFeedResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, ph.contextTimeout)
	defer cancel()

	// siapkan cursor
	var cursorTime *time.Time
	var cursorID *int32

	if req.Cursor != nil {
		t := req.Cursor.AsTime()
		cursorTime = &t
	}
	cursorID = &req.CursorID

	// ambil feed berdasarkan user_ids + cursor
	posts, hasMore, err := ph.PostUsecase.GetFollowFeed(ctx, req.UserIds, cursorTime, cursorID)
	if err != nil {
		ph.Logger.Error(ErrUsecase, zap.Error(err))
		switch {
		case errors.Is(err, ctx.Err()):
			return nil, status.Error(codes.DeadlineExceeded, ErrRequestTimeout)
		default:
			return nil, status.Error(codes.Internal, ErrInternalServer)
		}
	}

	// mapping ke response
	postResponses := make([]*pb.PostResponse, 0, len(posts))
	for _, post := range posts {
		postResponses = append(postResponses, &pb.PostResponse{
			Id:         post.ID,
			Content:    post.Content,
			UserId:     post.UserID,
			PostImages: post.PostImages,
			LikesCount: post.LikesCount,
			CreatedAt:  timestamppb.New(post.CreatedAt),
			UpdatedAt:  timestamppb.New(post.UpdatedAt),
		})
	}

	return &pb.GetFeedResponse{
		Posts:   postResponses,
		HasMore: hasMore,
	}, nil
}

func (ph *PostHandler) AccountDeletionHandle(ctx context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, ph.contextTimeout)
	defer cancel()

	userID, ok := ctx.Value(interceptor.CtxKeyUserID).(int)
	if !ok {
		ph.Logger.Error(ErrFailedGetUserID)
		return nil, status.Error(codes.Internal, ErrFailedGetUserID)
	}

	err := ph.PostUsecase.AccountDeletionHandle(ctx, int32(userID))
	if err != nil {
		ph.Logger.Error(ErrUsecase, zap.Error(err))
		switch err {
		case ctx.Err():
			return nil, status.Error(codes.DeadlineExceeded, ErrRequestTimeout)
		default:
			return nil, status.Error(codes.Internal, ErrInternalServer)
		}
	}

	return &emptypb.Empty{}, nil
}
