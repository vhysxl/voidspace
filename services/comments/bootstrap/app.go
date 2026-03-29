package bootstrap

import (
	"context"
	"log"
	"time"
	"voidspace/comments/config"
	"voidspace/comments/internal/domain"
	"voidspace/comments/internal/repository/comment"
	"voidspace/comments/internal/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	util_db "github.com/vhysxl/voidspace/shared/utils/database"
	"github.com/vhysxl/voidspace/shared/utils/helper"
	"go.uber.org/zap"
)

type Application struct {
	Config         *config.Config
	ContextTimeout time.Duration
	Logger         *zap.Logger
	DB             *pgxpool.Pool
	Validator      *validator.Validate
	CommentUseCase domain.CommentUsecase
}

func App() (*Application, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env not found, using fallbacks", err)
	}

	logger, err := helper.InitLogger()
	if err != nil {
		log.Println("logger failed to load", err)
	}

	defer func() { _ = logger.Sync() }()

	cfg := config.GetConfig()

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.ContextTimeout)*time.Second)
	defer cancel()

	db, err := util_db.PostgresDatabase(ctx, cfg.DBConnString)
	if err != nil {
		logger.Error("Failed to connect to database", zap.Error(err))
		return nil, err
	}

	app := &Application{
		Config:         cfg,
		DB:             db,
		Validator:      validator.New(),
		ContextTimeout: time.Duration(cfg.ContextTimeout) * time.Second,
		Logger:         logger,
	}

	commentRepo := comment.NewCommentRepository(db)
	app.CommentUseCase = usecase.NewCommentUsecase(commentRepo, app.ContextTimeout)

	return app, nil
}
