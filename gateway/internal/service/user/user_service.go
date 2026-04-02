package user

import (
	"crypto/rsa"
	"time"

	commentpb "voidspaceGateway/proto/generated/comments/v1"
	postpb "voidspaceGateway/proto/generated/posts/v1"
	userpb "voidspaceGateway/proto/generated/users/v1"

	"go.temporal.io/sdk/client"
	"go.uber.org/zap"
)

type UserService struct {
	ContextTimeout  time.Duration
	Logger          *zap.Logger
	PublicKey       rsa.PublicKey
	UserClient      userpb.UserServiceClient
	PostClient      postpb.PostServiceClient
	CommentClient   commentpb.CommentServiceClient
	TemporalClient  client.Client
	TemporalService string
}

func NewUserService(
	timeout time.Duration,
	logger *zap.Logger,
	PublicKey rsa.PublicKey,
	userClient userpb.UserServiceClient,
	postClient postpb.PostServiceClient,
	commentClient commentpb.CommentServiceClient,
	temporalClient client.Client,
	temporalService string,
) *UserService {
	return &UserService{
		ContextTimeout:  timeout,
		Logger:          logger,
		PublicKey:       PublicKey,
		UserClient:      userClient,
		PostClient:      postClient,
		CommentClient:   commentClient,
		TemporalClient:  temporalClient,
		TemporalService: temporalService,
	}
}
