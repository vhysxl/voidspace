package bootstrap

import (
	"context"
	"crypto/rsa"
	"database/sql"
	"fmt"
	"log"
	"time"
	"voidspace/users/config"
	"voidspace/users/database"
	"voidspace/users/internal/domain"
	"voidspace/users/internal/repository"
	"voidspace/users/internal/usecase"
	"voidspace/users/logger"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type Application struct {
	Config                 *config.Config
	DB                     *sql.DB
	ContextTimeout         time.Duration
	AccessTokenDuration    time.Duration
	RefreshTokenDuration   time.Duration
	PrivateKey             *rsa.PrivateKey
	Logger                 *zap.Logger
	InstanceConnectionName string
	RSAPrivateKey          string
	// use cases
	FollowUsecase  domain.FollowUsecase
	AuthUsecase    usecase.AuthUsecase
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

	fmt.Printf(`
Config:
  PublicHost:             %s
  Port:                   %s
  DBUser:                 %s
  DBPassword:             %s
  DBAddress:              %s
  DBName:                 %s
  ContextTimeout:         %d
  AccessTokenDuration:    %d
  RefreshTokenDuration:   %d
  InstanceConnectionName: %s
  RSAPrivateKey:          %s
`,
		cfg.PublicHost,
		cfg.Port,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBAddress,
		cfg.DBName,
		cfg.ContextTimeout,
		cfg.AccessTokenDuration,
		cfg.RefreshTokenDuration,
		cfg.InstanceConnectionName,
		cfg.RSAPrivateKey,
	)

	privateKey, err := config.LoadPrivateKey(cfg.RSAPrivateKey)
	if err != nil {
		logger.Error("Failed to load private key", zap.Error(err))
	}

	var dbConfig = &mysql.Config{
		User:   cfg.DBUser,
		Passwd: cfg.DBPassword,
		DBName: cfg.DBName,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.ContextTimeout)*time.Second)
	defer cancel()

	db, err := database.MySqlDatabase(ctx, dbConfig, cfg.InstanceConnectionName)
	if err != nil {
		logger.Error("Failed to connect to database", zap.Error(err))
		return nil, err
	}

	// registering repos, usecases and deps to the app
	userRepo := repository.NewUserRepository(db)
	profileRepo := repository.NewProfileRepository(db)
	followRepo := repository.NewFollowRepository(db)

	authUsecase := usecase.NewAuthUsecase(userRepo, time.Duration(cfg.ContextTimeout)*time.Second)
	userUsecase := usecase.NewUserUsecase(userRepo, time.Duration(cfg.ContextTimeout)*time.Second)
	profileUsecase := usecase.NewProfileUsecase(profileRepo, time.Duration(cfg.ContextTimeout)*time.Second)
	followUsecase := usecase.NewFollowUsecase(userRepo, followRepo, time.Duration(cfg.ContextTimeout)*time.Second)

	return &Application{
		Config:               cfg,
		DB:                   db,
		ContextTimeout:       time.Duration(cfg.ContextTimeout) * time.Second,
		AccessTokenDuration:  time.Duration(cfg.AccessTokenDuration) * time.Minute,
		RefreshTokenDuration: time.Duration(cfg.RefreshTokenDuration) * 24 * time.Hour,
		PrivateKey:           privateKey,
		Logger:               logger,
		AuthUsecase:          authUsecase,
		UserUsecase:          userUsecase,
		ProfileUsecase:       profileUsecase,
		FollowUsecase:        followUsecase,
	}, nil
}
