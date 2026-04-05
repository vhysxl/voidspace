package post

import (
	"context"

	"voidspaceGateway/internal/models"
	postpb "voidspaceGateway/proto/generated/posts/v1"
	"voidspaceGateway/utils"

	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

func (s *PostService) Create(
	ctx context.Context,
	username, userID string,
	req *models.CreatePostRequest,
) (*postpb.Post, error) {
	ctx, cancel := context.WithTimeout(ctx, s.ContextTimeout)
	defer cancel()

	md := utils.MetaDataHandler(userID, username)
	ctx = metadata.NewOutgoingContext(ctx, md)

	postImages := make([]*postpb.PostImage, len(req.PostImages))

	for _, image := range req.PostImages {
		postImages = append(postImages, &postpb.PostImage{
			Url:    image.ImageURL,
			Order:  int64(image.Order),
			Width:  int64(image.Width),
			Height: int64(image.Height),
		})
	}

	data := &postpb.CreatePostRequest{
		Content: req.Content,
		Images:  postImages,
	}

	res, err := s.PostClient.CreatePost(ctx, data)
	if err != nil {
		s.Logger.Error("failed to call PostService.CreatePost", zap.Error(err))
		return nil, err
	}

	return res, nil
}
