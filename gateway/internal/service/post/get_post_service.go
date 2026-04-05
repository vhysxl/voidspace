package post

import (
	"context"

	"voidspaceGateway/internal/models"
	commentpb "voidspaceGateway/proto/generated/comments/v1"
	postpb "voidspaceGateway/proto/generated/posts/v1"
	userpb "voidspaceGateway/proto/generated/users/v1"
	"voidspaceGateway/utils"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/metadata"
)

func (s *PostService) GetPost(
	ctx context.Context,
	postID int64,
	username string,
	userID string,
) (*models.Post, error) {
	ctx, cancel := context.WithTimeout(ctx, s.ContextTimeout)
	defer cancel()

	md := utils.MetaDataHandler(userID, username)
	ctx = metadata.NewOutgoingContext(ctx, md)

	postRes, err := s.PostClient.GetPost(ctx, &postpb.GetPostRequest{
		PostId: postID,
	})
	if err != nil {
		s.Logger.Error("Failed to call PostService.GetPost", zap.Error(err))
		return nil, err
	}

	// Fetch user and comment count in parallel
	g, gCtx := errgroup.WithContext(ctx)

	var userRes *userpb.GetUserResponse
	g.Go(func() error {
		var err error
		userRes, err = s.UserClient.GetUserById(gCtx, &userpb.GetUserByIdRequest{
			UserId: postRes.GetUserId(),
		})
		if err != nil {
			s.Logger.Error("Failed to call UserService.GetUserById", zap.Error(err))
		}
		return err
	})

	var commentCount int64
	g.Go(func() error {
		commentRes, err := s.CommentClient.GetAllCommentsByPostId(gCtx, &commentpb.GetAllCommentsByPostIdRequest{
			PostId: postRes.GetId(),
		})
		if err != nil {
			s.Logger.Error("Failed to call CommentService.GetAllCommentsByPostId", zap.Error(err))
			return err
		}
		commentCount = commentRes.GetCommentsCount()
		return nil
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}

	author := utils.UserMapper(userRes.GetUser())

	return utils.PostMapper(postRes, author, int(commentCount)), nil
}
