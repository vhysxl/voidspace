package service

import (
	"context"
	"time"
	"voidspace/posts/internal/domain"
	"voidspace/posts/internal/usecase"
	pb "voidspace/posts/proto/generated/posts"
	"voidspace/posts/utils/interceptor"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
