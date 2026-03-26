package utils

import (
	"voidspace/posts/internal/domain"
	pb "voidspace/posts/proto/generated/posts/v1"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func MapDomainPostToPb(p *domain.Post) *pb.Post {
	images := make([]*pb.PostImage, len(p.PostImages))
	for i, img := range p.PostImages {
		images[i] = &pb.PostImage{
			Url:    img.Url,
			Order:  int64(img.Order),
			Width:  int64(img.Width),
			Height: int64(img.Height),
		}
	}

	return &pb.Post{
		Id:            int64(p.ID),
		Content:       p.Content,
		UserId:        int64(p.UserID),
		Images:        images,
		LikesCount:    int64(p.LikesCount),
		CommentsCount: int64(p.CommentsCount),
		CreatedAt:     timestamppb.New(p.CreatedAt),
		UpdatedAt:     timestamppb.New(p.UpdatedAt),
		IsLiked:       p.IsLiked,
		IsOwner:       p.IsOwner,
	}
}

func MapPbPostImageToDomain(pbImages []*pb.PostImage) []domain.PostImage {
	images := make([]domain.PostImage, len(pbImages))
	for i, img := range pbImages {
		images[i] = domain.PostImage{
			Url:    img.GetUrl(),
			Order:  int(img.GetOrder()),
			Width:  int(img.GetWidth()),
			Height: int(img.GetHeight()),
		}
	}
	return images
}
