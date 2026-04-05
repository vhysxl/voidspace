package post

import (
	"context"

	"voidspaceGateway/internal/models"
	postpb "voidspaceGateway/proto/generated/posts/v1"
	"voidspaceGateway/utils"

	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

func (ps *PostService) Update(ctx context.Context, req *models.CreatePostRequest, postID int64, username string, userID string) error {
	ctx, cancel := context.WithTimeout(ctx, ps.ContextTimeout)
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

	data := &postpb.UpdatePostRequest{
		PostId:  postID,
		Content: req.Content,
		Images:  postImages,
	}

	_, err := ps.PostClient.UpdatePost(ctx, data)
	if err != nil {
		ps.Logger.Error("failed to call PostService.UpdatePost", zap.Error(err))
		return err
	}

	return nil
}
