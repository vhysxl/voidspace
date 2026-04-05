package post

import (
	"time"
	commentpb "voidspaceGateway/proto/generated/comments/v1"
	postpb "voidspaceGateway/proto/generated/posts/v1"
	userpb "voidspaceGateway/proto/generated/users/v1"
	"go.temporal.io/sdk/client"
	"go.uber.org/zap"
)

type PostService struct {
	ContextTimeout  time.Duration
	Logger          *zap.Logger
	UserClient      userpb.UserServiceClient
	PostClient      postpb.PostServiceClient
	CommentClient   commentpb.CommentServiceClient
	TemporalClient  client.Client
	TemporalService string
}

func NewPostService(
	contextTimeout time.Duration,
	logger *zap.Logger,
	userClient userpb.UserServiceClient,
	postClient postpb.PostServiceClient,
	commentClient commentpb.CommentServiceClient,
	temporalClient client.Client,
	temporalService string,
) *PostService {
	return &PostService{
		ContextTimeout:  contextTimeout,
		Logger:          logger,
		UserClient:      userClient,
		PostClient:      postClient,
		CommentClient:   commentClient,
		TemporalClient:  temporalClient,
		TemporalService: temporalService,
	}
}
