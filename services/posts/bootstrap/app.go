package bootstrap

import (
	"context"
	"database/sql"
	"log"
	"time"
	"voidspace/posts/config"
	"voidspace/posts/database"
	"voidspace/posts/internal/repository"
	"voidspace/posts/internal/usecase"
	"voidspace/posts/logger"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type Application struct {
	Config         *config.Config
	ContextTimeout time.Duration
	Validator      *validator.Validate
	Logger         *zap.Logger
	DB             *sql.DB
	// usecase
	LikeUsecase usecase.LikeUsecase
}

func App() (*Application, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env not found, using fallbacks", err)
	}
	cfg := config.GetConfig()

	// Initialize logger
	logger, err := logger.InitLogger()
	if err != nil {
		log.Println("Logger failed to load", err)
		return nil, err
	}

	// Initialize database
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	db, err := database.PostgresDatabase(ctx, cfg.GetDBConnectionString())
	if err != nil {
		logger.Error("Failed to connect to database", zap.Error(err))
		return nil, err
	}

	// Initialize validator
	validator := validator.New()

	likeRepo := repository.NewLikeRepository(db)

	likeUsecase := usecase.NewLikeUsecase(likeRepo, time.Duration(cfg.ContextTimeout)*time.Second)

	logger.Info("Application bootstrapped successfully")

	return &Application{
		Config:         cfg,
		ContextTimeout: time.Duration(cfg.ContextTimeout) * time.Second,
		Validator:      validator,
		Logger:         logger,
		DB:             db,
		LikeUsecase:    likeUsecase,
	}, nil
}
