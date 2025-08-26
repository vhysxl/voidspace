package service

import (
	"context"
	"time"
	"voidspaceGateway/internal/models"
	postpb "voidspaceGateway/proto/generated/posts"

	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

type LikeService struct {
	ContextTimeout time.Duration
	Logger         *zap.Logger
	LikeClient     postpb.LikesServiceClient
}

func NewLikeService(timeout time.Duration, logger *zap.Logger, likeClient postpb.LikesServiceClient) *LikeService {
	return &LikeService{
		ContextTimeout: timeout,
		Logger:         logger,
		LikeClient:     likeClient,
	}
}

func (ls *LikeService) Like(ctx context.Context, req *models.LikeRequest) (*models.LikeResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, ls.ContextTimeout)
	defer cancel()

	md := metadata.New(map[string]string{
		"user_id":  req.UserID,
		"username": req.Username,
	})

	ctx = metadata.NewOutgoingContext(ctx, md)

	res, err := ls.LikeClient.Like(ctx, &postpb.LikeRequest{
		PostId: int32(req.PostID),
	})
	if err != nil {
		ls.Logger.Error("failed to call LikeService.Create", zap.Error(err))
		return nil, err
	}

	return &models.LikeResponse{
		NewLikesCount: int(res.GetNewLikesCount()),
	}, nil
}

func (ls *LikeService) Unlike(ctx context.Context, req *models.LikeRequest) (*models.LikeResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, ls.ContextTimeout)
	defer cancel()

	md := metadata.New(map[string]string{
		"user_id":  req.UserID,
		"username": req.Username,
	})

	ctx = metadata.NewOutgoingContext(ctx, md)

	res, err := ls.LikeClient.Unlike(ctx, &postpb.LikeRequest{
		PostId: int32(req.PostID),
	})
	if err != nil {
		ls.Logger.Error("failed to call LikeService.Create", zap.Error(err))
		return nil, err
	}

	return &models.LikeResponse{
		NewLikesCount: int(res.GetNewLikesCount()),
	}, nil
}
