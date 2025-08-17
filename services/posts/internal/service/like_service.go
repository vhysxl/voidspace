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
)

type LikeHandler struct {
	pb.UnimplementedPostServiceServer

	LikeUsecase    usecase.LikeUsecase
	Logger         *zap.Logger
	Validator      *validator.Validate
	ContextTimeout time.Duration
}

func NewLikeHandler(
	likeUsecase usecase.LikeUsecase,
	validator *validator.Validate,
	timeout time.Duration,
	logger *zap.Logger,
) *LikeHandler {
	return &LikeHandler{
		LikeUsecase:    likeUsecase,
		Validator:      validator,
		ContextTimeout: timeout,
		Logger:         logger,
	}
}

func (lh *LikeHandler) LikePost(ctx context.Context, req *pb.LikeRequest) (*pb.LikeResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, lh.ContextTimeout)
	defer cancel()

	userId, ok := ctx.Value(interceptor.CtxKeyUserID).(int)
	if !ok {
		lh.Logger.Error(ErrFailedGetUserID)
		return nil, status.Error(codes.Unauthenticated, ErrFailedGetUserID)
	}

	data := &domain.Like{
		UserID:    int32(userId),
		PostID:    req.PostId,
		CreatedAt: time.Now(),
	}

	err := lh.LikeUsecase.LikePost(ctx, data)
	if err != nil {
		lh.Logger.Error(ErrUsecase, zap.Error(err))
		switch err {
		case ctx.Err():
			return nil, status.Error(codes.DeadlineExceeded, ErrRequestTimeout)
		// case domain.ErrPostNotFound:
		// 	return nil, status.Error(codes.NotFound, err.Error())
		default:
			return nil, status.Error(codes.Internal, ErrInternalServer)
		}
	}

	return &pb.LikeResponse{
		Success: true,
	}, nil

}

func (lh *LikeHandler) UnlikePost(ctx context.Context, req *pb.LikeRequest) (*pb.LikeResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, lh.ContextTimeout)
	defer cancel()

	userId, ok := ctx.Value(interceptor.CtxKeyUserID).(int)
	if !ok {
		lh.Logger.Error(ErrFailedGetUserID)
		return nil, status.Error(codes.Unauthenticated, ErrFailedGetUserID)
	}

	data := &domain.Like{
		UserID:    int32(userId),
		PostID:    req.PostId,
		CreatedAt: time.Now(),
	}

	err := lh.LikeUsecase.UnlikePost(ctx, data)
	if err != nil {
		lh.Logger.Error(ErrUsecase, zap.Error(err))
		switch err {
		case ctx.Err():
			return nil, status.Error(codes.DeadlineExceeded, ErrRequestTimeout)
		// case domain.ErrPostNotFound:
		// 	return nil, status.Error(codes.NotFound, err.Error())
		default:
			return nil, status.Error(codes.Internal, ErrInternalServer)
		}
	}

	return &pb.LikeResponse{
		Success: true,
	}, nil

}
