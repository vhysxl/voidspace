package bootstrap

import (
	"context"
	"log"
	"time"
	"voidspaceGateway/config"
	"voidspaceGateway/internal/service"
	comment_service "voidspaceGateway/internal/service/comment"
	post_service "voidspaceGateway/internal/service/post"
	user_service "voidspaceGateway/internal/service/user"

	commentpb "voidspaceGateway/proto/generated/comments/v1"
	postpb "voidspaceGateway/proto/generated/posts/v1"
	userpb "voidspaceGateway/proto/generated/users/v1"

	"github.com/vhysxl/voidspace/shared/utils/helper"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type Application struct {
	Config          *config.Config
	ApiSecret       string
	ContextTimeout  time.Duration
	Validator       *validator.Validate
	Logger          *zap.Logger
	TemporalService *TemporalService
	UserService     *user_service.UserService
	PostService     *post_service.PostService
	UploadService   *service.UploadService
	CommentService  *comment_service.CommentService
}

func App() (*Application, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env not found, using fallbacks", err)
	}

	config := config.GetConfig()

	validator := validator.New()

	logger, err := helper.InitLogger()
	if err != nil {
		log.Println("Logger failed to load", err)
		return nil, err
	}

	temporalService, err := TemporalServiceInit(logger, config.TemporalPort)
	if err != nil {
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
	userService := user_service.NewUserService(
		time.Duration(config.ContextTimeout)*time.Second,
		logger,
		*config.PublicKey,
		userpb.NewUserServiceClient(userConn),
		postpb.NewPostServiceClient(postConn),
		commentpb.NewCommentServiceClient(commentConn),
		temporalService.Client,
		temporalService.Service,
	)

	postService := post_service.NewPostService(
		time.Duration(config.ContextTimeout)*time.Second,
		logger,
		userpb.NewUserServiceClient(userConn),
		postpb.NewPostServiceClient(postConn),
		commentpb.NewCommentServiceClient(commentConn),
		temporalService.Client,
		temporalService.Service,
	)

	commentService := comment_service.NewCommentService(
		time.Duration(config.ContextTimeout)*time.Second,
		logger,
		postpb.NewPostServiceClient(postConn),
		commentpb.NewCommentServiceClient(commentConn),
		userpb.NewUserServiceClient(userConn),
	)

	uploadService, err := service.NewUploadService(context.Background(), config.BucketName)
	if err != nil {
		panic(err)
	}

	// Register Activities
	logger.Info("Gateway Ready")

	return &Application{
		Config:          config,
		ApiSecret:       config.ApiSecret,
		ContextTimeout:  time.Duration(config.ContextTimeout) * time.Second,
		Validator:       validator,
		Logger:          logger,
		TemporalService: temporalService,
		UserService:     userService,
		PostService:     postService,
		UploadService:   uploadService,
		CommentService:  commentService,
	}, nil
}
