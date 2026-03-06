package bootstrap

import (
	"context"
	"crypto/rsa"
	"log"
	"time"
	"voidspace/users/config"
	"voidspace/users/database"
	"voidspace/users/internal/domain"
	follow_repository "voidspace/users/internal/repository/follow"
	profile_repository "voidspace/users/internal/repository/profile"
	user_repository "voidspace/users/internal/repository/user"
	follow_usecase "voidspace/users/internal/usecase/follow"
	profile_usecase "voidspace/users/internal/usecase/profile"
	user_usecase "voidspace/users/internal/usecase/user"
	"voidspace/users/logger"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type Application struct {
	Config               *config.Config
	ContextTimeout       time.Duration
	AccessTokenDuration  time.Duration
	RefreshTokenDuration time.Duration
	PrivateKey           *rsa.PrivateKey
	Logger               *zap.Logger
	DB                   *pgxpool.Pool
	// InstanceConnectionString string
	// use cases
	FollowUsecase  domain.FollowUsecase
	ProfileUsecase domain.ProfileUsecase
	UserUsecase    domain.UserUsecase
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

	privateKey, err := config.LoadPrivateKey(cfg.SecretPath)
	if err != nil {
		logger.Error("Failed to load private key", zap.Error(err))
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.ContextTimeout)*time.Second)
	defer cancel()

	db, err := database.PostgresDatabase(ctx, cfg.DBConnectionString)
	if err != nil {
		logger.Error("Failed to connect to database", zap.Error(err))
		return nil, err
	}

	// registering repos, usecases and deps to the app
	userRepository := user_repository.NewUserRepository(db)
	profileRepository := profile_repository.NewProfileRepository(db)
	followRepository := follow_repository.NewFollowRepository(db)

	userUsecase := user_usecase.NewUserUsecase(userRepository, followRepository, time.Duration(cfg.ContextTimeout)*time.Second)
	profileUsecase := profile_usecase.NewProfileUsecase(profileRepository, time.Duration(cfg.ContextTimeout)*time.Second)
	followUsecase := follow_usecase.NewFollowUsecase(followRepository, time.Duration(cfg.ContextTimeout)*time.Second)

	return &Application{
		Config:               cfg,
		DB:                   db,
		ContextTimeout:       time.Duration(cfg.ContextTimeout) * time.Second,
		AccessTokenDuration:  time.Duration(cfg.AccessTokenDuration) * time.Minute,
		RefreshTokenDuration: time.Duration(cfg.RefreshTokenDuration) * 24 * time.Hour,
		PrivateKey:           privateKey,
		Logger:               logger,
		UserUsecase:          userUsecase,
		ProfileUsecase:       profileUsecase,
		FollowUsecase:        followUsecase,
	}, nil
}
