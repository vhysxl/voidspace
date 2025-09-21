package service

import (
	"context"
	"time"
	"voidspace/comments/internal/domain"
	"voidspace/comments/internal/usecase"
	pb "voidspace/comments/proto/generated/comments"
	utils "voidspace/comments/utils/common"
	"voidspace/comments/utils/interceptor"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CommentHandler struct {
	pb.UnimplementedCommentServiceServer

	ContextTimeout time.Duration
	Logger         *zap.Logger
	CommentUsecase usecase.CommentUsecase
}

func NewCommentService(timeout time.Duration, logger *zap.Logger, CommentUsecase usecase.CommentUsecase) *CommentHandler {
	return &CommentHandler{
		CommentUsecase: CommentUsecase,
		Logger:         logger,
		ContextTimeout: timeout,
	}
}

func (ch *CommentHandler) CreateComment(ctx context.Context, req *pb.CreateCommentRequest) (*pb.CommentResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, ch.ContextTimeout)
	defer cancel()

	userID, ok := ctx.Value(interceptor.CtxKeyUserID).(int)
	if !ok {
		ch.Logger.Error(ErrFailedGetUserID)
		return nil, status.Error(codes.Internal, ErrFailedGetUserID)
	}

	data := &domain.Comment{
		UserID:  int32(userID),
		PostID:  req.PostID,
		Content: req.Content,
	}

	comment, err := ch.CommentUsecase.CreateComment(ctx, data)
	if err != nil {
		ch.Logger.Error(ErrUsecase, zap.Error(err))
		switch err {
		case ctx.Err():
			return nil, status.Error(codes.DeadlineExceeded, ErrRequestTimeout)
		default:
			return nil, status.Error(codes.Internal, ErrInternalServer)
		}
	}

	return utils.CommentMapper(comment), nil
}

func (ch *CommentHandler) DeleteComment(ctx context.Context, req *pb.DeleteCommentRequest) (*pb.DeleteCommentResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, ch.ContextTimeout)
	defer cancel()

	userID, ok := ctx.Value(interceptor.CtxKeyUserID).(int)
	if !ok {
		ch.Logger.Error(ErrFailedGetUserID)
		return nil, status.Error(codes.Internal, ErrFailedGetUserID)
	}

	res, err := ch.CommentUsecase.DeleteComment(ctx, req.CommentId, int32(userID))
	if err != nil {
		ch.Logger.Error(ErrUsecase, zap.Error(err))
		switch err {
		case ctx.Err():
			return nil, status.Error(codes.DeadlineExceeded, ErrRequestTimeout)
		case domain.ErrCommentsNotFound:
			return nil, status.Error(codes.NotFound, err.Error())
		default:
			return nil, status.Error(codes.Internal, ErrInternalServer)
		}
	}

	return &pb.DeleteCommentResponse{PostId: int32(res)}, nil
}

func (ch *CommentHandler) GetAllCommentsByPostID(ctx context.Context, req *pb.GetAllCommentsByPostIDRequest) (*pb.GetBatchCommentsResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, ch.ContextTimeout)
	defer cancel()

	comments, err := ch.CommentUsecase.GetAllCommentsByPostID(ctx, req.PostId)
	if err != nil {
		ch.Logger.Error(ErrUsecase, zap.Error(err))
		switch err {
		case ctx.Err():
			return nil, status.Error(codes.DeadlineExceeded, ErrRequestTimeout)
		default:
			return nil, status.Error(codes.Internal, ErrInternalServer)
		}
	}

	var pbComments []*pb.CommentResponse // optimize pakai map
	for _, c := range comments {
		pbComments = append(pbComments, utils.CommentMapper(c))
	}

	return &pb.GetBatchCommentsResponse{
		Comments: pbComments,
	}, nil
}

func (ch *CommentHandler) GetAllCommentsByUserID(ctx context.Context, req *pb.GetAllCommentsByUserIDRequest) (*pb.GetBatchCommentsResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, ch.ContextTimeout)
	defer cancel()

	comments, err := ch.CommentUsecase.GetAllCommentsByUserID(ctx, req.UserId)
	if err != nil {
		ch.Logger.Error(ErrUsecase, zap.Error(err))
		switch err {
		case ctx.Err():
			return nil, status.Error(codes.DeadlineExceeded, ErrRequestTimeout)
		default:
			return nil, status.Error(codes.Internal, ErrInternalServer)
		}
	}

	var pbComments []*pb.CommentResponse
	for _, c := range comments {
		pbComments = append(pbComments, utils.CommentMapper(c))
	}

	return &pb.GetBatchCommentsResponse{
		Comments: pbComments,
	}, nil
}

func (ch *CommentHandler) AccountDeletionHandle(ctx context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, ch.ContextTimeout)
	defer cancel()

	userID, ok := ctx.Value(interceptor.CtxKeyUserID).(int)
	if !ok {
		ch.Logger.Error(ErrFailedGetUserID)
		return nil, status.Error(codes.Internal, ErrFailedGetUserID)
	}

	err := ch.CommentUsecase.AccountDeletionHandle(ctx, int32(userID))
	if err != nil {
		ch.Logger.Error(ErrUsecase, zap.Error(err))
		switch err {
		case ctx.Err():
			return nil, status.Error(codes.DeadlineExceeded, ErrRequestTimeout)
		default:
			return nil, status.Error(codes.Internal, ErrInternalServer)
		}
	}

	return &emptypb.Empty{}, nil
}
