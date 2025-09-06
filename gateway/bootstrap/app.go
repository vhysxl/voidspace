package bootstrap

import (
	"context"
	"log"
	"time"
	"voidspaceGateway/config"
	"voidspaceGateway/internal/service"
	logger "voidspaceGateway/loggger"
	postpb "voidspaceGateway/proto/generated/posts"
	userpb "voidspaceGateway/proto/generated/users"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Application struct {
	Config         *config.Config
	ContextTimeout time.Duration
	Validator      *validator.Validate
	Logger         *zap.Logger
	AuthService    *service.AuthService
	UserService    *service.UserService
	PostService    *service.PostService
	LikeService    *service.LikeService
	FeedService    *service.FeedService
	UploadService  *service.UploadService
}

func App() (*Application, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env not found, using fallbacks", err)
	}

	config := config.GetConfig()

	validator := validator.New()

	logger, err := logger.InitLogger()
	if err != nil {
		log.Println("Logger failed to load", err)
		return nil, err
	}

	// gRPC connections
	authConn, err := grpc.NewClient(config.UserServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	userConn, err := grpc.NewClient(config.UserServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	postConn, err := grpc.NewClient(config.PostServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	likeConn, err := grpc.NewClient(config.PostServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	logger.Info("bucket name", zap.String("bucket", config.BucketName))

	// Services
	authService := service.NewAuthService(time.Duration(config.ContextTimeout)*time.Second, logger, userpb.NewAuthServiceClient(authConn), *config.PublicKey)
	userService := service.NewUserService(time.Duration(config.ContextTimeout)*time.Second, logger, userpb.NewUserServiceClient(userConn), postpb.NewPostServiceClient(postConn))
	postService := service.NewPostService(time.Duration(config.ContextTimeout)*time.Second, logger, postpb.NewPostServiceClient(postConn), userpb.NewUserServiceClient(userConn))
	likeService := service.NewLikeService(time.Duration(config.ContextTimeout)*time.Second, logger, postpb.NewLikesServiceClient(likeConn))
	feedService := service.NewFeedService(time.Duration(config.ContextTimeout)*time.Second, logger, postpb.NewPostServiceClient(postConn), userpb.NewUserServiceClient(userConn))
	uploadService, err := service.NewUploadService(context.Background(), config.BucketName, config.GCSCredentialPath)
	if err != nil {
		panic(err)
	}

	logger.Info("Gateway Ready")

	return &Application{
		Config:         config,
		ContextTimeout: time.Duration(config.ContextTimeout) * time.Second,
		Validator:      validator,
		Logger:         logger,
		AuthService:    authService,
		UserService:    userService,
		PostService:    postService,
		LikeService:    likeService,
		FeedService:    feedService,
		UploadService:  uploadService,
	}, nil

}
