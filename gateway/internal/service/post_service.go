package service

import (
	"context"
	"fmt"
	"time"
	"voidspaceGateway/internal/models"
	postpb "voidspaceGateway/proto/generated/posts"
	userpb "voidspaceGateway/proto/generated/users"
	"voidspaceGateway/utils"

	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

type PostService struct {
	ContextTimeout time.Duration
	Logger         *zap.Logger
	PostClient     postpb.PostServiceClient
	UserClient     userpb.UserServiceClient
}

func NewPostService(timeout time.Duration, logger *zap.Logger, postClient postpb.PostServiceClient, userClient userpb.UserServiceClient) *PostService {
	return &PostService{
		ContextTimeout: timeout,
		Logger:         logger,
		PostClient:     postClient,
		UserClient:     userClient,
	}
}

func (ps *PostService) Create(ctx context.Context, username string, userID string, req *models.PostRequest) (*models.Post, error) {
	ctx, cancel := context.WithTimeout(ctx, ps.ContextTimeout)
	defer cancel()

	md := metadata.New(map[string]string{
		"user_id":  userID,
		"username": username,
	})

	ctx = metadata.NewOutgoingContext(ctx, md)

	data := &postpb.CreatePostRequest{
		Content:    req.Content,
		PostImages: req.PostImages,
	}

	res, err := ps.PostClient.CreatePost(ctx, data)
	if err != nil {
		ps.Logger.Error("failed to call PostService.Create", zap.Error(err))
		return nil, err
	}

	return &models.Post{
		ID:            int(res.GetId()),
		Content:       res.GetContent(),
		UserID:        int(res.GetUserId()),
		PostImages:    res.GetPostImages(),
		LikesCount:    int(res.GetLikesCount()),
		CommentsCount: int(res.GetCommentsCount()),
		CreatedAt:     res.GetCreatedAt().AsTime(),
		UpdatedAt:     res.GetUpdatedAt().AsTime(),
	}, nil
}

func (ps *PostService) GetPost(ctx context.Context, req *models.GetPostRequest, username string, userID string) (*models.Post, error) {
	ctx, cancel := context.WithTimeout(ctx, ps.ContextTimeout)
	defer cancel()

	md := metadata.New(map[string]string{
		"user_id":  userID,
		"username": username,
	})

	ctx = metadata.NewOutgoingContext(ctx, md)

	if md, ok := metadata.FromOutgoingContext(ctx); ok {
		fmt.Println("outgoing metadata:", md)
	} else {
		fmt.Println("no outgoing metadata")
	}

	postRes, err := ps.PostClient.GetPost(ctx, &postpb.GetPostRequest{
		Id: int32(req.ID),
	})
	if err != nil {
		ps.Logger.Error("Failed to call PostService.GetPost")
		return nil, err
	}

	userRes, err := ps.UserClient.GetUserById(ctx, &userpb.GetUserByIDRequest{
		UserID: postRes.UserId,
	})
	if err != nil {
		ps.Logger.Error("Failed to call UserService.GetUserByUserID")
		return nil, err
	}

	user := utils.UserMapper(userRes)

	post := utils.PostMapper(postRes, &user)
	return &post, nil
}

func (ps *PostService) Update(ctx context.Context, req *models.PostRequest, postID int, username string, userID string) error {
	ctx, cancel := context.WithTimeout(ctx, ps.ContextTimeout)
	defer cancel()

	md := metadata.New(map[string]string{
		"user_id":  userID,
		"username": username,
	})

	ctx = metadata.NewOutgoingContext(ctx, md)

	data := &postpb.UpdatePostRequest{
		Id:         int32(postID),
		Content:    req.Content,
		PostImages: req.PostImages,
	}

	_, err := ps.PostClient.UpdatePost(ctx, data)
	if err != nil {
		ps.Logger.Error("failed to call PostService.UpdatePost", zap.Error(err))
		return err
	}

	return nil
}

func (ps *PostService) Delete(ctx context.Context, postID int, username string, userID string) error {
	ctx, cancel := context.WithTimeout(ctx, ps.ContextTimeout)
	defer cancel()

	md := metadata.New(map[string]string{
		"user_id":  userID,
		"username": username,
	})

	ctx = metadata.NewOutgoingContext(ctx, md)

	data := &postpb.DeletePostRequest{
		Id: int32(postID),
	}

	_, err := ps.PostClient.DeletePost(ctx, data)
	if err != nil {
		ps.Logger.Error("failed to call PostService.DeletePost", zap.Error(err))
		return err
	}

	return nil
}

func (ps *PostService) GetUserPosts(ctx context.Context, username string) (*models.GetFeedResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, ps.ContextTimeout)
	defer cancel()

	user, err := ps.UserClient.GetUser(ctx, &userpb.GetUserRequest{
		Username: username,
	})
	if err != nil {
		ps.Logger.Error("failed to call UserService.GetUser", zap.Error(err))
		return nil, err
	}

	data := &postpb.GetAllPostsRequest{
		UserId: user.GetUser().GetId(),
	}

	res, err := ps.PostClient.GetAllPosts(ctx, data)
	if err != nil {
		ps.Logger.Error("failed to call PostService.GetAllPosts", zap.Error(err))
		return nil, err
	}

	posts := make([]models.Post, 0, len(res.GetPosts()))
	for _, post := range res.GetPosts() {
		posts = append(posts, models.Post{
			ID:            int(post.GetId()),
			Content:       post.GetContent(),
			UserID:        int(post.GetUserId()),
			PostImages:    post.GetPostImages(),
			CommentsCount: int(post.GetCommentsCount()),
			LikesCount:    int(post.GetLikesCount()),
			CreatedAt:     post.GetCreatedAt().AsTime(),
			UpdatedAt:     post.GetUpdatedAt().AsTime(),
		})
	}
	return &models.GetFeedResponse{
		Posts:   posts,
		HasMore: res.GetHasMore(),
	}, nil
}
