package bootstrap

import (
	"context"
	"log"
	"time"
	"voidspace/posts/config"
	"voidspace/posts/internal/domain"
	"voidspace/posts/internal/repository"
	"voidspace/posts/internal/repository/post"
	"voidspace/posts/internal/usecase"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	util_db "github.com/vhysxl/voidspace/shared/utils/database"
	"github.com/vhysxl/voidspace/shared/utils/helper"
	"go.uber.org/zap"
)

type Application struct {
	Config                 *config.Config
	ContextTimeout         time.Duration
	Logger                 *zap.Logger
	DB                     *pgxpool.Pool
	InstanceConnectionName string
	// usecase
	LikeUsecase domain.LikeUsecase
	PostUsecase domain.PostUsecase
}

func App() (*Application, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env not found, using fallbacks", err)
	}

	// Initialize logger
	logger, err := helper.InitLogger()
	if err != nil {
		log.Println("logger failed to load", err)
	}

	defer func() {
		if err := logger.Sync(); err != nil {
			log.Printf("failed to flush log: %v", err)
		}
	}()

	cfg := config.GetConfig()

	// Initialize database
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	db, err := util_db.PostgresDatabase(ctx, cfg.DBConnString)
	if err != nil {
		logger.Error("Failed to connect to database", zap.Error(err))
		return nil, err
	}

	// Initialize validator

	likeRepo := repository.NewLikeRepository(db)
	postRepo := post.NewPostRepository(db)

	likeUsecase := usecase.NewLikeUsecase(likeRepo, time.Duration(cfg.ContextTimeout)*time.Second)
	postUsecase := usecase.NewPostUsecase(postRepo, likeRepo, time.Duration(cfg.ContextTimeout)*time.Second)

	logger.Info("Application bootstrapped successfully")

	return &Application{
		Config:         cfg,
		ContextTimeout: time.Duration(cfg.ContextTimeout) * time.Second,
		Logger:         logger,
		DB:             db,
		LikeUsecase:    likeUsecase,
		PostUsecase:    postUsecase,
	}, nil
}
