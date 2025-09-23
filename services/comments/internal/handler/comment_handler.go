package service

import (
	"context"
	"time"
	"voidspace/comments/internal/domain"
	pb "voidspace/comments/proto/generated/comments"
	utils "voidspace/comments/utils/common"
	errorutils "voidspace/comments/utils/error"
	"voidspace/comments/utils/interceptor"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CommentHandler struct {
	pb.UnimplementedCommentServiceServer

	ContextTimeout time.Duration
	Logger         *zap.Logger
	CommentUsecase domain.CommentUsecase
}

func NewCommentHandler(timeout time.Duration, logger *zap.Logger, commentUsecase domain.CommentUsecase) *CommentHandler {
	return &CommentHandler{
		CommentUsecase: commentUsecase,
		Logger:         logger,
		ContextTimeout: timeout,
	}
}

func (ch *CommentHandler) CreateComment(ctx context.Context, req *pb.CreateCommentRequest) (*pb.CommentResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, ch.ContextTimeout)
	defer cancel()

	userId, err := errorutils.GetUserIDFromContext(ctx, interceptor.CtxKeyUserID)
	if err != nil {
		return nil, errorutils.HandleAuthError(nil, ch.Logger)
	}

	data := &domain.Comment{
		UserID:  userId,
		PostID:  int(req.PostId),
		Content: req.Content,
	}

	comment, err := ch.CommentUsecase.CreateComment(ctx, data)
	if err != nil {
		return nil, errorutils.HandleError(err, ch.Logger, "CreateComment")
	}

	return utils.CommentMapper(comment), nil
}

func (ch *CommentHandler) DeleteComment(ctx context.Context, req *pb.DeleteCommentRequest) (*pb.DeleteCommentResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, ch.ContextTimeout)
	defer cancel()

	userId, err := errorutils.GetUserIDFromContext(ctx, interceptor.CtxKeyUserID)
	if err != nil {
		return nil, errorutils.HandleAuthError(nil, ch.Logger)
	}

	res, err := ch.CommentUsecase.DeleteComment(ctx, int(req.CommentId), userId)
	if err != nil {
		return nil, errorutils.HandleError(err, ch.Logger, "DeleteComment")
	}

	return &pb.DeleteCommentResponse{PostId: int32(res)}, nil
}

func (ch *CommentHandler) GetAllCommentsByPostId(ctx context.Context, req *pb.GetAllCommentsByPostIdRequest) (*pb.GetBatchCommentsResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, ch.ContextTimeout)
	defer cancel()

	comments, err := ch.CommentUsecase.GetAllCommentsByPostID(ctx, int(req.PostId))
	if err != nil {
		return nil, errorutils.HandleError(err, ch.Logger, "GetAllCommentsByPostId")
	}

	pbComments := make([]*pb.CommentResponse, 0, len(comments))
	for _, c := range comments {
		pbComments = append(pbComments, utils.CommentMapper(c))
	}

	return &pb.GetBatchCommentsResponse{
		Comments: pbComments,
	}, nil
}

func (ch *CommentHandler) GetAllCommentsByUserId(ctx context.Context, req *pb.GetAllCommentsByUserIdRequest) (*pb.GetBatchCommentsResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, ch.ContextTimeout)
	defer cancel()

	comments, err := ch.CommentUsecase.GetAllCommentsByUserID(ctx, int(req.UserId))
	if err != nil {
		return nil, errorutils.HandleError(err, ch.Logger, "GetAllCommentsByUserId")
	}

	pbComments := make([]*pb.CommentResponse, 0, len(comments))
	for _, c := range comments {
		pbComments = append(pbComments, utils.CommentMapper(c))
	}

	return &pb.GetBatchCommentsResponse{
		Comments: pbComments,
	}, nil
}

func (ch *CommentHandler) CountCommentsByPostId(ctx context.Context, req *pb.CountCommentsByPostIdRequest) (*pb.GetCommentsCountResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, ch.ContextTimeout)
	defer cancel()

	count, err := ch.CommentUsecase.CountCommentsByPostID(ctx, int(req.PostId))
	if err != nil {
		return nil, errorutils.HandleError(err, ch.Logger, "CountCommentsByPostId")
	}

	return &pb.GetCommentsCountResponse{
		CommentsCount: int32(count),
	}, nil
}

func (ch *CommentHandler) GetCommentsCountByPostIds(ctx context.Context, req *pb.GetCommentsCountByPostIdsRequest) (*pb.GetCommentsCountByPostIdsResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, ch.ContextTimeout)
	defer cancel()

	postIds := make([]int, len(req.PostIds))
	for i, postId := range req.PostIds {
		postIds[i] = int(postId)
	}

	countMap, err := ch.CommentUsecase.GetCommentsCountByPostIDs(ctx, postIds)
	if err != nil {
		return nil, errorutils.HandleError(err, ch.Logger, "GetCommentsCountByPostIds")
	}

	pbCountMap := make(map[int32]int32)
	for postId, count := range countMap {
		pbCountMap[int32(postId)] = int32(count)
	}

	return &pb.GetCommentsCountByPostIdsResponse{
		PostCommentsCount: pbCountMap,
	}, nil
}

func (ch *CommentHandler) AccountDeletionHandle(ctx context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, ch.ContextTimeout)
	defer cancel()

	userId, err := errorutils.GetUserIDFromContext(ctx, interceptor.CtxKeyUserID)
	if err != nil {
		return nil, errorutils.HandleAuthError(nil, ch.Logger)
	}

	err = ch.CommentUsecase.AccountDeletionHandle(ctx, userId)
	if err != nil {
		return nil, errorutils.HandleError(err, ch.Logger, "AccountDeletionHandle")
	}

	return &emptypb.Empty{}, nil
}
