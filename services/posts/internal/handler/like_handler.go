package service

import (
	"context"
	"time"
	"voidspace/posts/internal/domain"
	pb "voidspace/posts/proto/generated/posts"
	errorutils "voidspace/posts/utils/error"
	"voidspace/posts/utils/interceptor"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type LikeHandler struct {
	pb.UnimplementedLikesServiceServer

	LikeUsecase    domain.LikeUsecase
	Logger         *zap.Logger
	Validator      *validator.Validate
	ContextTimeout time.Duration
}

func NewLikeHandler(
	likeUsecase domain.LikeUsecase,
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

func (lh *LikeHandler) Like(ctx context.Context, req *pb.LikeRequest) (*pb.LikeResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, lh.ContextTimeout)
	defer cancel()

	userId, err := errorutils.GetUserIDFromContext(ctx, interceptor.CtxKeyUserID)
	if err != nil {
		return nil, errorutils.HandleAuthError(nil, lh.Logger)
	}

	data := &domain.Like{
		UserID:    int32(userId),
		PostID:    req.PostId,
		CreatedAt: time.Now(),
	}

	likeCount, err := lh.LikeUsecase.LikePost(ctx, data)
	if err != nil {
		return nil, errorutils.HandleError(err, lh.Logger, "Like")
	}

	return &pb.LikeResponse{
		Success:       true,
		NewLikesCount: likeCount,
	}, nil
}

func (lh *LikeHandler) Unlike(ctx context.Context, req *pb.LikeRequest) (*pb.LikeResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, lh.ContextTimeout)
	defer cancel()

	userId, err := errorutils.GetUserIDFromContext(ctx, interceptor.CtxKeyUserID)
	if err != nil {
		return nil, errorutils.HandleAuthError(nil, lh.Logger)
	}

	data := &domain.Like{
		UserID: int32(userId),
		PostID: req.PostId,
	}

	likeCount, err := lh.LikeUsecase.UnlikePost(ctx, data)
	if err != nil {
		return nil, errorutils.HandleError(err, lh.Logger, "Unlike")
	}

	return &pb.LikeResponse{
		Success:       true,
		NewLikesCount: likeCount,
	}, nil
}
