package bootstrap

import (
	"context"
	"database/sql"
	"log"
	"time"
	"voidspace/comments/config"
	"voidspace/comments/database"
	"voidspace/comments/internal/domain"
	"voidspace/comments/internal/repository"
	"voidspace/comments/internal/usecase"
	"voidspace/comments/logger"

	"github.com/go-playground/validator/v10"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type Application struct {
	Config         *config.Config
	DB             *sql.DB
	Validator      *validator.Validate
	ContextTimeout time.Duration
	Logger         *zap.Logger
	CommentUseCase domain.CommentUsecase
}

func App() (*Application, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env not found, using fallbacks", err)
	}

	logger, err := logger.InitLogger()
	if err != nil {
		log.Println("logger failed to load", err)
	}
	defer func() {
		if err := logger.Sync(); err != nil {
			log.Printf("failed to flush log: %v", err)
		}
	}()

	cfg := config.GetConfig()

	var dbConfig = mysql.Config{
		User:      cfg.DBUser,
		Passwd:    cfg.DBPassword,
		DBName:    cfg.DBName,
		Net:       "tcp",
		Addr:      cfg.DBAddress,
		ParseTime: true,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.ContextTimeout)*time.Second)
	defer cancel()

	db, err := database.MySqlDatabase(ctx, dbConfig)
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

	commentRepo := repository.NewCommentRepository(db)
	app.CommentUseCase = usecase.NewCommentUsecase(commentRepo, app.ContextTimeout)

	return app, nil
}
