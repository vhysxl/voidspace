package bootstrap

import (
	"context"
	"log"
	"time"
	"voidspaceGateway/config"
	"voidspaceGateway/internal/service"
	logger "voidspaceGateway/logger"
	commentpb "voidspaceGateway/proto/generated/comments"
	postpb "voidspaceGateway/proto/generated/posts"
	userpb "voidspaceGateway/proto/generated/users"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type Application struct {
	Config         *config.Config
	ApiSecret      string
	ContextTimeout time.Duration
	Validator      *validator.Validate
	Logger         *zap.Logger
	AuthService    *service.AuthService
	UserService    *service.UserService
	PostService    *service.PostService
	LikeService    *service.LikeService
	FeedService    *service.FeedService
	UploadService  *service.UploadService
	CommentService *service.CommentsService
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

	// gRPC Connections to microservices
	userConn, err := NewConn(config.UserServiceAddr, true)
	if err != nil {
		return nil, err
	}
	postConn, err := NewConn(config.PostServiceAddr, true)
	if err != nil {
		return nil, err
	}

	commentConn, err := NewConn(config.CommentServiceAddr, true)
	if err != nil {
		return nil, err
	}

	logger.Info("bucket name", zap.String("bucket", config.BucketName))

	// Services
	authService := service.NewAuthService(time.Duration(config.ContextTimeout)*time.Second, logger, userpb.NewAuthServiceClient(userConn), *config.PublicKey)
	userService := service.NewUserService(time.Duration(config.ContextTimeout)*time.Second, logger, userpb.NewUserServiceClient(userConn), postpb.NewPostServiceClient(postConn), commentpb.NewCommentServiceClient(commentConn))
	postService := service.NewPostService(time.Duration(config.ContextTimeout)*time.Second, logger, postpb.NewPostServiceClient(postConn), userpb.NewUserServiceClient(userConn), commentpb.NewCommentServiceClient(commentConn))
	likeService := service.NewLikeService(time.Duration(config.ContextTimeout)*time.Second, logger, postpb.NewLikesServiceClient(postConn))
	feedService := service.NewFeedService(time.Duration(config.ContextTimeout)*time.Second, logger, postpb.NewPostServiceClient(postConn), userpb.NewUserServiceClient(userConn), commentpb.NewCommentServiceClient(commentConn))
	commentService := service.NewCommentService(time.Duration(config.ContextTimeout)*time.Second, logger, postpb.NewPostServiceClient(postConn), userpb.NewUserServiceClient(userConn), commentpb.NewCommentServiceClient(commentConn))
	uploadService, err := service.NewUploadService(context.Background(), config.BucketName)
	if err != nil {
		panic(err)
	}

	logger.Info("Gateway Ready")

	return &Application{
		Config:         config,
		ApiSecret:      config.ApiSecret,
		ContextTimeout: time.Duration(config.ContextTimeout) * time.Second,
		Validator:      validator,
		Logger:         logger,
		AuthService:    authService,
		UserService:    userService,
		PostService:    postService,
		LikeService:    likeService,
		FeedService:    feedService,
		UploadService:  uploadService,
		CommentService: commentService,
	}, nil
}
